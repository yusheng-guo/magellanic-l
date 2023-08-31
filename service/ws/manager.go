package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yushengguo557/magellanic-l/global"
	"log"
	"time"
)

const (
	ExchangeName = "websocket-messages-router"
	ExchangeType = "direct"

	QueueName = "websocket-messages-queue"

	WebSocketClientToServerMap = "websocket-client-to-server-map"
)

// WebSocketManager websocket管理器
type WebSocketManager struct {
	ID           string
	Clients      map[string]*Client
	Channels     map[string]*Channel
	Messages     chan Message // 消息处理通道
	MessageQueue MessageQueue // 当前服务器独占的消息队列 用于接收来自其他服务器发送的 websocket 消息
}

type MessageQueue struct {
	amqp.Queue
}

// Publish 发送消息 wid => websocket id;
func (mq *MessageQueue) Publish(wid string, msg Message) {
	body, err := json.Marshal(msg)
	if err != nil {
		log.Fatalln("marshal message, err:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = global.App.MQChannel.PublishWithContext(
		ctx,
		ExchangeName, // exchange
		wid,          // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Println("publish message to message queue, err:", err)
	}
}

// Consume 消费消息
func (mq *MessageQueue) Consume() (msgs chan Message, err error) {
	var ch = global.App.MQChannel
	var delivery <-chan amqp.Delivery
	var msg Message

	delivery, err = ch.Consume(
		mq.Name, // queue
		"",      // consumer
		false,   // auto-ack 消息确认
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	if err != nil {
		return nil, fmt.Errorf("register a consumer, err: %w", err)
	}

	go func() {
		for d := range delivery {
			if err = json.Unmarshal(d.Body, &msg); err != nil {
				log.Fatalln("unmarshal ")
			}
			msgs <- msg
		}
	}()
	return msgs, nil
}

// NewWebSocketManager 实例化websocket管理器
// cap: 消息通道容量
func NewWebSocketManager(cap int) *WebSocketManager {
	var err error
	var ch = global.App.MQChannel
	var id = uuid.NewString()

	// 1.声明交换机
	err = ch.ExchangeDeclare(
		ExchangeName, // name
		ExchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalln("queue declare, err:", err)
	}

	// 2.声明消息队列
	// wsmq: websocket message queue
	wsmq, err := ch.QueueDeclare(
		QueueName+"-"+id, // name
		false,            // durable
		false,            // delete when unused
		true,             // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Fatalln("queue declare, err:", err)
	}

	// 3.消息队列与交换机进行绑定
	err = ch.QueueBind(
		wsmq.Name,    // queue name
		id,           // routing key
		ExchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		log.Fatalln("queue bind, err:", err)
	}

	return &WebSocketManager{
		ID:           id,
		Clients:      make(map[string]*Client),
		Channels:     make(map[string]*Channel),
		Messages:     make(chan Message, cap),
		MessageQueue: MessageQueue{wsmq},
	}
}

// Register 使用 WebSocketManager 对 client 进行管理 & 接收客户端发送过来的所有消息
func (m *WebSocketManager) Register(client *Client) {
	var rdb = global.App.Redis
	var ret string
	var err error
	var msg Message

	// 1.保存 client -> manager id 映射 & client uid -> client 映射
	rdb.HSet(context.Background(), WebSocketClientToServerMap, client.UID, m.ID)
	m.Clients[client.UID] = client
	for {
		fmt.Printf("online population: %d\r", len(m.Clients))
		msg, err = client.Read()
		if err != nil {
			log.Println("read data when registering, err:", err)
			// 注销客户端 退出循环
			m.Logout(client.UID)
			break
		}
		ret, err = rdb.HGet(context.Background(), WebSocketClientToServerMap, msg.To).Result()
		if err != nil {
			return
		}
		if ret == m.ID {
			m.Messages <- msg
		} else {
			m.MessageQueue.Publish(msg.To, msg)
		}
	}
}

// Logout 取消 WebSocketManager 对 client 的管理 & 从所有频道中移除该客户端
func (m *WebSocketManager) Logout(uid string) {
	// 1.关闭连接
	m.Clients[uid].Conn.Close()

	// 2.移除管理
	delete(m.Clients, uid)

	// 3.移出频道
	for _, ch := range m.Channels {
		delete(ch.Members, uid)
	}

	// 4.删除映射
	val := global.App.Redis.HDel(context.Background(), WebSocketClientToServerMap, uid).Val()
	if val != 1 {
		log.Printf("failed delete client [%s] in redis\n", uid)
	}
}

// Broadcast 广播消息
func (m *WebSocketManager) Broadcast(msg Message) (err error) {
	for _, c := range m.Clients {
		err = c.Write(msg)
	}
	return err
}

// ReceiveMessage 接收消息
func (m *WebSocketManager) ReceiveMessage() {
	msgs, err := m.MessageQueue.Consume()
	if err != nil {
		log.Fatalln("receive message, err:", err)
	}

	for msg := range msgs {
		m.Messages <- msg
	}
}

// HandlerMessage 处理消息
func (m *WebSocketManager) HandlerMessage() {
	var err error
	var echo Message
	for msg := range m.Messages {
		switch msg.Type {
		case MessageTypeRegister:
			if m.IsManaged(msg.From) {
				echo = NewMessage(MessageTypeRegister, []byte("success"), "", msg.From)
			} else {
				echo = NewMessage(MessageTypeRegister, []byte("failed"), "", msg.From)
			}
			err = m.Clients[echo.To].Write(echo)
		case MessageTypeLogout:
			if m.IsManaged(msg.From) {
				m.Logout(msg.From)
			}
		case MessageTypeHeartbeat:
			if m.IsManaged(msg.From) {
				echo = NewMessage(MessageTypeHeartbeat, []byte("success"), "", msg.From)
				err = m.Clients[echo.To].Write(echo)
			}
		case MessageTypeOneOnOne:
			// TODO: 私聊
		case MessageTypeGroup:
			// TODO: 群聊
		case MessageTypeChannel:
		// err = m.Channels[msg.To].Write(msg)
		case MessageTypeBroadcast:
			err = m.Broadcast(msg)
		case MessageTypeEcho:
			if m.IsManaged(msg.To) {
				err = m.Clients[msg.To].Write(msg)
			}
		default:
			echo = NewMessage(MessageTypeEcho, []byte("format err, can't parse"), "", msg.From)
			err = m.Clients[echo.To].Write(echo)
		}
		if err != nil {
			log.Println("handle message, err:", err)
		}
	}
}

func (m *WebSocketManager) PushMessage(msg Message) {
	m.Messages <- msg
}

// IsManaged 判断用户是否被 websocket 管理器 所管理
func (m *WebSocketManager) IsManaged(uid string) bool {
	_, ok := m.Clients[uid]
	return ok
}

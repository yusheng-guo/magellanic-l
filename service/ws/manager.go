package ws

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yushengguo557/magellanic-l/global"
	"log"
	"time"
)

// WebSocketManager websocket管理器
type WebSocketManager struct {
	Clients      map[string]*Client
	Channels     map[string]*Channel
	Messages     chan Message
	MessageQueue MessageQueue
}

type MessageQueue struct {
	amqp.Queue
}

// Publish 发送消息
func (mq *MessageQueue) Publish(msg Message) {
	body, err := json.Marshal(msg)
	if err != nil {
		log.Fatalln("marshal message, err:", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err = global.App.MQChannel.PublishWithContext(
		ctx,
		"",      // exchange
		mq.Name, // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Println("publish message to message queue, err:", err)
	}
	cancel()
}

// Consume 消费消息
func (mq *MessageQueue) Consume() error {
	msgs, err := global.App.MQChannel.Consume(
		global.App.WebSocketManager.MessageQueue.Name, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return fmt.Errorf("register a consumer, err:%w", err)
	}
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
	}
	return nil
}

// NewWebSocketManager 实例化websocket管理器
// cap: 消息通道容量
func NewWebSocketManager(cap int) *WebSocketManager {
	// wsmq: websocket message queue
	wsmq, err := global.App.MQChannel.QueueDeclare(
		"websocket-messages", // name
		false,                // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		log.Fatalln("queue declare, err:", err)
	}
	return &WebSocketManager{
		Clients:      make(map[string]*Client),
		Channels:     make(map[string]*Channel),
		Messages:     make(chan Message, cap),
		MessageQueue: MessageQueue{wsmq},
	}
}

// Register 使用 WebSocketManager 对 client 进行管理 & 接收客户端发送过来的所有消息
func (m *WebSocketManager) Register(client *Client) {
	m.Clients[client.UID] = client
	for {
		fmt.Printf("online population: %d\r", len(m.Clients))
		msg, err := client.Read()
		if err != nil {
			log.Println("read data when registering, err:", err)
			// 注销客户端 退出循环
			m.Logout(client.UID)
			break
		}
		m.Messages <- msg
		m.MessageQueue.Publish(msg)
	}
}

// Logout 取消 WebSocketManager 对 client 的管理 & 从所有频道中移除该客户端
func (m *WebSocketManager) Logout(uid string) {
	// 1.关闭连接
	m.Clients[uid].Conn.Close()

	// 2.移除管理
	delete(m.Clients, uid)

	for _, ch := range m.Channels {
		delete(ch.Members, uid)
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
func (m *WebSocketManager) ReceiveMessage() error {
	return nil
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
		case MessageTypeOneToOne:
			// TODO: 私聊
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
			// 避免循环引用
			//global.App.Log.Error(err.Error())
			//zap.L().Error(err.Error())
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

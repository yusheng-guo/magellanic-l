package ws

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type MessageQueue struct {
	Queue   *amqp.Queue
	Channel *amqp.Channel
}

// NewMessageQueue 实例化 websocket 消息队列
func NewMessageQueue(id string, ch *amqp.Channel) MessageQueue {
	var err error
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
	queue, err := ch.QueueDeclare(
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
		queue.Name,   // queue name
		id,           // routing key
		ExchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		log.Fatalln("queue bind, err:", err)
	}
	return MessageQueue{
		Queue:   &queue,
		Channel: ch,
	}
}

// Publish 发送消息 wid => websocket id;
func (mq *MessageQueue) Publish(wid string, msg Message) {
	body, err := json.Marshal(msg)
	if err != nil {
		log.Fatalln("marshal message, err:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = mq.Channel.PublishWithContext(
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
	var delivery <-chan amqp.Delivery
	var msg Message

	delivery, err = mq.Channel.Consume(
		mq.Queue.Name, // queue
		"",            // consumer
		false,         // auto-ack 消息确认
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
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

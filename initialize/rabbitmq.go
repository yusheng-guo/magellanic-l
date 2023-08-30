package initialize

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yushengguo557/magellanic-l/global"
	"log"
)

// InitRabbitMQ initialize message queue
func InitRabbitMQ() {
	conn, err := amqp.Dial("amqp://rabbitmq:8f4a4e91d3@119.91.204.226:5672/")
	if err != nil {
		log.Fatalln("connect to RabbitMQ, err:", err)
	}

	task1 := global.NewDeferTask(func(a ...any) {
		var err error
		err = a[0].(*amqp.Connection).Close()
		if err != nil {
			log.Println("close RabbitMQ connect, err:", err)
		}
	}, conn)
	global.DeferTaskQueue = append(global.DeferTaskQueue, task1)

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln("open RabbitMQ channel, err:", err)
	}

	task2 := global.NewDeferTask(func(a ...any) {
		var err error
		err = a[0].(*amqp.Channel).Close()
		if err != nil {
			log.Println("close RabbitMQ channel, err:", err)
		}
	}, ch)
	global.DeferTaskQueue = append(global.DeferTaskQueue, task2)

	global.App.MQChannel = ch
	log.Println("You successfully connected to RabbitMQ!")

	//q, err := ch.QueueDeclare(
	//	"hello", // name
	//	false,   // durable
	//	false,   // delete when unused
	//	false,   // exclusive
	//	false,   // no-wait
	//	nil,     // arguments
	//)
	//if err != nil {
	//	log.Fatalln("declare a queue, err:", err)
	//}
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	//body := "Hello World!"
	//err = ch.PublishWithContext(ctx,
	//	"",     // exchange
	//	q.Name, // routing key
	//	false,  // mandatory
	//	false,  // immediate
	//	amqp.Publishing{
	//		ContentType: "text/plain",
	//		Body:        []byte(body),
	//	})
	//if err != nil {
	//	log.Fatalln("publish a message, err:", err)
	//}
	//log.Printf(" [x] Sent %s\n", body)

	//msgs, err := ch.Consume(
	//	q.Name, // queue
	//	"",     // consumer
	//	true,   // auto-ack
	//	false,  // exclusive
	//	false,  // no-local
	//	false,  // no-wait
	//	nil,    // args
	//)
	//if err != nil {
	//	log.Fatalln("deliver queued messages, err:", err)
	//}

	//go func() {
	//	for d := range msgs {
	//		log.Printf("Received a message: %s\n", d.Body)
	//	}
	//}()
	//log.Println(" [*] Waiting for messages. To exit press CTRL+C")
}

package initialize

import (
	"context"
	"fmt"
	"github.com/yushengguo557/magellanic-l/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// InitMongoDB 初始化 MongoDB
func InitMongoDB() {
	// 1.Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI("mongodb://mongo:8f4a4e91d3@119.91.204.226:27017/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPI)

	// 2.Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalln(fmt.Errorf("connect MongoDB, err: %w", err))
	}

	// 3.close connect
	task := global.NewDeferTask(func(a ...any) {
		var err error
		err = a[0].(*mongo.Client).Disconnect(context.TODO())
		if err != nil {
			log.Println("disconnect MongoDB, err:", err)
		}
	}, client)
	global.DeferTaskQueue = append(global.DeferTaskQueue, task)

	// 4.Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		log.Fatalln(fmt.Errorf("ping MongoDB, err: %w", err))
	}
	log.Println("You successfully connected to MongoDB!")

	// 5.Assigning a value to a global variable
	global.App.MongoDB = client
}

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
		ApplyURI("mongodb+srv://yushengguo557:Hr8CKHgYhjInISpK@cluster0.zecytgm.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPI)

	// 2.Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalln(fmt.Errorf("connect MongoDB, err: %w", err))
	}

	// 3.close connect
	global.DeferFuncList.Push(
		func() {
			if err = client.Disconnect(context.TODO()); err != nil {
				log.Fatalln(fmt.Errorf("disconnect MongoDB, err: %w", err))
			}
		})

	// 4.Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		log.Fatalln(fmt.Errorf("ping, err: %w", err))
	}
	log.Println("You successfully connected to MongoDB!")

	// 5.Assigning a value to a global variable
	global.App.MongoDB = client
}

package main

import "github.com/yushengguo557/magellanic-l/initialize"

func init() {
	// 1.initial config
	initialize.InitConfig()

	// 2.initial log
	initialize.InitLog()

	// 3.initial database
	//MongoDB TiDB DGraph Redis
	initialize.InitMongoDB()
	initialize.InitTiDB()
	initialize.InitRedis()
	initialize.InitDGraph()

	// 4.initial RabbitMQ
	initialize.InitRabbitMQ()

	// 5.initial websocket manager
	initialize.InitWebSocketManager()

	// 5.initial gin
	initialize.InitEngine()
}

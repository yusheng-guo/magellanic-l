package main

import "github.com/yushengguo557/magellanic-l/initialize"

func init() {
	// 1.initial config
	initialize.InitConfig()

	// 2.initial log
	initialize.InitLog()

	//TODO 初始化数据库连接
	//MongoDB TiDB DGraph Redis
	initialize.InitMongoDB()
	initialize.InitTiDB()
	initialize.InitRedis()

	//TODO 初始化消息队列 RabbitMQ
	initialize.InitRabbitMQ()

	// 4.initial gin
	initialize.InitEngine()
}

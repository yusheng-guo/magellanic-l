package main

import "github.com/yushengguo557/magellanic-l/initialize"

func init() {
	// 1.initial config
	initialize.InitConfig()

	//TODO 初始化日志

	//TODO 初始化数据库连接

	//TODO 初始化Gin框架引擎
	initialize.InitEngine()
}

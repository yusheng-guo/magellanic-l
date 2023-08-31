package initialize

import (
	"github.com/yushengguo557/magellanic-l/global"
	"github.com/yushengguo557/magellanic-l/service/ws"
	"log"
)

const MessageChannelCapacity = 1024

func InitWebSocketManager() {
	// 1.实例化管理器
	manager := ws.NewWebSocketManager(MessageChannelCapacity)

	// 2.接收消息 (来自其他服务器)
	go manager.ReceiveMessage()

	// 2.处理消息
	go manager.HandlerMessage()

	// 3.赋值到全局变量
	global.App.WebSocketManager = manager

	log.Println("You successfully init websocket manager!")
}

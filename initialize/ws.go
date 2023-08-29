package initialize

import (
	"github.com/yushengguo557/magellanic-l/global"
	"github.com/yushengguo557/magellanic-l/service/ws"
)

const MessageChannelCapacity = 1024

func InitWebSocketManager() {
	// 1.实例化管理器
	manager := ws.NewWebSocketManager(MessageChannelCapacity)

	// 2.处理消息
	go manager.HandlerMessage()

	// 3.赋值到全局变量
	global.App.WebSocketManager = manager
}

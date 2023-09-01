package initialize

import (
	"github.com/google/uuid"
	"github.com/yushengguo557/magellanic-l/global"
	"github.com/yushengguo557/magellanic-l/service/ws"
	"log"
)

const MessageChannelCapacity = 1024

func InitWebSocketManager() {
	// 1.实例化管理器
	id := uuid.NewString()
	mq := ws.NewMessageQueue(id, global.App.MQChannel)
	manager := ws.NewWebSocketManager(id, MessageChannelCapacity, global.App.Redis, mq)

	// 2.接收消息 (来自消息队列)
	go manager.ReceiveMessage()

	// 2.处理消息
	go manager.HandleMessage()

	// 3.赋值到全局变量
	global.App.WebSocketManager = manager
	log.Println("You successfully init websocket manager!")

	// 4.注销当前管理器所有的客户端
	task := global.NewDeferTask(func(a ...any) {
		for uid := range manager.Clients {
			a[0].(*ws.WebSocketManager).Logout(uid)
		}
		log.Printf("logout all client of manager [%s]\n", manager.ID)
	}, manager)
	global.DeferTaskQueue = append(global.DeferTaskQueue, task)
}

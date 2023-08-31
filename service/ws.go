package service

import (
	"github.com/yushengguo557/magellanic-l/global"
	"github.com/yushengguo557/magellanic-l/service/ws"
	"golang.org/x/net/websocket"
	"net/http"
)

func WebSocketHandel(uid string, w http.ResponseWriter, req *http.Request) {
	websocket.Handler(func(conn *websocket.Conn) {
		client := ws.NewClient(uid, conn)
		manager := global.App.WebSocketManager
		manager.Register(client) // 注册客户端
		manager.Logout(uid)      // 注销客户端
	}).ServeHTTP(w, req)
}

//func Echo(conn *websocket.Conn) {
//	defer conn.Close()
//	var err error
//
//	for {
//		var reply string
//
//		if err = websocket.Message.Receive(conn, &reply); err != nil {
//			fmt.Println("Can't receive")
//			break
//		}
//
//		fmt.Println("Received back from client: " + reply)
//
//		msg := reply
//		fmt.Println("Sending to client: " + msg)
//
//		if err = websocket.Message.Send(conn, msg); err != nil {
//			fmt.Println("Can't send")
//			break
//		}
//	}
//}

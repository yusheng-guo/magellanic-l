package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

func WebSocket(c *gin.Context) {
	handler := websocket.Handler(Echo)
	handler.ServeHTTP(c.Writer, c.Request)
}

func Echo(conn *websocket.Conn) {
	defer conn.Close()
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(conn, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(conn, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

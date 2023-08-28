package v1

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

func WebSocket(c *gin.Context) {
	handler := websocket.Handler(EchoServer)
	handler.ServeHTTP(c.Writer, c.Request)
}

func EchoServer(conn *websocket.Conn) {
	var msg = make([]byte, 1024)
	for {
		conn.Read(msg)
		conn.Write(msg)
	}
}

package router

import (
	v1 "github.com/yushengguo557/magellanic-l/controllers/api/v1"
)

// AddWebSocket 添加 websocket 相关路由
func (g *Group) AddWebSocket() {
	g.GET("/ws", v1.WebSocket)
	//g.GET("/ws", func(c *gin.Context) {
	//	websocket.Handler(echo).ServeHTTP(c.Writer, c.Request)
	//})
}

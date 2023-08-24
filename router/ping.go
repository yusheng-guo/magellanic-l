package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddPing 添加 ping 路由 测试服务器是否正常运行
func (g *Group) AddPing() {
	g.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
}

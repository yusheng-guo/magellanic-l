package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/magellanic-l/internal/middleware"
	"net/http"
)

func Handler(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	r.GET("/ws", WebSocket)

	user := r.Group("/user")
	{
		user.POST("/login", Login)       // 登录
		user.POST("/register", Register) // 注册

		user.PUT("/info/:id", middleware.JwtAuth(), EditInfo)   // 修改个人信息
		user.GET("/info/:id", middleware.JwtAuth(), ObtainInfo) // 获取个人信息
	}
}

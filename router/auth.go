package router

import (
	v1 "github.com/yushengguo557/magellanic-l/controllers/api/v1"
	"github.com/yushengguo557/magellanic-l/middleware"
)

func (g *Group) AddAuth() {
	auth := g.POST("/auth")
	auth.POST("/login", v1.Login)       // 登录
	auth.POST("/register", v1.Register) // 注册

	auth.PUT("/info/:id", middleware.JwtAuth(), v1.EditInfo)   // 修改个人信息
	auth.GET("/info/:id", middleware.JwtAuth(), v1.ObtainInfo) // 获取个人信息
}

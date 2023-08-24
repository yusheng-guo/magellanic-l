package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/magellanic-l/global"
	"os"
)

// InitEngine 初始化路由器
func InitEngine() {
	// 1.设置Gin模式
	switch os.Getenv("GOENV") {
	case "test":
		gin.SetMode(gin.TestMode)
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	// 2.创建引擎
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	// 3.赋值全局变量
	global.App.Engine = engine
}

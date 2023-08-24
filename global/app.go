package global

import (
	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/magellanic-l/config"
	"go.uber.org/zap"
)

var App = new(Application)

type Application struct {
	Config config.Configuration
	Log    *zap.Logger
	Engine *gin.Engine
}

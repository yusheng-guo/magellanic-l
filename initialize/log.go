package initialize

import (
	"fmt"
	"github.com/yushengguo557/magellanic-l/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

// InitLog 初始化日志
func InitLog() {
	var err error
	//log := global.App.Config.Log
	zapConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     zap.NewProductionEncoderConfig(),
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		InitialFields:     nil,
	}
	global.App.Log, err = zapConfig.Build()
	if err != nil {
		log.Panic(fmt.Errorf("init log, err: %w", err))
	}
}

package initialize

import (
	"github.com/yushengguo557/magellanic-l/global"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"log"
	"os"
)

// InitLog 初始化日志
func InitLog() {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	global.App.Log = zap.New(core, zap.AddCaller())

	// 添加延迟函数 程序结束后执行
	global.DeferFuncList.Push(
		func() {
			var err error
			err = global.App.Log.Sync()
			if err != nil {
				log.Fatalf("zap logger sync, err:%+v\n", err)
			}
		})

	// zap.ReplaceGlobals(logger)
	// zap.L()
}

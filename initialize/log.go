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
	logger := zap.New(core, zap.AddCaller())

	// 添加延迟函数 程序结束后执行
	task := global.NewDeferTask(func(a ...any) {
		var err error
		err = a[0].(*zap.Logger).Sync()
		if err != nil {
			log.Fatalf("zap logger sync, err:%+v\n", err)
		}
	}, logger)
	global.DeferTaskQueue = append(global.DeferTaskQueue, task)

	global.App.Log = logger
	// zap.ReplaceGlobals(logger)
	// zap.L()
}

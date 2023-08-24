package main

import (
	"context"
	"github.com/yushengguo557/magellanic-l/global"
	"github.com/yushengguo557/magellanic-l/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	var err error

	// 定义路由
	router.DefineRouter()

	// 启动服务
	server := http.Server{
		Addr:    ":9999",
		Handler: global.App.Engine,
	}
	go func() {
		if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed:%+v\n", err)
	}
	log.Println("server shutdown")

	global.DeferFuncList.Run() // 释放资源

	global.App.Log.Info("hi")
}

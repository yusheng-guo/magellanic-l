package main

import (
	"context"
	"github.com/yushengguo557/magellanic-l/global"
	"github.com/yushengguo557/magellanic-l/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

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
		if err = server.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed:%+v\n", err)
	}
	log.Println("server shutdown")
}

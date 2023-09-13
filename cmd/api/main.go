package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/magellanic-l/global"
	"github.com/yushengguo557/magellanic-l/internal/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func setGinMode() {
	switch os.Getenv("GOENV") {
	case "prod":
		fmt.Println("currently a production environment")
		gin.SetMode(gin.ReleaseMode)
	case "test":
		fmt.Println("currently a test environment")
		gin.SetMode(gin.TestMode)
	default:
		fmt.Println("currently a development environment")
		gin.SetMode(gin.DebugMode)
	}
}

func main() {
	var err error

	// set gin engine mode
	setGinMode()
	var r = gin.New()
	handlers.Handler(r)

	fmt.Println("starting magellanic-l service...")

	// 启动服务
	server := http.Server{
		Addr:    ":9999",
		Handler: r,
	}
	go func() {
		if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		log.Println("listen:", server.Addr)
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

	// 释放资源
	releaseResources()
}

func releaseResources() {
	wg := sync.WaitGroup{}
	wg.Add(len(global.DeferTaskQueue))
	for _, deferTask := range global.DeferTaskQueue {
		go func(deferTask *global.DeferTask) {
			defer wg.Done()
			deferTask.Execute()
			log.Printf("%#v\n", deferTask.Params)
		}(deferTask)
	}
	wg.Wait()
}

package bootstrap

import (
	"context"
	"course-system/global"
	"course-system/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// RegisterRouter 注册路由
func RegisterRouter() *gin.Engine {
	router := gin.Default()
	// 注册 api 分组路由
	apiGroup := router.Group("/api/v1")
	routes.SetApiGroupRoutes(apiGroup)
	return router
}

// RunServer 启动服务器
func RunServer() {
	r := RegisterRouter()

	server := http.Server{
		Addr:         ":" + global.App.Config.App.Port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server listen err:%s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// 在此阻塞
	<-quit
	server.SetKeepAlivesEnabled(false)
	ctx, channel := context.WithTimeout(context.Background(), 1*time.Second)

	defer channel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown error")
	}
	fmt.Println("server exiting...")
}

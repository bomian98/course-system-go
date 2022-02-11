package main

import (
	"course-system/bootstrap"
	"course-system/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()

	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 启动服务器
	r.Run(":" + global.App.Config.App.Port)
}

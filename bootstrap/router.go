package bootstrap

import (
	"System/global"
	"System/routes"
	"github.com/gin-gonic/gin"
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
	//r.Run(":" + global.App.Config.App.Port)
	r.Run(":" + global.App.Config.App.Port) // 默认使用8080端口
}
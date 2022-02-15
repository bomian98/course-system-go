package routes

import (
	"System/app/controller"
	"github.com/gin-gonic/gin"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {

	// 成员管理
	router.POST("/member/create", controller.CreateUser)
	router.GET("/member")
	router.GET("/member/list")
	router.POST("/member/update")
	router.POST("/member/delete")

	// 登录
	router.POST("/auth/login")
	router.POST("/auth/logout")
	router.GET("/auth/whoami")

	// 排课
	router.POST("/course/create")
	router.GET("/course/get")

	router.POST("/teacher/bind_course")
	router.POST("/teacher/unbind_course")
	router.GET("/teacher/get_course")
	router.POST("/course/schedule")

	// 抢课
	router.POST("/student/book_course")
	router.GET("/student/course")
}

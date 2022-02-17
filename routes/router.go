package routes

import (
	"System/app/controllers"
	"System/global"
	"System/utils"
	"github.com/gin-gonic/gin"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {

	// 成员管理
	router.POST("/member/create", utils.Authorize(global.App.E), controllers.CreateUser)
	router.GET("/member", controllers.GetUser)
	router.GET("/member/list", controllers.GetsUser)
	router.POST("/member/update", controllers.UpdateUser)
	router.POST("/member/delete", controllers.DeleteUser)

	// 登录
	router.POST("/auth/login", controllers.Login)
	router.POST("/auth/logout", controllers.Logout)
	router.GET("/auth/whoami", controllers.WhoAmI)

	// 排课
	router.POST("/course/create")
	router.GET("/course/get")

	router.POST("/teacher/bind_course")
	router.POST("/teacher/unbind_course")
	router.GET("/teacher/get_course")
	router.POST("/course/schedule")

	// 抢课
	router.POST("/student/book_course", controllers.BookCourse)
	router.GET("/student/course", controllers.GetStudentCourse)
}

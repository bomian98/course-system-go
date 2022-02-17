package controllers

import (
	"System/app/common"
	"System/app/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func BookCourse(c *gin.Context) {
	fmt.Println("访问到该controller了")                 // 不需要 or 后期合并时，注释掉
	var request common.BookCourseRequest           // 声明待绑定的输入数据
	if err := c.ShouldBind(&request); err != nil { // 入参绑定错误，返回错误
		c.JSON(http.StatusOK, common.BookCourseResponse{
			Code: common.ParamInvalid,
		})
		return
	}
	// 判断学生是否存在
	//if !services.UserCourseService.IsStudentExisted(stu_id_uint) {
	//	c.JSON(http.StatusOK, common.BookCourseResponse{Code: common.StudentNotExisted})
	//	return
	//}
	// 判断课程编号是否存在
	//if !services.CourseService.IsCourseExisted(cos_id_uint) {
	//	c.JSON(http.StatusOK, common.BookCourseResponse{Code: common.CourseNotExisted})
	//	return
	//}
	// 插入数据
	stu_id, _ := strconv.ParseInt(request.StudentID, 10, 64)
	cos_id, _ := strconv.ParseInt(request.CourseID, 10, 64)
	//usercourse := models.UserCourse{UserID: stu_id, CourseID: cos_id}
	//services.UserCourseService.InsertUserCourse(usercourse)
	// 请求课程
	code := services.UserCourseService.BookCourse(stu_id, cos_id)
	c.JSON(http.StatusOK, common.BookCourseResponse{Code: code})
	return
}

func GetStudentCourse(c *gin.Context) {
	fmt.Println("访问到该controller了") // 不需要 or 后期合并时，注释掉
	var request common.GetStudentCourseRequest
	if err := c.ShouldBindJSON(&request); err != nil { // 入参绑定错误，返回错误
		c.JSON(http.StatusOK, common.GetStudentCourseResponse{
			Code: common.ParamInvalid,
		})
		return
	}
	//stu_id, _ := strconv.ParseInt(request.StudentID, 10, 64)
	//services.UserCourseService.GetUserCourses(stu_id)
	//// 学生不存在，返回错误
	//if !services.UserCourseService.IsStudentExisted(stu_id_uint) {
	//	c.JSON(http.StatusOK, common.BookCourseResponse{Code: common.StudentNotExisted})
	//	return
	//}
	//c.JSON(http.StatusOK, common.GetStudentCourseResponse{
	//	Code: common.StudentNotExisted,
	//})
	// 获取学生的课程
	// 获取学生课表
	// 返回
	// ctx.JSON()
}

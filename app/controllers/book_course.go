package controllers

import (
	"course-system/app/common"
	"course-system/app/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BookCourse(c *gin.Context) {
	//fmt.Println("访问到该controller了")                 // 不需要 or 后期合并时，注释掉
	var request common.BookCourseRequest               // 声明待绑定的输入数据
	if err := c.ShouldBindJSON(&request); err != nil { // 入参绑定错误，返回错误
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
	//usercourse := models.UserCourse{UserID: stu_id, CourseID: cos_id}
	//services.UserCourseService.InsertUserCourse(usercourse)
	// 请求课程
	code := services.UserCourseService.BookCourse(request.StudentID, request.CourseID)
	c.JSON(http.StatusOK, gin.H{"Code": code})
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

	code, courseListStruct := services.UserCourseService.GetUserCourses(request.StudentID)
	c.JSON(http.StatusOK, common.GetStudentCourseResponse{
		Code: code, Data: courseListStruct,
	})
	return
}

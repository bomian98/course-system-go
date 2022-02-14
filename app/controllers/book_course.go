package controllers

import (
	"course-system/app/common"
	"course-system/app/models"
	"course-system/app/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func BookCourse(c *gin.Context) {
	fmt.Println("访问到该controller了") // 不需要 or 后期合并时，注释掉
	var request common.BookCourseRequest
	if err := c.ShouldBind(&request); err != nil { // 入参绑定错误，返回错误
		c.JSON(http.StatusBadRequest, common.BookCourseResponse{
			Code: common.ParamInvalid,
		})
		return
	}
	stu_id, _ := strconv.ParseInt(request.StudentID, 10, 64)
	cos_id, _ := strconv.ParseInt(request.CourseID, 10, 64)
	usercourse := models.UserCourse{UserID: stu_id, CourseID: cos_id}
	services.UserCourseService.InsertUserCourse(usercourse)

	//if !services.UserCourseService.IsStudentExisted(stu_id_uint) {
	//	c.JSON(http.StatusOK, common.BookCourseResponse{Code: common.StudentNotExisted})
	//	return
	//}
	//if !services.CourseService.IsCourseExisted(cos_id_uint) {
	//	c.JSON(http.StatusOK, common.BookCourseResponse{Code: common.CourseNotExisted})
	//	return
	//}
	// 请求课程
	// 课程容量不足
	c.JSON(http.StatusOK, common.BookCourseResponse{
		Code: common.CourseNotAvailable,
	})
	// 课程请求成功
	c.JSON(http.StatusOK, common.BookCourseResponse{
		Code: common.CourseHasBound,
	})
	return
}

func GetStudentCourse(c *gin.Context) {
	fmt.Println("访问到该controller了") // 不需要 or 后期合并时，注释掉
	var request common.GetStudentCourseRequest
	if err := c.ShouldBindJSON(&request); err != nil { // 入参绑定错误，返回错误
		c.JSON(http.StatusBadRequest, common.GetStudentCourseResponse{
			Code: common.ParamInvalid,
		})
		return
	}
	stu_id, _ := strconv.ParseInt(request.StudentID, 10, 64)
	services.UserCourseService.GetUserCourses(stu_id)

	//// 学生不存在，返回错误
	//c.JSON(http.StatusOK, common.GetStudentCourseResponse{
	//	Code: common.StudentNotExisted,
	//})
	//// 课程不存在，返回错误
	//c.JSON(http.StatusOK, common.GetStudentCourseResponse{
	//	Code: common.CourseNotExisted,
	//})

	// 获取学生的课程
	// 获取学生课表

	// 返回
	// ctx.JSON()
}

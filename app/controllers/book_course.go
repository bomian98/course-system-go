package controllers

import (
	"course-system/app/common"
	"course-system/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BookCourse(c *gin.Context) {
	var request common.BookCourseRequest // 声明待绑定的输入数据
	var response common.BookCourseResponse
	if err := c.ShouldBindJSON(&request); err != nil { // 入参绑定错误，返回错误
		response.Code = common.ParamInvalid
	} else {
		response.Code = services.UserCourseService.BookCourse(request.StudentID, request.CourseID)
	}
	c.JSON(http.StatusOK, response)
}

func GetStudentCourse(c *gin.Context) {
	var request common.GetStudentCourseRequest
	var response common.GetStudentCourseResponse
	if err := c.ShouldBindJSON(&request); err != nil { // 入参绑定错误，返回错误
		response.Code = common.ParamInvalid
		response.Data.CourseList = make([]common.TCourse, 0)
	} else {
		response.Data.CourseList, response.Code = services.UserCourseService.GetUserCourses(request.StudentID)
	}
	c.JSON(http.StatusOK, response)
}

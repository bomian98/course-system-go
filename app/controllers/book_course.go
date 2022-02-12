package controllers

import (
	"course-system/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BookCourse(c *gin.Context) {
	var request common.BookCourseRequest
	if err := c.ShouldBindJSON(&request); err != nil { // 入参绑定错误，返回错误
		c.JSON(http.StatusBadRequest, common.BookCourseResponse{
			Code: common.ParamInvalid,
		})
		return
	}
	// 请求课程
	// 课程容量不足
	c.JSON(http.StatusOK, common.BookCourseResponse{
		Code: common.CourseNotAvailable,
	})
	//
	c.JSON(http.StatusOK, common.BookCourseResponse{
		Code: common.CourseHasBound,
	})
	return
}

func GetStudentCourse(c *gin.Context) {
	var request common.GetStudentCourseRequest
	if err := c.ShouldBindJSON(&request); err != nil { // 入参绑定错误，返回错误
		c.JSON(http.StatusBadRequest, common.GetStudentCourseResponse{
			Code: common.ParamInvalid,
		})
		return
	}
	// 学生不存在，返回错误
	c.JSON(http.StatusOK, common.GetStudentCourseResponse{
		Code: common.StudentNotExisted,
	})
	// 课程不存在，返回错误
	c.JSON(http.StatusOK, common.GetStudentCourseResponse{
		Code: common.CourseNotExisted,
	})

	// 获取学生的课程
	// 获取学生课表

	// 返回
	// ctx.JSON()
}

package controllers

import (
	"course-system/app/common"
	"course-system/app/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCourse(c *gin.Context) {
	var request common.CreateCourseRequest
	var response common.CreateCourseResponse
	if err := c.ShouldBind(&request); err != nil { // 入参绑定错误，返回错误
		response.Code = common.ParamInvalid
		response.Data.CourseID = ""
		fmt.Println(err)
	} else {
		courseID, code := services.CourseService.CreateCourse(request.Name, request.Cap)
		response.Code = code
		response.Data.CourseID = courseID
	}
	c.JSON(http.StatusOK, response)
	return
}

func GetCourse(c *gin.Context) {
	var request common.GetCourseRequest
	var response common.GetCourseResponse
	if err := c.ShouldBind(&request); err != nil { // 入参绑定错误，返回错误
		response.Code = common.ParamInvalid
		response.Data = *new(common.TCourse)
	} else {
		tCourse, code := services.CourseService.GetCourse(request.CourseID)
		response.Code = code
		response.Data = tCourse
	}
	c.JSON(http.StatusOK, response)
	return
}

func BindCourse(c *gin.Context) {
	var request common.BindCourseRequest
	var response common.BindCourseResponse
	if err := c.ShouldBind(&request); err != nil { // 入参绑定错误，返回错误
		response.Code = common.ParamInvalid
	} else {
		response.Code = services.CourseService.BindCourse(request.CourseID, request.TeacherID)
	}
	c.JSON(http.StatusOK, response)
	return
}

func UnbindCourse(c *gin.Context) {
	var request common.UnbindCourseRequest
	var response common.UnbindCourseResponse
	if err := c.ShouldBind(&request); err != nil { // 入参绑定错误，返回错误
		response.Code = common.ParamInvalid
	} else {
		response.Code = services.CourseService.UnBindCourse(request.CourseID, request.TeacherID)
	}
	c.JSON(http.StatusOK, response)
	return
}

func GetTeacherCourse(c *gin.Context) {
	var request common.GetTeacherCourseRequest
	var response common.GetTeacherCourseResponse
	var CourseList []*common.TCourse
	if err := c.ShouldBind(&request); err != nil { // 入参绑定错误，返回错误
		response.Code = common.ParamInvalid
		response.Data.CourseList = CourseList
	} else {
		response.Data.CourseList, response.Code = services.CourseService.GetTeacherCourse(request.TeacherID)
	}
	c.JSON(http.StatusOK, response)
	return
}

func MakeSchedule(c *gin.Context) {
	var request common.ScheduleCourseRequest
	var response common.ScheduleCourseResponse
	if err := c.ShouldBind(&request); err != nil { // 入参绑定错误，返回错误
		response.Code = common.ParamInvalid
		response.Data = *new(map[string]string)
	} else {
		response.Code = common.OK
		response.Data = services.KM(request.TeacherCourseRelationShip)
	}
	c.JSON(http.StatusOK, response)
	return
}

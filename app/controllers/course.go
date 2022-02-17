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
	fmt.Println("访问到controller")
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
		Map_teacher = request.TeacherCourseRelationShip
		KM()
		c.JSON(http.StatusOK, gin.H{"Code": 0, "Data": p})
	}
	c.JSON(http.StatusOK, response)
	return
}

var Map_teacher = map[string][]string{}
var Map_course = map[string][]string{}
var vis = map[string]bool{} //记录课程是否已被访问过
var p = map[string]string{} //记录当前课程被哪位老师选中

func match(teacherID string) bool {
	for _, courseID := range Map_teacher[teacherID] {
		if !vis[courseID] { //有边且未访问
			vis[courseID] = true //记录状态未访问过
			_, ok := p[courseID]
			if !ok || match(p[courseID]) { //如果暂无匹配，或者原来匹配的左侧元素可以找到新的匹配
				p[courseID] = teacherID //当前左侧元素成为当前右侧元素的新匹配
				return true             //返回匹配成功
			}
		}
	}
	return false
}

func KM() {
	for teacherID, courseIDs := range Map_teacher {
		for _, courseID := range courseIDs {
			Map_course[teacherID] = append(Map_course[teacherID], courseID)
		}
	}
	for teacherID, _ := range Map_teacher {
		match(teacherID)
	}
}

package controllers

import (
	"course-system/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Map_teacher = map[string][]string{}
var Map_course = map[string][]string{}
var vis = map[string]bool{} //记录课程是否已被访问过
var p = map[string]string{} //记录当前课程被哪位老师选中

func ScheduleCourse(c *gin.Context) {
	//fmt.Println("访问到该controller了")                 // 不需要 or 后期合并时，注释掉
	var request common.ScheduleCourseRequest           // 声明待绑定的输入数据
	if err := c.ShouldBindJSON(&request); err != nil { // 入参绑定错误，返回错误
		c.JSON(http.StatusOK, common.ScheduleCourseResponse{
			Code: common.ParamInvalid,
		})
		return
	}
	Map_teacher = request.TeacherCourseRelationShip
	KM()
	c.JSON(http.StatusOK, gin.H{"Code": 0, "Data": p})
	return
}

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

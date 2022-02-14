package dao

import (
	"course-system/app/models"
	"course-system/global"
	"fmt"
)

type userCourseDao struct {
}

var UserCourseDao = new(userCourseDao)

func (userCourseDao *userCourseDao) InsertUserCourse(usercourse models.UserCourse) (models.UserCourse, error) {
	err := global.App.DB.Create(&usercourse).Error
	if err != nil { // 如果出错了，打印输出情况。调试时使用，之后注释掉
		fmt.Println(err)
	} else {
		fmt.Println(usercourse) // 打印结果
	}
	return usercourse, err
}

func (userCourseDao *userCourseDao) GetUserCourseList(userID int64) (usercourses []models.UserCourse, err error) {
	err = global.App.DB.Where("ID=?", userID).Find(&usercourses).Error
	return usercourses, err
}

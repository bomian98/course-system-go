package dao

import (
	"course-system/app/models"
	"course-system/global"
)

type userCourseDao struct {
}

func (userCourseDao *userCourseDao) InsertUserCourse(usercourse models.UserCourse) (res_usercourse models.UserCourse, err error) {
	err = global.App.DB.Create(&usercourse).Error
	return res_usercourse, err
}

func (userCourseDao *userCourseDao) GetUserCourseList(userID uint) (usercourses []models.UserCourse, err error) {
	err = global.App.DB.Where("ID=?", userID).Find(&usercourses).Error
	return usercourses, err
}

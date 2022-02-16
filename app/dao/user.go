package dao

import (
	"course-system/app/models"
	"course-system/global"
)

type userDao struct {
}

var UserDao = new(userDao)

// 将寻找到的数据绑定到 user 中，同时返回错误信息
func (userDao *userDao) GetUser(userID int64) (user models.User, err error) {
	err = global.App.DB.Where("ID=?", userID).First(&user).Error
	return
}

func (userDao *userDao) GetUserByUsername(username string) (user *models.User, err error) {
	err = global.App.DB.Where("username=?", username).First(&user).Error
	return
}

package dao

import (
	"course-system/app/models"
	"course-system/global"
)

type userDao struct {
}

var UserDao = new(userDao)

func (userDao *userDao) GetAllStuList() (users []*models.User) {
	global.App.DB.Where("user_type = ?", 2).Find(&users)
	return
}

func (userDao *userDao) GetUserByUsername(username string) (user *models.User, err error) {
	err = global.App.DB.Where("username = ?", username).Find(&user).Error
	return
}

func (userDao *userDao) GetUserByUsername2(username string) (user *models.User, err error) {
	err = global.App.DB.Unscoped().Where("username = ?", username).Find(&user).Error
	return
}

func (userDao *userDao) CreateUser(user *models.User) (err error) {
	err = global.App.DB.Create(&user).Error
	return
}

func (userDao *userDao) UpdateUser(user models.User, id int64) (err error) {
	err = global.App.DB.Where("ID = ?", id).Updates(&user).Error
	return
}

func (userDao *userDao) GetUserByID2(id int64) (user *models.User, err error) {
	err = global.App.DB.Unscoped().Find(&user, "ID = ?", id).Error
	return
}

func (userDao *userDao) GetUserByID(id int64) (user *models.User, err error) {
	err = global.App.DB.Find(&user, "ID = ?", id).Error
	return
}

func (userDao *userDao) DeleteUser(id int64) (err error) {
	err = global.App.DB.Where("ID = ?", id).Delete(&models.User{}).Error
	return
}

func (userDao *userDao) GetUsers(Offset int, Limit int) (users []*models.User, err error) {
	err = global.App.DB.Limit(Limit).Offset(Offset).Find(&users).Error
	return
}

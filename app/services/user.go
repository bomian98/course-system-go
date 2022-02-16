package services

import (
	"System/app/common"
	"System/app/models"
	"System/global"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

func CreateUseServices(request common.CreateMemberRequest) (common.ErrNo, string) {
	var result = global.App.DB.Unscoped().Where("username = ?", request.Username).Find(&models.User{})
	if err := result.Error; err != nil {
		return common.UnknownError, ""
	}
	if result.RowsAffected != 0 { //用户名已存在
		return common.UserHasExisted, ""
	}
	user := models.User{Username: request.Username, Nickname: request.Nickname, Password: request.Password, UserType: request.UserType}
	if err := global.App.DB.Create(&user).Error; err != nil {
		return common.UnknownError, ""
	}
	return common.OK, strconv.FormatInt(user.ID.ID, 10)
}

func UpdateServices(request common.UpdateMemberRequest) common.ErrNo {
	user := models.User{Nickname: request.Nickname}
	id, _ := strconv.ParseInt(request.UserID, 10, 64)
	if code := userStatus(id); code != common.OK {
		return code
	}
	if err := global.App.DB.Where("ID = ?", id).Updates(&user).Error; err != nil {
		return common.UnknownError
	}
	return common.OK
}

func DeleteServices(request common.DeleteMemberRequest) common.ErrNo {
	id, _ := strconv.ParseInt(request.UserID, 10, 64)
	if code := userStatus(id); code != common.OK {
		return code
	}
	if err := global.App.DB.Where("ID = ?", id).Delete(&models.User{}).Error; err != nil {
		return common.UnknownError
	}
	return common.OK
}

func GetServices(request common.GetMemberRequest) (common.ErrNo, models.User) {
	id, _ := strconv.ParseInt(request.UserID, 10, 64)
	user := models.User{}
	if code := userStatus(id); code != common.OK {
		return code, user
	}
	if err := global.App.DB.First(&user, "ID = ?", id).Error; err != nil {
		return common.UnknownError, user
	}
	return common.OK, user
}

func GetsServices(request common.GetMemberListRequest) (common.ErrNo, []models.User) {
	var users []models.User
	if err := global.App.DB.Limit(int(request.Limit)).Offset(int(request.Offset)).Find(&users).Error; err != nil {
		return common.UnknownError, users
	}
	fmt.Println(users)
	return common.OK, users
}

func userStatus(ID int64) common.ErrNo { //用户是否不存在,是否已删除
	var result *gorm.DB
	user := new(models.User)
	result = global.App.DB.Unscoped().Find(&user, "ID = ?", ID)
	if err := result.Error; err != nil {
		return common.UnknownError
	}
	if result.RowsAffected == 0 {
		return common.UserNotExisted
	}
	if user.DeletedAt.Valid == true {
		return common.UserHasDeleted
	}
	return common.OK
}

func GetUserType(ID int64) (common.ErrNo, common.UserType) {
	var result *gorm.DB
	user := new(models.User)
	result = global.App.DB.Find(&user, "ID = ?", ID)
	if err := result.Error; err != nil {
		return common.UnknownError, 0
	} else {
		return common.OK, user.UserType
	}

}

package services

import (
	"System/app/common"
	"System/app/dao"
	"System/app/models"
	"strconv"
)

func CreateUseServices(request common.CreateMemberRequest) (common.ErrNo, string) {
	temp, err := dao.UserDao.GetUserByUsername(request.Username)
	if err != nil {
		return common.UnknownError, ""
	}
	if temp.Username != "" {
		return common.UserHasExisted, ""
	}
	user := models.User{Username: request.Username, Nickname: request.Nickname, Password: request.Password, UserType: request.UserType}
	if err := dao.UserDao.CreateUser(user); err != nil {
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
	if err := dao.UserDao.UpdateUser(user, id); err != nil {
		return common.UnknownError
	}
	return common.OK
}

func DeleteServices(request common.DeleteMemberRequest) common.ErrNo {
	id, _ := strconv.ParseInt(request.UserID, 10, 64)
	if code := userStatus(id); code != common.OK {
		return code
	}
	if err := dao.UserDao.DeleteUser(id); err != nil {
		return common.UnknownError
	}
	return common.OK
}

func GetServices(request common.GetMemberRequest) (common.ErrNo, models.User) {
	id, _ := strconv.ParseInt(request.UserID, 10, 64)
	if code := userStatus(id); code != common.OK {
		return code, models.User{}
	}
	user, err := dao.UserDao.GetUserByID(id)
	if err != nil {
		return common.UnknownError, user
	}
	return common.OK, user
}

func GetsServices(request common.GetMemberListRequest) (common.ErrNo, []models.User) {
	var users []models.User
	users, err := dao.UserDao.GetUsers(int(request.Offset), int(request.Limit))
	if err != nil {
		return common.UnknownError, users
	}
	return common.OK, users
}

func userStatus(ID int64) common.ErrNo { //用户是否不存在,是否已删除
	user, err := dao.UserDao.GetUserByID2(ID)
	if err != nil {
		return common.UnknownError
	}
	if user.Username == "" {
		return common.UserNotExisted
	}
	if user.DeletedAt.Valid == true {
		return common.UserHasDeleted
	}
	return common.OK
}

func GetUserType(ID int64) (common.ErrNo, common.UserType) {
	if code := userStatus(ID); code != common.OK {
		return code, 0
	}
	user, err := dao.UserDao.GetUserByID(ID)
	if err != nil {
		return common.UnknownError, 0
	} else {
		return common.OK, user.UserType
	}

}

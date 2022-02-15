package services

import (
	"System/app/common"
	"System/app/models"
	"System/global"
	"fmt"
	"reflect"
	"strconv"
)

func CreateUseServices(request common.CreateMemberRequest) (common.ErrNo, string) {
	var result = global.App.DB.Where("username = ?", request.Username).First(&models.User{})
	if result.RowsAffected != 0 {
		return common.UserHasExisted, ""
	}
	user := models.User{Username: request.Username, Nickname: request.Nickname, Password: request.Password, UserType: request.UserType}
	fmt.Println("添加", user)
	if err := global.App.DB.Create(&user).Error; err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(user.ID))
	return common.OK, strconv.FormatInt(user.ID.ID, 10)
}

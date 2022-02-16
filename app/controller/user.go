package controller

import (
	"System/app/common"
	"System/app/services"
	"System/app/validParam"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var json common.CreateMemberRequest
	var res common.CreateMemberResponse
	if err := c.ShouldBindJSON(&json); err != nil {
		//绑定错误
		panic(err)
	}
	if validParam.CreateUserValidParam(json) == false {
		//参数错误
		res.Code = common.ParamInvalid
		c.JSON(http.StatusOK, res)
	} else {
		code, userID := services.CreateUseServices(json)
		res = common.CreateMemberResponse{Code: code, Data: struct{ UserID string }{userID}}
		fmt.Println("create success!")
		c.JSON(http.StatusOK, res)
	}
}

func UpdateUser(c *gin.Context) {
	var json common.UpdateMemberRequest
	var res common.UpdateMemberResponse
	if err := c.ShouldBindJSON(&json); err != nil {
		//绑定错误
		panic(err)
	}
	if validParam.UpdateUserValidParam(json) == false {
		//参数错误
		res.Code = common.ParamInvalid
		c.JSON(http.StatusOK, res)
	} else {
		res.Code = services.UpdateServices(json)
		fmt.Println("update success!")
		c.JSON(http.StatusOK, res)
	}
}

func DeleteUser(c *gin.Context) {
	var json common.DeleteMemberRequest
	var res common.DeleteMemberResponse
	if err := c.ShouldBindJSON(&json); err != nil {
		//绑定错误
		panic(err)
	} else {
		res.Code = services.DeleteServices(json)
		fmt.Println("delete success!")
		c.JSON(http.StatusOK, res)
	}
}

func GetUser(c *gin.Context) {
	var request common.GetMemberRequest
	var res common.GetMemberResponse
	if err := c.ShouldBind(&request); err != nil {
		//绑定错误
		panic(err)
	} else {
		code, user := services.GetServices(request)
		member := common.TMember{UserID: strconv.FormatInt(user.ID.ID, 10), Nickname: user.Nickname, Username: user.Username, UserType: user.UserType}
		res = common.GetMemberResponse{Code: code, Data: member}
		fmt.Println("get success!")
		c.JSON(http.StatusOK, res)
	}
}

func GetsUser(c *gin.Context) {
	var request common.GetMemberListRequest
	//var res common.GetMemberListResponse
	if err := c.ShouldBind(&request); err != nil {
		//绑定错误
		panic(err)
	} else {
		fmt.Println(request)
		services.GetsServices(request)

		//code, user := services.GetsServices(json)
		//member := common.TMember{UserID: strconv.FormatInt(user.ID.ID, 10), Nickname: user.Nickname, Username: user.Username, UserType: user.UserType}
		//test := common.GetMemberListResponse{Code: code, Data: member}
		//fmt.Println("get success!")
		//c.JSON(http.StatusOK,gin.H{code,user})
	}
}

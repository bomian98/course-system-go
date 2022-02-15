package controller

import (
	"System/app/common"
	"System/app/services"
	"System/app/validParam"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var json common.CreateMemberRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		//绑定错误
		panic(err)
	}
	fmt.Println("----", json)
	if validParam.CreateUserValidParam(json) == false {
		//参数错误
		c.JSON(http.StatusOK, gin.H{"code": "错误"})
	} else {
		code, userID := services.CreateUseServices(json)
		test := common.CreateMemberResponse{Code: code, Data: struct{ UserID string }{userID}}
		c.JSON(http.StatusOK, test)
	}
}

func UpdateUser(c *gin.Context) {
	var json common.UpdateMemberRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		//绑定错误
		panic(err)
	}
	if validParam.UpdateUserValidParam(json) == false {
		//参数错误
		c.JSON(http.StatusOK, gin.H{"code": "错误"})
	} else {
		code := services.UpdateServices(json)
		test := common.UpdateMemberResponse{Code: code}
		fmt.Println("update success!")
		c.JSON(http.StatusOK, test)
	}
}

func DeleteUser(c *gin.Context) {
	var json common.DeleteMemberRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		//绑定错误
		panic(err)
	} else {
		code := services.DeleteServices(json)
		test := common.UpdateMemberResponse{Code: code}
		fmt.Println("delete success!")
		c.JSON(http.StatusOK, test)
	}
}

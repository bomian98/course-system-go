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

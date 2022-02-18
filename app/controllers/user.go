package controllers

import (
	"course-system/app/common"
	"course-system/app/services"
	"course-system/app/validParam"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var json common.CreateMemberRequest
	var res common.CreateMemberResponse
	if err := c.ShouldBindJSON(&json); err != nil {
		//绑定错误
		res.Code = common.UnknownError
		c.JSON(http.StatusOK, res)
	} else if validParam.CreateUserValidParam(json) == false {
		//参数错误
		res.Code = common.ParamInvalid
		c.JSON(http.StatusOK, res)
	} else {
		code, userID := services.CreateUseServices(json)
		res = common.CreateMemberResponse{Code: code, Data: struct{ UserID string }{userID}}
		log.Println("create success!")
		c.JSON(http.StatusOK, res)
	}
}

func UpdateUser(c *gin.Context) {
	var json common.UpdateMemberRequest
	var res common.UpdateMemberResponse
	if err := c.ShouldBindJSON(&json); err != nil {
		//绑定错误
		res.Code = common.UnknownError
		c.JSON(http.StatusOK, res)
	} else if validParam.UpdateUserValidParam(json) == false {
		//参数错误
		res.Code = common.ParamInvalid
		c.JSON(http.StatusOK, res)
	} else {
		res.Code = services.UpdateServices(json)
		log.Println("update success!")
		c.JSON(http.StatusOK, res)
	}
}

func DeleteUser(c *gin.Context) {
	var json common.DeleteMemberRequest
	var res common.DeleteMemberResponse
	if err := c.ShouldBindJSON(&json); err != nil {
		//绑定错误
		res.Code = common.UnknownError
		c.JSON(http.StatusOK, res)
	} else {
		res.Code = services.DeleteServices(json)
		log.Println("delete success!")
		c.JSON(http.StatusOK, res)
	}
}

func GetUser(c *gin.Context) {
	var request common.GetMemberRequest
	var res common.GetMemberResponse
	if err := c.ShouldBind(&request); err != nil {
		//绑定错误
		res.Code = common.UnknownError
		c.JSON(http.StatusOK, res)
	} else {
		code, user := services.GetServices(request)
		res.Code = code
		if code != common.OK {
			c.JSON(http.StatusOK, res)
		} else {
			member := common.TMember{UserID: strconv.FormatInt(user.ID.ID, 10), Nickname: user.Nickname, Username: user.Username, UserType: user.UserType}
			res.Data = member
			log.Println("get success!")
			c.JSON(http.StatusOK, res)
		}

	}
}

func GetsUser(c *gin.Context) {
	var request common.GetMemberListRequest
	var res common.GetMemberListResponse
	if err := c.ShouldBind(&request); err != nil {
		//绑定错误
		res.Code = common.UnknownError
		c.JSON(http.StatusOK, res)
	} else {
		code, users := services.GetsServices(request)
		members := make([]common.TMember, 0)
		for i := 0; i < len(users); i++ {
			members = append(members, common.TMember{UserID: strconv.FormatInt(users[i].ID.ID, 10), Nickname: users[i].Nickname, Username: users[i].Username, UserType: users[i].UserType})
		}
		res.Code = code
		res.Data.MemberList = members
		log.Println("gets success!")
		c.JSON(http.StatusOK, res)
	}
}

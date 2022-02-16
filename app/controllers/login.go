package controllers

import (
	"course-system/app/common"
	"course-system/app/services"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Login(c *gin.Context) {
	fmt.Println("访问到该controller了") // 不需要 or 后期合并时，注释掉
	var request common.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil { // 入参绑定错误，返回错误
		c.JSON(http.StatusBadRequest, common.LoginResponse{
			Code: common.ParamInvalid,
		})
		return
	}
	// 获取用户并检查密码
	user, errno := services.UserService.GetUserByUsername(request.Username)
	if errno != common.OK || strings.Compare(request.Password, user.Password) != 0 {
		// 清除session
		s := sessions.Default(c)
		s.Delete("camp-session")
		if s.Save() != nil {
			c.JSON(http.StatusOK, common.LoginResponse{
				Code: common.UnknownError,
			})
			return
		}
		// 清除cookie
		c.SetCookie("camp-session", "", -1, "/", "", false, true)
		c.JSON(http.StatusOK, common.LoginResponse{
			Code: common.WrongPassword,
		})
		return
	} else {
		useridStr := strconv.FormatInt(user.ID.ID, 10)
		var response common.LoginResponse
		response.Code = common.OK
		response.Data.UserID = useridStr
		// 生成token
		token := services.UserService.UserMD5(useridStr + strconv.FormatInt(time.Now().Unix(), 10))
		// 设置session
		s := sessions.Default(c)
		s.Set("camp-session", token)
		s.Set("userid", useridStr)
		if s.Save() != nil {
			c.JSON(http.StatusOK, common.LoginResponse{
				Code: common.UnknownError,
			})
			return
		}
		// 设置cookie
		c.SetCookie("camp-session", token, common.CookieAge, "/", "", false, true)
		c.JSON(http.StatusOK, response)
		return
	}
}

func Logout(c *gin.Context) {
	fmt.Println("访问到该controller了") // 不需要 or 后期合并时，注释掉
	var request common.LogoutRequest
	if err := c.ShouldBindJSON(&request); err != nil { // 入参绑定错误，返回错误
		c.JSON(http.StatusBadRequest, common.LogoutResponse{
			Code: common.ParamInvalid,
		})
		return
	}
	// 获取cookie
	if cookie, err := c.Cookie("camp-session"); err != nil { // 未获取到cookie
		c.JSON(http.StatusOK, common.LogoutResponse{
			Code: common.LoginRequired,
		})
		return
	} else {
		// 清除session
		s := sessions.Default(c)
		s.Delete("camp-session")
		if s.Save() != nil {
			c.JSON(http.StatusOK, common.LogoutResponse{
				Code: common.UnknownError,
			})
			return
		}
		// 清除cookie
		c.SetCookie("camp-session", cookie, -1, "/", "", false, true)
		c.JSON(http.StatusOK, common.LogoutResponse{
			Code: common.OK,
		})
		return
	}
}

func WhoAmI(c *gin.Context) {
	fmt.Println("访问到该controller了") // 不需要 or 后期合并时，注释掉
	var request common.WhoAmIRequest
	if err := c.ShouldBindJSON(&request); err != nil { // 入参绑定错误，返回错误
		c.JSON(http.StatusBadRequest, common.WhoAmIResponse{
			Code: common.ParamInvalid,
		})
		return
	}
	// 获取cookie
	if cookie, err := c.Cookie("camp-session"); err != nil { // 未获取到cookie
		// 清除session
		s := sessions.Default(c)
		s.Delete("camp-session")
		if s.Save() != nil {
			c.JSON(http.StatusOK, common.WhoAmIResponse{
				Code: common.UnknownError,
			})
			return
		}
		c.JSON(http.StatusOK, common.WhoAmIResponse{
			Code: common.LoginRequired,
		})
		return
	} else {
		s := sessions.Default(c)
		token := s.Get("camp-session")
		if token == nil {
			c.JSON(http.StatusOK, common.WhoAmIResponse{
				Code: common.UnknownError,
			})
			return
		}
		if strings.Compare(token.(string), cookie) != 0 {
			// 清除session
			s.Delete("camp-session")
			if s.Save() != nil {
				c.JSON(http.StatusOK, common.WhoAmIResponse{
					Code: common.UnknownError,
				})
				return
			}
			c.SetCookie("camp-session", "", -1, "/", "", false, true)
			c.JSON(http.StatusOK, common.WhoAmIResponse{
				Code: common.LoginRequired,
			})
			return
		}
		// 获取TMember
		useridSession := s.Get("userid")
		if useridSession == nil {
			c.JSON(http.StatusOK, common.WhoAmIResponse{
				Code: common.UnknownError,
			})
			return
		}
		if userId, err := strconv.ParseInt(useridSession.(string), 10, 64); err != nil { // cookie有问题，直接清除cookie
			// 清除session
			s.Delete("camp-session")
			if s.Save() != nil {
				c.JSON(http.StatusOK, common.WhoAmIResponse{
					Code: common.UnknownError,
				})
				return
			}
			c.SetCookie("camp-session", "", -1, "/", "", false, true)
			c.JSON(http.StatusOK, common.WhoAmIResponse{
				Code: common.UnknownError,
			})
			return
		} else {
			tMember, errno := services.UserService.GetTMember(userId)
			if errno == common.OK {
				// 刷新cookie?
				// c.SetCookie("camp-session", cookie, common.CookieAge, "/", "", false, true)
				c.JSON(http.StatusOK, common.WhoAmIResponse{
					Code: common.OK,
					Data: tMember,
				})
			} else {
				// userid有问题，清除cookie
				c.SetCookie("camp-session", cookie, -1, "/", "", false, true)
				c.JSON(http.StatusOK, common.WhoAmIResponse{
					Code: common.LoginRequired,
				})
			}
			return
		}
	}
}

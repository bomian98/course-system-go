package middleware

import (
	"System/app/common"
	"System/app/services"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authorize(e *casbin.Enforcer) gin.HandlerFunc {

	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("userid")
		if sessionValue == nil {
			c.JSON(http.StatusOK, common.CreateMemberResponse{Code: common.LoginRequired})
		}
		ID := sessionValue.(int64)
		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		if code, sub := services.GetUserType(ID); code == common.OK {
			//判断策略中是否存在
			if ok := e.Enforce(sub, obj, act); ok {
				fmt.Println("恭喜您,权限验证通过")
				c.Next()
			} else {
				fmt.Println("很遗憾,权限验证没有通过")
				c.JSON(http.StatusOK, common.CreateMemberResponse{Code: common.PermDenied})
			}
		} else {
			c.JSON(http.StatusOK, common.CreateMemberResponse{Code: code})
		}
	}
}

package utils

import (
	"System/app/common"
	"System/app/services"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Authorize(e *casbin.Enforcer) gin.HandlerFunc {

	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("userid")
		if sessionValue != nil {
			fmt.Println("sessionValue:", sessionValue)
			ID := sessionValue.(string)
			id, _ := strconv.ParseInt(ID, 10, 64)
			//获取请求的URI
			obj := c.Request.URL.Path
			//获取请求方法
			act := c.Request.Method
			//获取用户的角色
			fmt.Println(obj, "    ", act)
			if code, temp := services.GetUserType(id); code == common.OK {
				//判断策略中是否存在
				sub := strconv.Itoa(int(temp))
				if ok := e.Enforce(sub, obj, act); ok {
					fmt.Println("恭喜您,权限验证通过")
					c.Next()
				} else {
					fmt.Println("很遗憾,权限验证没有通过")

					c.JSON(http.StatusOK, common.CreateMemberResponse{Code: common.PermDenied})
					c.Abort()
				}
			} else {
				c.JSON(http.StatusOK, common.CreateMemberResponse{Code: code})
			}
		} else {
			c.JSON(http.StatusOK, common.CreateMemberResponse{Code: common.LoginRequired})
		}

	}
}

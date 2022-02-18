package utils

import (
	"course-system/app/common"
	"course-system/app/services"
	"github.com/casbin/casbin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Authorize(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("userid")
		if sessionValue != nil {
			log.Println("sessionValue:", sessionValue)
			ID := sessionValue.(string)
			//获取请求的URI
			obj := c.Request.URL.Path
			//获取请求方法
			act := c.Request.Method
			//获取用户的角色

			if code, sub := services.GetUserType(ID); code == common.OK {
				//判断策略中是否存在
				log.Println(sub, " ", obj, "    ", act)
				if ok := e.Enforce(sub, obj, act); ok {
					log.Println("恭喜您,权限验证通过")
					c.Next()
				} else {
					log.Println("很遗憾,权限验证没有通过")
					c.JSON(http.StatusOK, common.CreateMemberResponse{Code: common.PermDenied})
					c.Abort()
				}
			} else {
				c.JSON(http.StatusOK, common.CreateMemberResponse{Code: code})
				c.Abort()
			}
		} else {
			c.JSON(http.StatusOK, common.CreateMemberResponse{Code: common.LoginRequired})
			c.Abort()
		}

	}
}

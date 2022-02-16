package bootstrap

import (
	"github.com/casbin/casbin"
)

func InitCasbin() *casbin.Enforcer {

	e := casbin.NewEnforcer("./casbin_rbac_model.conf", "./Policy.csv")
	return e
}

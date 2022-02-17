package global

import (
	"course-system/config"
	"github.com/casbin/casbin"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Log         *zap.Logger
	Redis       *redis.Client
	DB          *gorm.DB
	E           *casbin.Enforcer
}

var App = new(Application)

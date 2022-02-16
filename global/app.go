package global

import (
	"System/config"
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
	DB          *gorm.DB
	E           *casbin.Enforcer
	Redis       *redis.Client
}

var App = new(Application)

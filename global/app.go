package global

import (
	"course-system/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Redis       *redis.Client
	DB          *gorm.DB
}

var App = new(Application)

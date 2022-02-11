package global

import (
	"course-system/config"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Redis       *redis.Client
}

var App = new(Application)

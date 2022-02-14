package bootstrap

import (
	"course-system/global"
	"fmt"
	"github.com/go-redis/redis"
)

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.App.Config.Redis.Host + ":" + global.App.Config.Redis.Port,
		Password: global.App.Config.Redis.Password, // no password set
		DB:       global.App.Config.Redis.DB,       // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Redis connect ping failed, err:", err)
		return nil
	} else {
		fmt.Println("Redis 加载成功")
	}
	return client
}

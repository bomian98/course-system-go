package bootstrap

import (
	"System/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.App.Config.Redis.Host + ":" + global.App.Config.Redis.Port,
		Password: global.App.Config.Redis.Password, // no password set
		DB:       global.App.Config.Redis.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Redis connect ping failed, err:", err)
		return nil
	} else {
		fmt.Println("Redis 加载成功")
	}
	return client
}

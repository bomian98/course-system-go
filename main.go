package main

import (
	"course-system/bootstrap"
	"course-system/global"
	"log"
)

func main() {

	// 初始化配置
	bootstrap.InitializeConfig()

	bootstrap.InitializeLog()

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 初始化Redis
	global.App.Redis = bootstrap.InitializeRedis()

	// 开启抢课管道消费者
	bootstrap.ConsumerOpen()

	log.Println("1")
	// 启动服务器
	bootstrap.RunServer()

}

package main

import (
	"course-system/bootstrap"
	"course-system/global"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 配置路由
	bootstrap.RunServer()

}

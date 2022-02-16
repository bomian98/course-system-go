package main

import (
	"System/bootstrap"
	"System/global"
)

func main() {
	//初始化配置
	bootstrap.InitializeConfig()

	//初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()
	//初始化casbin
	global.App.E = bootstrap.InitCasbin()
	// 路由
	bootstrap.RunServer()

}

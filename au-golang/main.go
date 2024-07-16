package main

import (
	"au-go/core"
	"au-go/global"
	"au-go/initialize"
)

func main() {
	//初始化全局变量
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库

	if global.GVA_DB != nil {
		initialize.MysqlTables(global.GVA_DB) // 初
		// 始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

	core.RunWindowsServer()
}

package main

import (
	"au-golang/core"
	"au-golang/global"
	"au-golang/initialize"
	"au-golang/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//设置为发行模式
	gin.SetMode(gin.ReleaseMode)

	global.GVA_VP = core.Viper() //初始化Viper
	//global.GVA_LOG = core.Zap()
	global.GVA_DB = initialize.Gorm()
	if global.GVA_DB != nil {
		//initialize.MysqlTables(global.GVA_DB)
		// 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	// 小程序access token定时刷新
	if err := utils.MPTokenRefresh(); err != nil {
		fmt.Printf("小程序access获取失败 %v \n", err.Error())

		return
	}
	// 从excel中导入社联报名数据
	//utils.AddInterviewFromExcel()

	//从excel中导入社团报名数据
	//utils.AddInterviewUsers()
	core.RunWindowsServer()
}

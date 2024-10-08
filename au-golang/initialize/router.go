package initialize

import (
	"au-go/global"
	"au-go/middleware"
	"au-go/router"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.GVA_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	//PublicGroup := Router.Group("")
	//{
	//	//router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	//	//router.InitInitRouter(PublicGroup) // 自动初始化相关
	//}
	PrivateGroup := Router.Group("")
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	PrivateGroup.Use()
	{
		router.InitUserRouter(PrivateGroup)
	}
	global.GVA_LOG.Info("router register success")
	return Router
}

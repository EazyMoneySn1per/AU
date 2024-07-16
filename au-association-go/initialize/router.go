package initialize

import (
	"au-golang/middleware"
	"au-golang/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middleware.Cors())
	systemRouter := router.RouterGroupApp.System
	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 静态页面设置
	//Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址

	//PublicGroup := Router.Group("")
	//{
	//	// 健康监测
	//	PublicGroup.GET("/health", func(c *gin.Context) {
	//		c.JSON(200, "ok")
	//	})
	//}
	//{
	//	systemRouter.InitBaseRouter(PublicGroup)
	//}
	PrivateGroup := Router.Group("")
	// 添加中间件
	PrivateGroup.Use()
	{
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitAccountRouter(PrivateGroup)
		systemRouter.InitActivityRouter(PrivateGroup)
		systemRouter.InitWxUserRouter(PrivateGroup)
		systemRouter.InitOutlayRouter(PrivateGroup)
		systemRouter.InitTwitterRouter(PrivateGroup)
		systemRouter.InitSynthesizeRouter(PrivateGroup)
		systemRouter.InitDepartmentRouter(PrivateGroup)
		systemRouter.InitUserMigrateRouter(PrivateGroup)
		systemRouter.InitDataShowRouter(PrivateGroup)
		systemRouter.InitAdminRouter(PrivateGroup)
		systemRouter.InitSysInterviewRouter(PrivateGroup)
		systemRouter.InitWxInterViewRouter(PrivateGroup)
		systemRouter.InitAdminsterRouter(PrivateGroup)
	}
	return Router
}

package system

import (
	v1 "au-golang/api"
	"au-golang/utils"
	"github.com/gin-gonic/gin"
)

type OutlayRouter struct{}

func (s *UserRouter) InitOutlayRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("outlay").Use(utils.JWTAuthMiddleware())
	outlayApi := v1.ApiGroupApp.SystemApiGroup.OutlayApi
	{
		userRouter.POST("outlaysubmit", outlayApi.OutlaySubmit)
		userRouter.GET("downFile", outlayApi.DownFile)
		userRouter.GET("getinfo", outlayApi.GetInfo)
		userRouter.GET("nextstep", outlayApi.Nextstep)
		userRouter.GET("setbackmsg", outlayApi.SetBackMsg)
		userRouter.GET("getOutlays", outlayApi.GetActivities)
		userRouter.GET("getRecord", outlayApi.GetRecord)
	}
}

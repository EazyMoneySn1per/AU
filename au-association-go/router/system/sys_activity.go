package system

import (
	v1 "au-golang/api"
	"au-golang/utils"
	"github.com/gin-gonic/gin"
)

type ActivityRouter struct{}

func (s *ActivityRouter) InitActivityRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("activity").Use(utils.JWTAuthMiddleware())
	activityApi := v1.ApiGroupApp.SystemApiGroup.ActivityApi
	{
		userRouter.POST("submit", activityApi.Submit)
		userRouter.GET("getinfo", activityApi.Getinfo)
		userRouter.GET("nextstep", activityApi.Nextstep)
		userRouter.GET("setbackmsg", activityApi.SetBackMsg)
		userRouter.GET("getactivities", activityApi.GetActivities)
		userRouter.GET("downFile", activityApi.DownFile)
		userRouter.GET("getRecord", activityApi.GetRecord)
	}

}

package system

import (
	v1 "au-golang/api"
	"au-golang/utils"
	"github.com/gin-gonic/gin"
)

type TwitterRouter struct{}

func (s *UserRouter) InitTwitterRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("twitter").Use(utils.JWTAuthMiddleware())
	twitterApi := v1.ApiGroupApp.SystemApiGroup.TwitterApi
	{
		userRouter.POST("twittersubmit", twitterApi.TwitterSubmit)
		userRouter.GET("downFile", twitterApi.DownFile)
		userRouter.GET("getinfo", twitterApi.GetInfo)
		userRouter.GET("nextstep", twitterApi.Nextstep)
		userRouter.GET("setbackmsg", twitterApi.SetBackMsg)
		userRouter.GET("gettwitters", twitterApi.GetActivities)
		userRouter.GET("getRecord", twitterApi.GetRecord)
	}
}

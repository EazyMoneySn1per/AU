package system

import (
	v1 "au-golang/api"
	"au-golang/utils"
	"github.com/gin-gonic/gin"
)

type DataShowRouter struct{}

func (s *UserRouter) InitDataShowRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("dataShow").Use(utils.JWTAuthMiddleware())
	dataShowApi := v1.ApiGroupApp.SystemApiGroup.DataShowApi
	{
		userRouter.GET("getUnfinishedEvent", dataShowApi.GetUnfinishedEvent)
		userRouter.GET("getFinishedEvent", dataShowApi.GetFinishedEvent)
		userRouter.GET("getTodayEvent", dataShowApi.GetTodayEvent)
	}
}

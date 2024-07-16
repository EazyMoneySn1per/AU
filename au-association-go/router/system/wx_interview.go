package system

import (
	v1 "au-golang/api"
	"au-golang/utils"
	"github.com/gin-gonic/gin"
)

type WxInterViewRouter struct{}

func (s *UserRouter) InitWxInterViewRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("wxapi/interview").Use(utils.JWTAuthMiddleware())
	interViewApi := v1.ApiGroupApp.SystemApiGroup.InterViewApi
	{
		userRouter.GET("getUserInterviews", interViewApi.GetUserInterviews)
		userRouter.GET("studentRefuse", interViewApi.StudentRefuse)
		userRouter.GET("studentConfirm", interViewApi.StudentConfirm)
		userRouter.POST("addInterviewUser", interViewApi.AddInterviewUser)
	}

}

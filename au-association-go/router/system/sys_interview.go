package system

import (
	v1 "au-golang/api"
	"au-golang/utils"
	"github.com/gin-gonic/gin"
)

type SysInterviewRouter struct {
}

func (s *SysInterviewRouter) InitSysInterviewRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("interview").Use(utils.JWTAuthMiddleware())
	sysinterviewApi := v1.ApiGroupApp.SystemApiGroup.SysInterviewApi
	{
		userRouter.GET("getAssInterviews", sysinterviewApi.GetInterviewUsers)
		userRouter.GET("AssConfirm", sysinterviewApi.AssConfirm)
		userRouter.GET("AssRefuse", sysinterviewApi.AssRefuse)
		userRouter.GET("getAssMessage", sysinterviewApi.GetAssMessage)
		userRouter.GET("setAssMessage", sysinterviewApi.SetAssShowMessage)
	}
}

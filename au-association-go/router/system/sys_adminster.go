package system

import (
	v1 "au-golang/api"
	"au-golang/utils"
	"github.com/gin-gonic/gin"
)

type AdminsterRouter struct{}

func (s *UserRouter) InitAdminsterRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("AMinister").Use(utils.JWTAuthMiddleware())
	adminsterApi := v1.ApiGroupApp.SystemApiGroup.AdministerApi
	{
		userRouter.POST("uploadAssLogo", adminsterApi.UploadAssLogo)
		userRouter.POST("setAssDescription", adminsterApi.SetAssDescription)
	}
}

package system

import (
	v1 "au-golang/api"
	"au-golang/utils"
	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (s *UserRouter) InitAdminRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("admin").Use(utils.JWTAuthMiddleware())
	adminApi := v1.ApiGroupApp.SystemApiGroup.AdminApi
	{
		userRouter.GET("getAssociations", adminApi.GetAssociation)
		userRouter.GET("getAssociationsByAssid", adminApi.GetAssociationsByAssid)
		userRouter.POST("setAssociationsByAssid", adminApi.SetAssociationsByAssid)
		userRouter.POST("addAssociations", adminApi.AddAssociations)
		userRouter.GET("deleteAssociations", adminApi.DeleteAssociations)
		userRouter.GET("getAssociationsNameMapAssid", adminApi.GetAssociationsNameMapAssid)
		userRouter.GET("getAssociationsName", adminApi.GetAssociationName)
		userRouter.GET("getAssociationsStudents", adminApi.GetAssociationStudents)
		userRouter.GET("getActivityCompleteList", adminApi.GetActivityCompleteList)
		userRouter.GET("getStudentsList", adminApi.GetStudentsList)
		userRouter.GET("getTwitterList", adminApi.GetTwitter)
		userRouter.POST("setActivityTwitter", adminApi.SetActivityTwitter)
		userRouter.GET("getOutlayCompleteList", adminApi.GetOutlayCompleteList)
		userRouter.POST("setMiniProgramModule", adminApi.SetMiniProgramModule)
		userRouter.GET("getMiniProgramModule", adminApi.GetMiniProgramModule)
	}
}

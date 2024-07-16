package migrate

import (
	v1 "au-golang/api"
	"au-golang/utils"
	"github.com/gin-gonic/gin"
)

type UserRouterMigrate struct{}

func (s *UserRouterMigrate) InitUserMigrateRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("wxapi/interview/").Use(utils.JWTAuthMiddleware())
	accountApi := v1.ApiGroupApp.SystemApiGroup.UserApiMigrate
	{
		userRouter.POST("goPost/AddInterviewUser", accountApi.AddInterviewUser)
		userRouter.GET("goGet/FindStudentInterviewLists", accountApi.FindStudentInterviewLists)
		// 后台接口
		userRouter.POST("goPost/SetInterviewUserStatusInBatch", accountApi.SetInterviewUserStatusInBatch)
		userRouter.GET("goGet/SetInterviewUerStatusFailed", accountApi.SetInterviewUerStatusFailed)
		userRouter.GET("goGet/SetInterviewUerStatusSuccess", accountApi.SetInterviewUerStatusSuccess)
		userRouter.GET("goGet/FindDepartmentUsers", accountApi.FindDepartmentUsers)
		userRouter.GET("goGet/FindDepartmentAllUsers", accountApi.FindDepartmentAllUsers)
		userRouter.POST("goPost/ExportExcelSetInterviewTime", accountApi.ExportExcelSetInterviewTime)
		userRouter.POST("goPost/SetInterviewTime", accountApi.SetInterviewTime)
	}
}

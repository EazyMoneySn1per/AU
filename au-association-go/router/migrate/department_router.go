package migrate

import (
	v1 "au-golang/api"
	"au-golang/utils"
	"github.com/gin-gonic/gin"
)

type DepartmentRouterMigrate struct{}

func (s *UserRouterMigrate) InitDepartmentRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("wxapi/interview/").Use(utils.JWTAuthMiddleware())
	accountApi := v1.ApiGroupApp.SystemApiGroup.DepartmentApiMigrate
	{
		userRouter.POST("goPost/SetDepartmentMessage", accountApi.SetDepartmentMessage)
		userRouter.GET("goGet/FindDepartmentMessage", accountApi.FindDepartmentMessage)
		userRouter.POST("goPost/SetWxUserAssInBatch", accountApi.SetWxUserAssInBatch)
		userRouter.GET("goGet/GetTotalJoinAssStudent", accountApi.GetTotalJoinAssStudent)
	}
}

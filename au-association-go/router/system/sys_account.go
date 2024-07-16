package system

import (
	v1 "au-golang/api"
	"au-golang/utils"
	"github.com/gin-gonic/gin"
)

type AccountRouter struct{}

func (s *UserRouter) InitAccountRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("account").Use(utils.JWTAuthMiddleware())
	accountApi := v1.ApiGroupApp.SystemApiGroup.AccountApi
	{
		userRouter.GET("deleteUser", accountApi.DeleteUser)
		userRouter.POST("updateInfo", accountApi.UpdateInfo)
		userRouter.GET("getLists", accountApi.GetLists)
	}
}

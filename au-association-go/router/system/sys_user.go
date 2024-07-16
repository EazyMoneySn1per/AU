package system

import (
	v1 "au-golang/api"
	"au-golang/utils"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("auth").Use(utils.JWTAuthMiddleware())
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("logout", userApi.Logout)
		userRouter.GET("info", userApi.GetInfo)
		userRouter.GET("exportUsersList", userApi.ExportUsersList)
	}
	//登录不需要中间件
	loginRouter := Router.Group("auth").Use()
	{
		loginRouter.POST("login", userApi.Login)
	}
}

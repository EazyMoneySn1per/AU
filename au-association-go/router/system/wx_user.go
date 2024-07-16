package system

import (
	v1 "au-golang/api"
	"au-golang/utils"
	"github.com/gin-gonic/gin"
)

type WxUserRouter struct {
}

func (w *WxUserRouter) InitWxUserRouter(Router *gin.RouterGroup) {

	userRouter := Router.Group("wxapi").Use(utils.JWTAuthMiddleware())
	wxUserApi := v1.ApiGroupApp.SystemApiGroup.WxUserApi
	{
		userRouter.GET("getTotalAssociation", wxUserApi.GetTotalAssociation)
		userRouter.GET("getExecllentAssociation", wxUserApi.GetExecllentAssociation)
		userRouter.GET("test", wxUserApi.Test)
		userRouter.GET("getAssociationsNameMapAssid", wxUserApi.GetAssociationsNameMapAssid)
		userRouter.GET("getAssociationsByAssid", wxUserApi.GetAssociationsByAssid)
		userRouter.GET("getTwitterPic", wxUserApi.GetTwitterPic)
		userRouter.GET("getAssByType", wxUserApi.GetAssByType)
		userRouter.GET("exitAss", wxUserApi.ExitAss)
		userRouter.GET("searchStudent", wxUserApi.SearchStudent)
	}
	//登录和认证不需要token校验
	loginRouter := Router.Group("wxapi")
	{
		loginRouter.GET("getUserInfo", wxUserApi.GetUserInfo)
		loginRouter.POST("addWxUser", wxUserApi.AddWxUser)
	}
}

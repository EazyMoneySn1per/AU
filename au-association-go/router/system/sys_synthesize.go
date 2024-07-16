package system

import (
	v1 "au-golang/api"
	"au-golang/utils"
	"github.com/gin-gonic/gin"
)

type SynthesizeRouter struct{}

func (s *UserRouter) InitSynthesizeRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("synthesize").Use(utils.JWTAuthMiddleware())
	synthesizeApi := v1.ApiGroupApp.SystemApiGroup.SynthesizeApi
	{
		userRouter.POST("synthesizesubmit", synthesizeApi.SynthesizeSubmit)
		userRouter.GET("synthesizegetlist", synthesizeApi.GetList)
	}
}

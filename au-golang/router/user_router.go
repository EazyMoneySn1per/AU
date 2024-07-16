package router

import (
	"au-go/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("")
	{
		UserRouter.POST("test", api.Test)
		UserRouter.POST("AddInterviewUser", api.AddInterviewUser)
		UserRouter.GET("FindStudentInterviewLists", api.FindStudentInterviewLists)
		// 后台接口
		UserRouter.POST("SetInterviewUserStatusInBatch", api.SetInterviewUserStatusInBatch)
		UserRouter.GET("SetInterviewUerStatusFailed", api.SetInterviewUerStatusFailed)
		UserRouter.GET("SetInterviewUerStatusSuccess", api.SetInterviewUerStatusSuccess)
		UserRouter.GET("FindDepartmentUsers", api.FindDepartmentUsers)
		UserRouter.GET("FindDepartmentAllUsers", api.FindDepartmentAllUsers)
		UserRouter.POST("ExportExcelSetInterviewTime", api.ExportExcelSetInterviewTime)
		UserRouter.POST("SetInterviewTime", api.SetInterviewTime)
		UserRouter.POST("SetDepartmentMessage", api.SetDepartmentMessage)
		UserRouter.GET("FindDepartmentMessage", api.FindDepartmentMessage)
		UserRouter.POST("SetWxUserAssInBatch", api.SetWxUserAssInBatch)
		UserRouter.GET("GetTotalJoinAssStudent", api.GetTotalJoinAssStudent)
	}
}

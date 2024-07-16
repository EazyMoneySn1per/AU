package api

import (
	"au-go/model/request"
	"au-go/model/response"
	"au-go/service"
	"github.com/gin-gonic/gin"
)

/**
设置部门的消息
*/
func SetDepartmentMessage(c *gin.Context) {
	var r request.DepartmentMessage
	_ = c.ShouldBindJSON(&r)
	department := service.FindDepartment(r.DepartmentName)
	department.StageOneMessage = r.StageOneMessage
	department.StageTwoMessage = r.StageTwoMessage
	department.StageOneFailedMessage = r.StageOneFailedMessage
	department.StageTwoSuccessMessage = r.StageTwoSuccessMessage
	department.StageTwoFailedMessage = r.StageTwoFailedMessage
	service.SetDepartmentMessage(department)
	response.OkWithMessage("设置成功", c)
}

/**
查找部门的消息
*/
func FindDepartmentMessage(c *gin.Context) {
	departmentName := c.Query("department")
	department := service.FindDepartment(departmentName)
	response.OkWithDetailed(department, "查找成功", c)
}

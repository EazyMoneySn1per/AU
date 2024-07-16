package api

import (
	"au-golang/model/migrate/request"
	"au-golang/model/migrate/response"
	service "au-golang/service/migrate"
	"github.com/gin-gonic/gin"
)

type DepartmentApiMigrate struct{}

/**
设置部门的消息
*/
func (departmentApiMigrate *DepartmentApiMigrate) SetDepartmentMessage(c *gin.Context) {
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
func (departmentApiMigrate *DepartmentApiMigrate) FindDepartmentMessage(c *gin.Context) {
	departmentName := c.Query("department")
	department := service.FindDepartment(departmentName)
	response.OkWithDetailed(department, "查找成功", c)
}

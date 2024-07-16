package api

import (
	"au-golang/model/migrate/request"
	service "au-golang/service/migrate"
	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {
	var r request.DepartmentAndStatusNoPage
	c.ShouldBindJSON(&r)
	users, _ := service.FindDepartmentAllUsers(r.Department, r.Status)
	var studentIds []string
	for _, u := range users {
		studentIds = append(studentIds, u.StudentId)
	}
	//wxUsers := service.FindWxUserByStudentId(studentIds)
	//for _,u := range wxUsers {
	//	//u.OpenId
	//}
}

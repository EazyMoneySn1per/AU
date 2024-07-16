package service

import (
	"au-go/global"
	"au-go/model"
)

func FindDepartment(name string) model.AuInterviewDepartment {
	var department model.AuInterviewDepartment
	global.GVA_DB.Where("department_name = ?", name).Find(&department)
	return department
}
func FindDepartmentById(id uint) model.AuInterviewDepartment {
	var department model.AuInterviewDepartment
	global.GVA_DB.Where("id = ?", id).Find(&department)
	return department
}

//func(db *gorm.DB) *gorm.DB {
//	return db.Order("au_interview_time_tables.month ASC")
//}
func FindDepartmentUsers(name string, status string, page int, limit int) ([]model.AuInterviewUser, int64) {
	department := FindDepartment(name)
	var users []model.AuInterviewUser
	var count int64
	//global.GVA_DB.Where("au_interview_department_id = ? AND status = ?", department.ID,status).Preload("AuInterviewTimeTable").Limit(limit).Offset((page - 1) * limit).Find(&users).Count(&count)
	global.GVA_DB.Preload("AuInterviewTimeTable").Where("au_interview_department_id = ? AND status = ?", department.ID, status).Limit(limit).Offset((page - 1) * limit).
		Find(&users).Offset(-1).Limit(-1).Count(&count)
	return users, count
}
func FindDepartmentAllUsers(name string, status string) ([]model.AuInterviewUser, int64) {
	department := FindDepartment(name)
	var users []model.AuInterviewUser
	var count int64
	global.GVA_DB.Where("au_interview_department_id = ? AND status = ?", department.ID, status).Preload("AuInterviewTimeTable").Find(&users).Count(&count)
	return users, count
}
func AllUsers(status string) ([]model.AuInterviewUser, int64) {
	var users []model.AuInterviewUser
	var count int64
	global.GVA_DB.Where("status = ?", status).Preload("AuInterviewTimeTable").Find(&users).Count(&count)
	return users, count
}
func SetDepartmentMessage(department model.AuInterviewDepartment) {
	global.GVA_DB.Save(department)
}

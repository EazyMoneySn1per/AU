package service

import (
	"au-golang/global"
	model "au-golang/model/migrate"
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

//	func(db *gorm.DB) *gorm.DB {
//		return db.Order("au_interview_time_tables.month ASC")
//	}
func FindDepartmentUsers(name string, status string, query string, page int, limit int, isId bool) ([]model.AuInterviewUser, int64) {
	department := FindDepartment(name)
	var users []model.AuInterviewUser
	db := global.GVA_DB.Model(&model.AuInterviewUser{}).
		Where("created_at BETWEEN ? AND ?", global.GVA_CONFIG.Interview.OpenTime, global.GVA_CONFIG.Interview.EndTime).Order("created_at DESC")
	var count int64
	if isId {
		//global.GVA_DB.Where("au_interview_department_id = ? AND status = ?", department.ID,status).Preload("AuInterviewTimeTable").Limit(limit).Offset((page - 1) * limit).Find(&users).Count(&count)
		db.Preload("AuInterviewTimeTable").Where("au_interview_department_id = ? AND status = ? AND student_id like CONCAT('%',?,'%')", department.ID, status, query).Limit(limit).Offset((page - 1) * limit).
			Find(&users).Offset(-1).Limit(-1).Count(&count)
	} else {
		//global.GVA_DB.Where("au_interview_department_id = ? AND status = ?", department.ID,status).Preload("AuInterviewTimeTable").Limit(limit).Offset((page - 1) * limit).Find(&users).Count(&count)
		db.Preload("AuInterviewTimeTable").Where("au_interview_department_id = ? AND status = ? AND name like CONCAT('%',?,'%')", department.ID, status, query).Limit(limit).Offset((page - 1) * limit).
			Find(&users).Offset(-1).Limit(-1).Count(&count)
	}

	return users, count
}
func FindDepartmentAllUsers(name string, status string) ([]model.AuInterviewUser, int64) {
	department := FindDepartment(name)
	var users []model.AuInterviewUser
	var count int64
	db := global.GVA_DB.Model(&model.AuInterviewUser{}).
		Where("created_at BETWEEN ? AND ?", global.GVA_CONFIG.Interview.OpenTime, global.GVA_CONFIG.Interview.EndTime)

	db.Where("au_interview_department_id = ? AND status = ?", department.ID, status).Order("created_at DESC").Preload("AuInterviewTimeTable").Find(&users).Count(&count)
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

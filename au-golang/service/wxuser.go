package service

import (
	"au-go/global"
	"au-go/model"
)

func FindWxUsersByStudentId(studentId []string) (wxUsers []model.Wxuser) {
	global.GVA_DB.Where("student_id IN ?", studentId).Find(&wxUsers)
	return wxUsers
}
func FindWxUserByStudentId(studentId string) (wxUsers model.Wxuser) {
	global.GVA_DB.Where("student_id = ?", studentId).Find(&wxUsers)
	return wxUsers
}

func UpdateWxUserFirstAss(studentId string, assId int) {
	global.GVA_DB.Model(model.Wxuser{}).Where("student_id = ? ", studentId).Updates(map[string]interface{}{"ass_entity_first": assId})
}
func UpdateWxUserFirstSecond(studentId string, assId int) {
	global.GVA_DB.Model(model.Wxuser{}).Where("student_id = ? ", studentId).Updates(map[string]interface{}{"ass_entity_second": assId})
}

func FindWxUser() (wxUsers []model.Wxuser) {
	global.GVA_DB.Find(&wxUsers)
	return wxUsers
}

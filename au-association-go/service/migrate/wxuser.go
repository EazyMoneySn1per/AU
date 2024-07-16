package service

import (
	"au-golang/global"
	"au-golang/model/WxBean"
)

func FindWxUserByStudentId(studentId string) (wxUsers WxBean.WxUser) {
	global.GVA_DB.Where("student_id = ?", studentId).Find(&wxUsers)
	return wxUsers
}

func UpdateWxUserFirstAss(studentId string, assId int) {
	global.GVA_DB.Model(WxBean.WxUser{}).Where("student_id = ? ", studentId).Updates(map[string]interface{}{"ass_entity_first": assId})
}
func UpdateWxUserFirstSecond(studentId string, assId int) {
	global.GVA_DB.Model(WxBean.WxUser{}).Where("student_id = ? ", studentId).Updates(map[string]interface{}{"ass_entity_second": assId})
}

func FindWxUser() (wxUsers []WxBean.WxUser) {
	global.GVA_DB.Find(&wxUsers)
	return wxUsers
}

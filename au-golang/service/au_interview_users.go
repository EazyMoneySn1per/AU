package service

import (
	"au-go/global"
	"au-go/model"
)

//func Test() () {
//	timeTable := model.AuInterviewTimeTable{
//		GVA_MODEL:         global.GVA_MODEL{},
//		AuInterviewUserID: 0,
//		UUID:              uuid.UUID{},
//		Month:             "9",
//		Date:              "10",
//		Hour:              "14",
//		Minute:            "10",
//		Location:          "502",
//	}
//	user := model.AuInterviewUser{
//		GVA_MODEL:               global.GVA_MODEL{},
//		UUID:                    uuid.UUID{},
//		Name:                    "费一凡2",
//		StudentId:               "201904020217",
//		PhoneNum:                "17835340505",
//		WxNum:                   "fran6k",
//		Description:             "我是fyf",
//		Sex:                     "男",
//		AuInterviewDepartmentID: 0,
//		AuInterviewTimeTable:    timeTable,
//	}
//	//global.GVA_DB.Create(&user)
//	var depart model.AuInterviewDepartment
//	global.GVA_DB.Where("id = ?",1).Find(&depart)
//	fmt.Println(depart)
//	users := append(depart.AuInterviewUsers,user)a
//	depart.AuInterviewUsers = users
//	global.GVA_DB.Save(depart)
//}
func AddUser(u model.AuInterviewUser) (err error, userInter model.AuInterviewUser) {
	err = global.GVA_DB.Create(&u).Error
	return err, u
}

func FindStudentInterviewLists(studentId string) (err error, user []model.AuInterviewUser) {
	global.GVA_DB.Preload("AuInterviewTimeTable").Where("student_id = ?", studentId).Find(&user)
	return err, user
}

func FindStudentInterviewAndDepartment(studentId string, departmentId uint) (err error, user model.AuInterviewUser) {
	global.GVA_DB.Preload("AuInterviewTimeTable").Where("student_id = ? AND au_interview_department_id = ?", studentId, departmentId).Find(&user)
	return err, user
}

func FindUserByUUID(uuid string) (err error, user model.AuInterviewUser) {
	global.GVA_DB.Preload("AuInterviewTimeTable").Where("uuid = ?", uuid).Find(&user)
	return err, user
}
func SetInterviewUserStatusInBatch(uuids []string, value string) (err error) {
	err = global.GVA_DB.Model(model.AuInterviewUser{}).Where("uuid IN ? ", uuids).Updates(map[string]interface{}{"status": value}).Error
	err = global.GVA_DB.Model(model.AuInterviewTimeTable{}).Where("uuid IN ? ", uuids).Updates(map[string]interface{}{"month": -1, "date": -1, "hour": -1, "minute": -1, "location": "-1"}).Error
	return err
}
func SetInterviewUserTime(uuid string, month interface{}, date interface{}, hour interface{}, minute interface{}, location string) (err error) {
	err = global.GVA_DB.Model(model.AuInterviewTimeTable{}).
		Where("uuid = ?", uuid).
		Updates(map[string]interface{}{"month": month, "date": date, "hour": hour, "minute": minute, "location": location}).Error
	return err
}

package system

import (
	"au-golang/global"
	"au-golang/model/WxBean"
	"gorm.io/gorm"
)

type InterViewUserService struct{}

func (i *InterViewUserService) AddUser(u WxBean.InterViewUser) (err error) {
	err = global.GVA_DB.Create(&u).Error
	return
}

func (interViewUserService *InterViewUserService) FindByStudentIdAndInterViewAss(studentId string, assId int) (exist bool) {
	var interviewUser []WxBean.InterViewUser
	db := global.GVA_DB.Model(WxBean.InterViewUser{})
	db.Where("student_id = ?", studentId).First(&interviewUser)
	for _, v := range interviewUser {
		if v.SubmitAssId == assId && v.InterViewStatus <= 3 {
			exist = true
			return
		}
	}
	exist = false
	return
}

// DeleteBatchByAssId 批量删除
func (interViewUserService *InterViewUserService) DeleteBatchByAssId(assId int) {
	db := global.GVA_DB.Model(WxBean.InterViewUser{})
	_ = db.Where("submit_ass_id = ?", assId).Delete(&WxBean.InterViewUser{}).Error
}

func (interViewUserService *InterViewUserService) Save(interViewUser WxBean.InterViewUser) {
	db := global.GVA_DB.Model(WxBean.InterViewUser{}).Debug()
	db.Where("id = ?", interViewUser.Id).Save(&interViewUser)
}

func (interViewUserService *InterViewUserService) Create(interViewUser *WxBean.InterViewUser) (err error) {
	db := global.GVA_DB
	var ass WxBean.InterViewAss
	err = db.Model(&WxBean.InterViewAss{}).Where("ass_id", interViewUser.SubmitAssId).Find(&ass).Error
	interViewUser.InterViewAss = ass
	err = db.Model(&WxBean.InterViewUser{}).Create(interViewUser).Error
	return
}

func (interViewUserService *InterViewUserService) FindById(id string) (WxBean.InterViewUser, error) {
	var interviewUser WxBean.InterViewUser
	db := global.GVA_DB.Model(WxBean.InterViewUser{})
	err := db.Preload("InterViewAss").Where("id = ?", id).First(&interviewUser).Error
	return interviewUser, err
}

func (interViewUserService *InterViewUserService) FindAlByStudentId(studentId string) []WxBean.InterViewUser {
	var interviewUsers []WxBean.InterViewUser
	db := global.GVA_DB.Model(WxBean.InterViewUser{})
	db.Preload("InterViewAss").Where("student_id = ?", studentId).Find(&interviewUsers)
	return interviewUsers
}

func (i *InterViewUserService) FindAlByInterviewAssCount(studentId string, assId int) int64 {
	var interviewUsers []WxBean.InterViewUser
	db := global.GVA_DB.Model(WxBean.InterViewUser{})
	var count1 int64
	var count2 int64
	var count3 int64
	db.Where("student_id", studentId).Where("inter_view_status", 1).Find(&interviewUsers).Count(&count1)
	db.Where("student_id", studentId).Where("inter_view_status", 2).Find(&interviewUsers).Count(&count2)
	db.Where("student_id", studentId).Where("inter_view_status", 3).Find(&interviewUsers).Count(&count3)
	return count1 + count2 + count3
}

func (interViewUserService *InterViewUserService) FindStudentInterviewLists(studentId string) (err error, user []WxBean.InterViewUser) {
	global.GVA_DB.Preload("AuInterviewTimeTable").Where("student_id = ?", studentId).Find(&user)
	return err, user
}

func (i *InterViewUserService) UpdateInfo(interviewUser WxBean.InterViewUser) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		//踩坑：Updates只会更新非零值，Save会更新所有字段
		if err := tx.Model(&WxBean.InterViewUser{}).Where("id=?", interviewUser.Id).Save(&interviewUser).Error; err != nil {
			return err
		}
		return nil
	})
	return

}

func (i *InterViewUserService) GetUserInfo(user WxBean.InterViewUser) (u WxBean.InterViewUser, err error) {
	err = global.GVA_DB.Model(&WxBean.InterViewUser{}).Preload("InterViewAss").Where("student_id=?", user.StudentId).Find(&u).Error
	return
}

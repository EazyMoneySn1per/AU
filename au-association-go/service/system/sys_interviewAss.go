package system

import (
	"au-golang/global"
	"au-golang/model/WxBean"
	"au-golang/model/system"

	"gorm.io/gorm"
)

type InterViewAssService struct{}

func (i *InterViewAssService) FindByAssname(name string) (system.Ass, error) {
	var Ass system.Ass
	err := global.GVA_DB.Model(&system.Ass{}).Where("assname", name).Find(&Ass).Error
	return Ass, err
}

func (interViewAssService *InterViewAssService) FindByAssId(assId int) WxBean.InterViewAss {
	var interviewAss WxBean.InterViewAss
	db := global.GVA_DB.Model(WxBean.InterViewAss{})
	db.Where("ass_id= ?", assId).First(&interviewAss)
	return interviewAss
}

func (interViewAssService *InterViewAssService) Save(ass *WxBean.InterViewAss) {
	db := global.GVA_DB.Model(WxBean.InterViewAss{})
	db.Where("ass_id = ?", ass.AssId).Save(ass)
}

func (i *InterViewAssService) GetInterViewUsers(assid, status, page, limit int) (interviewUsers []WxBean.InterViewUser, totalInterviewUsers, total int64, err error) {

	result := global.GVA_DB.Model(&WxBean.InterViewUser{}).
		Where("create_time BETWEEN ? AND ?", global.GVA_CONFIG.Interview.OpenTime, global.GVA_CONFIG.Interview.EndTime)

	result.Where("submit_ass_id", assid).Count(&totalInterviewUsers)
	result.Where("inter_view_status", status).Count(&total)

	if page == 0 {
		page = 1
	}

	offset := (page - 1) * limit
	err = result.Offset(offset).Limit(limit).Find(&interviewUsers).Order("create_time DESC").Error
	return
}

func (i *InterViewAssService) UpdateInfo(interviewAss WxBean.InterViewAss) {
	global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&WxBean.InterViewAss{}).Where("ass_id=?", interviewAss.AssId).Updates(
			WxBean.InterViewAss{ConfirmJoinMessage: interviewAss.ConfirmJoinMessage, ShowMessage: interviewAss.ShowMessage,
				PresidentWechat: interviewAss.PresidentWechat, PresidentName: interviewAss.PresidentName}).Error; err != nil {
			return err
		}
		return nil
	})
}
func (interViewAssService *InterViewAssService) Create(ass *WxBean.InterViewAss) {
	db := global.GVA_DB.Model(WxBean.InterViewAss{})
	db.Create(ass)
}

func (interViewAssService *InterViewAssService) Delete(assId int) {
	var interviewAss WxBean.InterViewAss
	db := global.GVA_DB.Model(WxBean.InterViewAss{})
	db.Where("ass_id = ?", assId).Delete(&interviewAss)
}

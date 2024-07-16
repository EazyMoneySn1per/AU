package system

import (
	"au-golang/global"
	"au-golang/model/WxBean"
)

type SysInterviewService struct{}

func (interviewService *SysInterviewService) FindByAssId(assId int) (WxBean.InterViewAss, error) {
	var interviewAss WxBean.InterViewAss
	db := global.GVA_DB.Model(&WxBean.InterViewAss{})
	err := db.Where("assid = ?", assId).First(&interviewAss).Error
	return interviewAss, err
}

package system

import (
	"au-golang/global"
	"au-golang/model/system"
)

type SysnthesizeSubmitService struct{}

func (SysnthesizeSubmitService *SysnthesizeSubmitService) Save(data *system.SynthesizeSubmit) {
	db := global.GVA_DB.Model(&system.SynthesizeSubmit{})
	db.Save(data)
}

func (SysnthesizeSubmitService *SysnthesizeSubmitService) FindByAssEntityOrderByCreateTimeDesc(assId string) []system.SynthesizeSubmit {
	var synthesizeSubmit []system.SynthesizeSubmit
	db := global.GVA_DB.Model(&system.SynthesizeSubmit{})
	db.Preload("Ass").Where("ass_entity = ?", assId).Order("create_time desc").Find(&synthesizeSubmit)
	return synthesizeSubmit
}

func (SysnthesizeSubmitService *SysnthesizeSubmitService) Create(data *system.SynthesizeSubmit) {
	db := global.GVA_DB.Model(&system.SynthesizeSubmit{})
	db.Create(data)
}

package system

import (
	"au-golang/global"
	"au-golang/model/common/request"
	"au-golang/model/system"
	"math/rand"
	"time"
)

type OutlayService struct{}

func (outlayService *OutlayService) FindByAssidAndStepIsNot(assid, finalstep int) (system.Outlay, error) {
	outlay := system.Outlay{}
	db := global.GVA_DB.Model(system.Outlay{})
	err := db.Where("assid = ?", assid).Not("step = ?", finalstep).First(&outlay).Error
	return outlay, err
}

func (outlayService *OutlayService) Create(outlay *system.Outlay) {
	rand.Seed(time.Now().Unix())
	Id := rand.Int31()
	outlay.Id = int(Id)
	db := global.GVA_DB.Model(system.Outlay{})
	db.Create(outlay)
}

func (outlayService *OutlayService) Save(outlay *system.Outlay) {
	db := global.GVA_DB.Model(system.Outlay{})
	db.Where("id = ?", outlay.Id).Save(outlay)
}

func (outlayService *OutlayService) FindAllByStepNot(finalstep int) []system.Outlay {
	var outlay []system.Outlay
	db := global.GVA_DB.Model(system.Outlay{})
	db.Not("step = ?", finalstep).Find(&outlay)
	return outlay
}

func (outlayService *OutlayService) FindAllByStep(finalstep int) []system.Outlay {
	var outlays []system.Outlay
	db := global.GVA_DB.Model(system.Outlay{})
	db.Debug().Where("step = ?", finalstep).Find(&outlays)
	return outlays
}

func (outlayService *OutlayService) FindTodayAll() []system.Outlay {
	var outlays []system.Outlay
	db := global.GVA_DB.Model(system.Outlay{})
	sql := "SELECT * FROM `outlay` WHERE DATE(create_time) = CURDATE()"
	db.Raw(sql).Find(&outlays)
	return outlays
}

func (o *OutlayService) FindByAssidAndStep(step, assid int, pageInfo request.PageInfo) (list []system.Outlay, total int64) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)

	result := global.GVA_DB.Model(&system.Outlay{}).Where("assid =? and step =?", assid, step).Count(&total)
	result.Limit(limit).Offset(offset).Find(&list)
	return
}

func (o *OutlayService) FindAllByStepLimit(step int, pageInfo request.PageInfo) (list []system.Outlay, total int64) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.GVA_DB.Model(&system.Outlay{}).Where("step=?", step).Count(&total)
	db.Limit(limit).Offset(offset).Scan(&list)
	return
}

func (o *OutlayService) GetRecord(assId, page, limit int) (outlay []system.Outlay, count int64) {
	offset := (page - 1) * limit
	global.GVA_DB.Model(&system.Outlay{}).Offset(offset).Limit(limit).Where("assid", assId).Find(&outlay).Count(&count)
	return
}

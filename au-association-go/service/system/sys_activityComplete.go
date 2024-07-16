package system

import (
	"au-golang/global"
	"au-golang/model/common/request"
	"au-golang/model/system"
)

type ActivityCompleteService struct{}

func (activityCompleteService *ActivityCompleteService) FindById(id int) system.ActivityComplete {
	var activityComplete system.ActivityComplete
	db := global.GVA_DB.Model(&system.ActivityComplete{})
	// 联表查询
	db.Preload("AssStruct").Where("id = ?", id).First(&activityComplete)
	return activityComplete
}

func (activityCompleteService *ActivityCompleteService) Save(event *system.ActivityComplete) {
	db := global.GVA_DB.Model(&system.ActivityComplete{})
	db.Save(event)
}

func (activityCompleteService *ActivityCompleteService) FindAll() []system.ActivityComplete {
	var activityCompletes []system.ActivityComplete
	db := global.GVA_DB.Model(&system.ActivityComplete{})
	// 联表查询
	db.Preload("AssStruct").Find(&activityCompletes)
	return activityCompletes
}

func (activityCompleteService *ActivityCompleteService) FindAllLimit(info request.PageInfo) (list []system.ActivityComplete, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&system.ActivityComplete{}).Count(&total)
	// 联表查询
	db.Limit(limit).Offset(offset).Preload("AssStruct").Scan(&list)
	return
}

func (activityCompleteService *ActivityCompleteService) FindallbyassAssnameorassAssidorderbydate(assName string, assId int, info request.PageInfo) (list []system.ActivityComplete, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&system.ActivityComplete{}).Where("ass = ?", assId).Count(&total)
	db.Limit(limit).Offset(offset).Preload("AssStruct").Order("date desc").Scan(&list)
	return
}

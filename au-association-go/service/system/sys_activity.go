package system

import (
	"au-golang/global"
	"au-golang/model/system"
)

type ActivityService struct{}

func (activityService *ActivityService) FindByAssidAndStepIsNot(assid string, step int) (system.Activity, error) {
	activity := system.Activity{}
	db := global.GVA_DB.Model(&system.Activity{})
	err := db.Where("assid = ?", assid).Not("step = ?", step).First(&activity).Error
	return activity, err
}

func (activityService *ActivityService) Save(activity *system.Activity) {
	db := global.GVA_DB.Model(&system.Activity{})
	db.Where("id = ?", activity.Id).Save(activity)
}

func (activityService *ActivityService) Create(activity *system.Activity) {
	db := global.GVA_DB.Model(&system.Activity{})
	db.Create(activity)
}

func (activityService *ActivityService) FindAllByStepNot(n int) []system.Activity {
	var activities []system.Activity
	db := global.GVA_DB.Model(system.Activity{})
	db.Not("step = ?", n).Find(&activities)
	return activities
}

func (activityService *ActivityService) FindTodayAll() []system.Activity {
	var activities []system.Activity
	db := global.GVA_DB.Model(system.Activity{})
	sql := "SELECT * FROM `activity` WHERE DATE(create_time) = CURDATE()"
	db.Debug().Raw(sql).Find(&activities)
	return activities
}

func (a *ActivityService) GetRecord(assId, page, limit int) (activity []system.Activity, count int64) {
	offset := (page - 1) * limit
	global.GVA_DB.Model(&system.Activity{}).Offset(offset).Limit(limit).Where("assid", assId).Find(&activity).Count(&count)
	return
}

package system

import (
	"au-golang/global"
	"au-golang/model/common/request"
	"au-golang/model/system"
)

type TwitterService struct{}

func (twitterService *TwitterService) FindByAssidAndStepIsNot(assid, finalstep int) (system.Twitter, error) {
	twitter := system.Twitter{}
	db := global.GVA_DB.Model(system.Twitter{})
	err := db.Where("assid = ?", assid).Not("step = ?", finalstep).First(&twitter).Error
	return twitter, err
}

func (twitterService *TwitterService) Save(twitter *system.Twitter) {
	db := global.GVA_DB.Model(system.Twitter{})
	_ = db.Create(twitter).Error
}

func (twitterService *TwitterService) FindAllByStepNot(finalstep int) []system.Twitter {
	var twitters []system.Twitter
	db := global.GVA_DB.Model(system.Twitter{})
	db.Debug().Not("step = ?", finalstep).Find(&twitters)
	return twitters
}

func (twitterService *TwitterService) FindAllByStep(finalstep int) []system.Twitter {
	var twitters []system.Twitter
	db := global.GVA_DB.Model(system.Twitter{})
	db.Where("step = ?", finalstep).Find(&twitters)
	return twitters
}

func (twitterService *TwitterService) FindAllByStepLimit(step int, info request.PageInfo) (list []system.Twitter, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&system.Twitter{}).Where("step = ?", step).Count(&total)
	db.Limit(limit).Offset(offset).Scan(&list)
	return
}

func (twitterService *TwitterService) FindTodayAll() []system.Twitter {
	var twitters []system.Twitter
	db := global.GVA_DB.Model(system.Twitter{})
	sql := "SELECT * FROM `twitter` WHERE DATE(create_time) = CURDATE()"
	db.Raw(sql).Find(&twitters)
	return twitters
}

func (twitterService *TwitterService) FindAllByAssIdAndStep(step, assId int, info request.PageInfo) (list []system.Twitter, total int64) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&system.Twitter{}).Where("assid = ? and step = ?", assId, step).Count(&total)
	//err = db.Where("assid = ? and step = ?", assId, step).Count(&total).Error
	db.Limit(limit).Offset(offset).Scan(&list)
	return
}

func (t *TwitterService) UpdateInfo(twitter *system.Twitter) (err error) {
	err = global.GVA_DB.Model(system.Twitter{}).Where("id=?", twitter.Id).Save(twitter).Error
	return
}

func (t *TwitterService) GetRecord(assId, page, limit int) (twitter []system.Twitter, count int64) {
	offset := (page - 1) * limit
	global.GVA_DB.Model(&system.Twitter{}).Offset(offset).Limit(limit).Where("assid", assId).Find(&twitter).Count(&count)
	return
}

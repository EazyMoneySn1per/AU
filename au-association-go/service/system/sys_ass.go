package system

import (
	"au-golang/global"
	"au-golang/model/system"
	"fmt"
)

type AssService struct{}

func (AssService *AssService) FindByAssname(assName string) (system.Ass, error) {
	ass := system.Ass{}
	db := global.GVA_DB.Model(&system.Ass{})
	err := db.Where("assname = ? ", assName).First(&ass).Error
	return ass, err
}

// 返回可以创建社团的最大id
func (AssService *AssService) FindAssIdMax() int {
	ass := system.Ass{}
	db := global.GVA_DB.Model(&system.Ass{})
	db.Order("assid desc").Limit(1).Find(&ass)
	return ass.Assid + 1
}

func (AssService *AssService) FindByAssid(n int) system.Ass {
	ass := system.Ass{}
	db := global.GVA_DB.Model(&system.Ass{})
	db.Preload("AssUsers").Where("assid = ?", n).First(&ass)
	return ass
}

func (AssService *AssService) Save(ass *system.Ass) {
	db := global.GVA_DB.Model(&system.Ass{})
	err := db.Debug().Where("assid = ?", ass.Assid).Save(ass)
	fmt.Println(err)
}

func (AssService *AssService) FindAll() []system.Ass {
	var ass []system.Ass
	db := global.GVA_DB.Model(&system.Ass{})
	db.Find(&ass)
	return ass
}

func (AssService *AssService) UpdateByAssId(data *system.Ass) {
	db := global.GVA_DB.Model(&system.Ass{})
	err := db.Debug().Where("assid = ?", data.Assid).Updates(data)
	fmt.Println(err)
}

func (AssService *AssService) Delete(data int) {
	db := global.GVA_DB.Model(&system.Ass{})
	err := db.Debug().Where("assid = ?", data).Delete(&system.Ass{})
	fmt.Println(err)
}

func (a *AssService) Create(ass system.Ass) (err error) {
	err = global.GVA_DB.Model(&system.Ass{}).Create(&ass).Error
	return
}

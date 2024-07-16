package system

import (
	"au-golang/global"
	"au-golang/model/WxBean"
	"au-golang/model/system"
)

type UserService struct{}

// 用于FindAllByRole()函数的结构体
type ResultUser struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Assid    int    `json:"assId"`
	Assname  string `json:"assName"`
}

func (userService *UserService) FindByName(name string) (error, system.User) {
	var userr system.User
	db := global.GVA_DB.Model(&system.User{})
	err := db.Where("name = ?", name).First(&userr).Error
	return err, userr
}

func (userService *UserService) FindByNameAndPassword(name, password string) (system.User, error) {
	var user system.User
	db := global.GVA_DB.Model(&system.User{})
	err := db.Where("name = ? and password = ?", name, password).First(&user).Error
	return user, err
}

func (userService *UserService) FindAllByRole(role string, page, limit int) (err error, list []system.User, total int64) {

	offset := limit * (page - 1)

	db := global.GVA_DB.Preload("Ass").Model(&system.User{}).Where("role = ?", role).Count(&total)
	err = db.Offset(offset).Limit(limit).Find(&list).Error

	return err, list, total
}

func (userService *UserService) DeleteUser(user system.User) {
	db := global.GVA_DB.Model(&system.User{})
	db.Delete(&user)
}

func (userService *UserService) DeleteBatchUser(assId int) {
	db := global.GVA_DB.Model(&system.User{})
	_ = db.Where("assid = ?", assId).Delete(&system.User{}).Error
}

func (userService *UserService) Save(user system.User) (err error) {
	db := global.GVA_DB.Model(&system.User{})
	err = db.Where("id=?", user.Id).Save(&user).Error
	return
}
func (u *UserService) FindAll(page, limit int) (user []system.User, total int64, err error) {

	offset := limit * (page - 1)
	result := global.GVA_DB.Preload("Ass").Model(&system.User{}).Count(&total)
	err = result.Offset(offset).Limit(limit).Find(&user).Error
	return
}

func (u *UserService) FindAssUsersByAssId(assId string) (users []WxBean.WxUser) {

	global.GVA_DB.Model(&WxBean.WxUser{}).Where("ass_entity_first = ? or ass_entity_second = ?", assId, assId).Find(&users)
	return
}

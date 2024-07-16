package request

import "mime/multipart"

type AddUser struct {
	Name        string `json:"name" binding:"required"`
	StudentId   string `json:"studentId" binding:"required"`
	PhoneNum    string `json:"phoneNum" binding:"required"`
	WxNum       string `json:"wxNum" binding:"required"`
	Description string `json:"description" binding:"required"`
	Sex         string `json:"sex" binding:"required"`
	Department  string `json:"department" binding:"required"`
	//PhoneNum    string  `json:"phoneNum" gorm:"default:'QMPlusUser'"`
	//HeaderImg   string `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	//AuthorityId string `json:"authorityId" gorm:"default:888"`
}

type DepartmentAndStatus struct {
	PageInfo   PageInfo
	Department string `json:"department" binding:"required" form:"department"`
	Status     string `json:"status" binding:"required" form:"status"`
}
type DepartmentAndStatusNoPage struct {
	Department string `json:"department" binding:"required" form:"department"`
	Status     string `json:"status" binding:"required" form:"status"`
}

type StatusAndUuids struct {
	Status string   `json:"status"`
	Uuids  []string `json:"uuids"`
}

type FileAndDepartment struct {
	File       multipart.File `json:"file" form:"file"`
	Department string         `json:"department" form:"department"`
}

type InterviewTime struct {
	Uuid     string `json:"uuid" binding:"required"`
	Month    int    `json:"month" binding:"required"`
	Date     int    `json:"date" binding:"required"`
	Hour     int    `json:"hour" binding:"required"`
	Minute   int    `json:"minute" binding:"required"`
	Location string `json:"location" binding:"required"`
}

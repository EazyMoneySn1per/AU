package WxBean

import (
	"au-golang/global"
	"time"
)

type InterViewUser struct {
	Id            string `json:"id" gorm:"primary_key"`
	Name          string `json:"name"`
	StudentId     string `json:"studentId"`
	Description   string `json:"description"`
	PhoneNum      string `json:"phoneNum"`
	WxNum         string `json:"wxNum"`
	BackMessage   string `json:"backMessage"`
	Sex           string `json:"sex"`
	ButtonControl int    `json:"buttonControl"`
	SubmitAssId   int    `json:"submitAssId"`

	InterViewStatus int

	InterViewAss InterViewAss `gorm:"foreignKey:SubmitAssId;references:AssId"`

	CreatedAt time.Time `gorm:"column:create_time" json:"createTime"`

	UpdatedAt time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (InterViewUser) TableName() string {
	return "inter_view_user"
}

func (InterViewUser) GetInfo(info InterViewUser) (result map[string]interface{}) {
	result = make(map[string]interface{})
	result["id"] = info.Id
	result["studentId"] = info.StudentId
	result["description"] = info.Description
	result["sex"] = info.Sex
	result["studentName"] = info.Name
	result["wxNum"] = info.WxNum
	result["phoneNum"] = info.PhoneNum
	result["assName"] = info.InterViewAss.AssName
	result["assLogo"] = info.InterViewAss.Logo
	result["presidentName"] = info.InterViewAss.PresidentName
	result["presidentWechat"] = info.InterViewAss.PresidentWechat
	result["interViewStatus"] = info.InterViewStatus
	result["interViewStatusMessage"] = global.InterViewEnum{}.GetMessageByStage(info.InterViewStatus)
	result["buttonControl"] = info.ButtonControl
	return
}

package model

// wxuser学生表
type Wxuser struct {
	Id string `json:"id"`
	Avatar string `json:"avatar"`
	Nickname string `json:"nickname"`
	OpenId string `json:"openIid"`
	RealName string `json:"realName"`
	AssEntityFirst int `json:"assEntityFirst"`
	AssEntitySecond int `json:"assEntitySecond"`
	StudentId string `json:"studentId"`
	PhoneNum string `json:"phoneNum"`
	WeChatId string `json:"weChatId"`
}
func (Wxuser) TableName() string {
	return "wxuser"
}
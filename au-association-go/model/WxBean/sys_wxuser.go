package WxBean

type WxUser struct {
	Id              string `json:"id" gorm:"primary_key"`
	Nickname        string `json:"nickName"`
	OpenId          string `json:"openId"`
	MpOpenId        string `json:"MpOpenId"` //新增字段，小程序用户openid
	RealName        string `json:"realName"`
	Avatar          string `json:"avatar"`
	StudentId       string `json:"studentId"`
	WeChatId        string `json:"weChatId"`
	PhoneNum        string `json:"phoneNum"`
	AssEntityFirst  int    `json:"ass_entity_first"`
	AssEntitySecond int    `json:"ass_entity_second"`
}

func (WxUser) TableName() string {
	return "wxuser"
}

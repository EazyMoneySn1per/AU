package WxBean

type InterViewAss struct {
	Id                 string `json:"id"`
	AssId              int    `gorm:"index"`
	AssName            string
	Logo               string
	PresidentName      string
	PresidentWechat    string
	ShowMessage        string
	ConfirmJoinMessage string
	CodeUrl            string
}

func (InterViewAss) TableName() string {
	return "inter_view_ass"
}

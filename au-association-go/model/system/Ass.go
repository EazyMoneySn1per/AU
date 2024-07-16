package system

type Ass struct {
	// 社团id
	Assid int    `json:"assId" form:"assId" gorm:"primary_key"`
	Logo  string `json:"logo" form:"-"`
	// 社团名称
	Assname string `json:"assName" form:"assName"`
	// 社团简介
	AssDescription string `json:"assDescription" form:"assDescription"`
	// 社团类型
	Asstype string `json:"assType" form:"assType"`
	// 社团经费
	Money float64 `json:"money" form:"money"`
	// 优秀社团
	IsExecllent int `json:"isExecllent" form:"isExecllent"`
	//社长名称
	Presidentname string `json:"presidentName" form:"presidentName"`
	// 社长学号
	PresidentId string `json:"presidentId" form:"presidentId"`
	// 指导老师姓名
	Teachername string `json:"teacherName" form:"teacherName"`
	// 知道老师联系方式
	Teacherphone string `json:"teacherPhone" form:"teacherPhone"`
	// 指导老师所在单位
	Teacherpost string `json:"teacherPost" form:"teacherPost"`

	//AssUsers []User `gorm:"foreignKey:AssEntity;references:Assid"` //后台系统用户

	//AssWxUsers []WxBean.WxUser `gorm:"foreignKey:AssEntityFirst;foreignKey:AssEntitySecond;AssociactionPrimariKey:Assid"`
}

func (Ass) TableName() string {
	return "ass"
}

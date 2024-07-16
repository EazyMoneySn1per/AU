package model

type AuInterviewDepartment struct {
	ID               uint   `json:"id" gorm:"primarykey"`
	DepartmentName   string `json:"departmentName" gorm:"comment:部门名称"`
	AuInterviewUsers []AuInterviewUser
	InterviewMsg
}

type InterviewMsg struct {
	StageOneMessage        string `json:"StageOneMessage" gorm:"comment:一面的消息"`
	StageTwoMessage        string `json:"StageTwoMessage" gorm:"comment:二面的消息"`
	StageTwoSuccessMessage string `json:"StageFinallyMessage" gorm:"comment:加入成功的消息"`
	StageOneFailedMessage  string `json:"StageOneFailedMessage" gorm:"comment:一面失败的消息"`
	StageTwoFailedMessage  string `json:"StageTwoFailedMessage" gorm:"comment:二面失败的消息"`
}

// GetInterviewMsg 获取部门面试消息
func (p *AuInterviewDepartment) GetInterviewMsg(status string) string {
	switch status {
	case StageOne:
		return p.StageOneMessage
	case StageTwo:
		return p.StageTwoMessage
	case StageTwoFailed:
		return p.StageTwoFailedMessage
	case StageOneFailed:
		return p.StageOneFailedMessage
	case StageTwoSuccess:
		return p.StageTwoSuccessMessage
	}
	return "未知错误"
}

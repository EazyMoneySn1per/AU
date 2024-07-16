package model

import (
	"au-go/global"
	"github.com/satori/go.uuid"
	"sort"
)

type AuInterviewUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Name        string    `json:"name" gorm:"comment:名字"`
	StudentId   string    `json:"studentId" gorm:"comment:学号"`
	PhoneNum    string    `json:"phoneNum" gorm:"comment:手机号"`
	WxNum       string    `json:"wxNum" gorm:"comment:微信号"`
	Description string    `json:"description" gorm:"comment:个人简介"`
	Sex         string    `json:"sex" gorm:"comment:性别"`
	/**
	1:一面
	2:二面
	3:面试成功
	9:一面失败
	8:二面失败
	*/
	Status                  string `json:"status" gorm:"comment:面试状态"`
	AuInterviewDepartmentID uint
	AuInterviewTimeTable    AuInterviewTimeTable
}

type SortUserList []AuInterviewUser

func (p SortUserList) Len() int {
	return len(p)
}
func (p SortUserList) Less(i, j int) bool {
	if p[i].AuInterviewTimeTable.Month > p[j].AuInterviewTimeTable.Month {
		return true
	}
	if p[i].AuInterviewTimeTable.Month < p[j].AuInterviewTimeTable.Month {
		return false
	}
	if p[i].AuInterviewTimeTable.Date > p[j].AuInterviewTimeTable.Date {
		return true
	}
	if p[i].AuInterviewTimeTable.Date < p[j].AuInterviewTimeTable.Date {
		return false
	}
	if p[i].AuInterviewTimeTable.Hour > p[j].AuInterviewTimeTable.Hour {
		return true
	}
	if p[i].AuInterviewTimeTable.Hour < p[j].AuInterviewTimeTable.Hour {
		return false
	}
	if p[i].AuInterviewTimeTable.Minute > p[j].AuInterviewTimeTable.Minute {
		return true
	} else {
		return false
	}
}
func (p SortUserList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p *SortUserList) Sort() {
	sort.Sort(p)
}

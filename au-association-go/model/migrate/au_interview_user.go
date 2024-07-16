package model

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"sort"
	"strconv"
)

type AuInterviewUser struct {
	gorm.Model
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

func (p *AuInterviewUser) GetStatus() string {
	switch p.Status {
	case "1":
		return StageOne
	case "2":
		return StageTwo
	case "3":
		return StageTwoSuccess
	case "8":
		return StageTwoFailed
	case "9":
		return StageOneFailed
	}
	return "未知错误"
}

func (p *AuInterviewUser) GetStatusByNumber(num string) string {
	switch num {
	case "1":
		return StageOne
	case "2":
		return StageTwo
	case "3":
		return StageTwoSuccess
	case "8":
		return StageTwoFailed
	case "9":
		return StageOneFailed
	}
	return "未知错误"
}

func (p *AuInterviewUser) GetInterviewTime() (timeBack, locationBack string) {

	// 判读有无面试时间和地点
	if p.AuInterviewTimeTable.Month == -1 && p.AuInterviewTimeTable.Date == -1 {
		timeBack = "暂无"
	} else {
		temMinute := strconv.Itoa(p.AuInterviewTimeTable.Minute)
		if len(temMinute) == 1 {
			temMinute = "0" + temMinute
		}
		timeBack = strconv.Itoa(p.AuInterviewTimeTable.Month) + "月" + strconv.Itoa(p.AuInterviewTimeTable.Date) + "日 " + strconv.Itoa(p.AuInterviewTimeTable.Hour) + ":" + temMinute
	}
	if p.AuInterviewTimeTable.Location == "-1" {
		locationBack = "暂无"
	} else {
		locationBack = p.AuInterviewTimeTable.Location
	}
	return
}

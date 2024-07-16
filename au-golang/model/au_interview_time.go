package model

import (
	"au-go/global"
	uuid "github.com/satori/go.uuid"
)

type AuInterviewTimeTable struct {
	global.GVA_MODEL
	AuInterviewUserID uint
	UUID              uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Month             int       `json:"month" gorm:"comment:月"`
	Date              int       `json:"date" gorm:"comment:日期"`
	Hour              int       `json:"hour" gorm:"comment:时"`
	Minute            int       `json:"minute" gorm:"comment:分"`
	Location          string    `json:"location" gorm:"comment:地点"`
}

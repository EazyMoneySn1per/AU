package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type AuInterviewTimeTable struct {
	gorm.Model
	AuInterviewUserID uint
	UUID              uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Month             int       `json:"month" gorm:"comment:月"`
	Date              int       `json:"date" gorm:"comment:日期"`
	Hour              int       `json:"hour" gorm:"comment:时"`
	Minute            int       `json:"minute" gorm:"comment:分"`
	Location          string    `json:"location" gorm:"comment:地点"`
}

package system

import "time"

type Activity struct {
	Id          string    `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Assid       int       `json:"assid"`
	Date        time.Time `json:"date"`
	Step        int       `json:"step"`
	File1       string    `json:"file1"`
	File2       string    `json:"file2"`
	Backmsg     string    `json:"backmsg"`
	Description string    `json:"description"`
	CreatedAt   time.Time `gorm:"column:create_time" json:"createTime"`
	UpdatedAt   time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (Activity) TableName() string {
	return "activity"
}

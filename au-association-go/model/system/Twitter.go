package system

import "time"

type Twitter struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Assid      int    `json:"assid"`
	PictureUrl string `json:"pictureUrl"`
	Step       int    `json:"step"`
	Backmsg    string `json:"backmsg"`

	CreatedAt time.Time `gorm:"column:create_time" json:"createTime"`
	UpdatedAt time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (Twitter) TableName() string {
	return "twitter"
}

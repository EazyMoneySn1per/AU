package system

import "time"

type Outlay struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Money     float64 `json:"money"`
	Assid     int     `json:"assid"`
	OutlayUrl string  `json:"outlayUrl"`
	Step      int     `json:"step"`
	Backmsg   string  `json:"backmsg"`

	CreatedAt time.Time `gorm:"column:create_time" json:"createTime"`

	UpdatedAt time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (Outlay) TableName() string {
	return "outlay"
}

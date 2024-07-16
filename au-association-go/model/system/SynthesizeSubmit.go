package system

import "time"

type SynthesizeSubmit struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Type        int    `json:"type"`
	Name        string `json:"name"`
	FileUrl     string `json:"fileUrl"`
	Backmsg     string `json:"backmsg"`
	IsHandle    int    `json:"isHandle"`
	AssEntity   int    `json:"assEntity"`
	//	Ass         Ass    `gorm:"foreignKey:ass_entity;AssociationPrimaryKey:assid"`

	CreatedAt time.Time `gorm:"column:create_time" json:"createTime"`

	UpdatedAt time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (SynthesizeSubmit) TableName() string {
	return "synthesizesubmit"
}

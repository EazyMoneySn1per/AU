package system

type User struct {
	Id        string `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"column:name"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Assid     int    `json:"assid"`
	AssEntity int    `json:"assEntity"`
	Ass       Ass    `gorm:"foreignKey:ass_entity;AssociationForeignKey:assid"`
}

func (User) TableName() string {
	return "user"
}

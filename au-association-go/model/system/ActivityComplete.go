package system

import "time"

type Model struct {
	ID       uint      `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreateAt time.Time `gorm:"column:create_time"`
	UpdateAt time.Time `gorm:"column:update_time"`
}

type ActivityComplete struct {
	Model
	ActivityName string    `json:"activityName"`
	Twitter      string    `json:"twitter"`
	Date         time.Time `json:"date"`
	File1        string    `json:"file1"`
	File2        string    `json:"file2"`
	Description  string    `json:"description"`
	Ass          int       `json:"ass"`
	AssStruct    Ass       `gorm:"foreignKey:ass;references:assid"`
}

func (ActivityComplete) TableName() string {
	return "activity_complete"
}

func (ActivityComplete) GetInfo(info ActivityComplete) (result map[string]interface{}) {
	result = make(map[string]interface{})
	result["id"] = info.ID
	result["activityName"] = info.ActivityName
	result["date"] = info.Date
	result["twitter"] = info.Twitter
	result["file1"] = info.File1
	result["file2"] = info.File2
	result["description"] = info.Description
	//result["assName"] = info.AssStruct.Assname
	//result["assId"] = info.AssStruct.Assid
	result["createTime"] = info.CreateAt
	result["updateTime"] = info.UpdateAt
	return
}

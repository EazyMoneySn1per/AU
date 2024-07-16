package response

import uuid "github.com/satori/go.uuid"

type WxUserReturn struct {
	UUID           uuid.UUID `json:"uuid"`
	Name           string    `json:"name"`
	BackMessage    string    `json:"backMessage"`
	Time           string    `json:"time"`
	Location       string    `json:"location"`
	Status         string    `json:"status"`
	StatusMsg      string    `json:"statusMsg"`
	DepartmentName string    `json:"departmentName"`
}

type StatusReturn struct {
	Status int `json:"status"`
}
type ErrImport struct {
	Name string `json:"name"`
	StudentId string `json:"studentId"`
}
type StudentWithTime struct {
	StudentId string `json:"studentId"`
	Name string `json:"name"`
	Month string  `json:"month"`
	Date string `json:"date"`
	Hour string `json:"hour"`
	Minute string `json:"minute"`
	Location string `json:"location"`
}
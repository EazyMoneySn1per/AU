package request

import "au-go/model"

type DepartmentMessage struct {
	DepartmentName string `json:"departmentName"`
	model.InterviewMsg
}

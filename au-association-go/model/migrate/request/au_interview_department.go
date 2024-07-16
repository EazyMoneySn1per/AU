package request

import model "au-golang/model/migrate"

type DepartmentMessage struct {
	DepartmentName string `json:"departmentName"`
	model.InterviewMsg
}

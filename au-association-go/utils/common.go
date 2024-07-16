package utils

import (
	model "au-golang/model/migrate"
)

func GetInterviewStatusMsg(status string) (statusMsg string) {
	switch status {
	case "1":
		statusMsg = model.StageOne
	case "2":
		statusMsg = model.StageTwo
	case "3":
		statusMsg = model.StageTwoSuccess
	case "9":
		statusMsg = model.StageOneFailed
	case "8":
		statusMsg = model.StageTwoFailed
	}
	return statusMsg
}
func GetInterviewDepartmentMsg(status string, department model.AuInterviewDepartment) (backMessage string) {
	switch status {
	case "1":
		backMessage = department.StageOneMessage
	case "2":
		backMessage = department.StageTwoMessage
	case "3":
		backMessage = department.StageTwoSuccessMessage
	case "9":
		backMessage = department.StageOneFailedMessage
	case "8":
		backMessage = department.StageTwoFailedMessage
	}
	return backMessage
}

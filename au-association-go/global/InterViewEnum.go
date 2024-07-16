package global

type InterViewEnum struct {
	step    int
	message string
}

var StageOneWaitB = InterViewEnum{step: 1, message: "等待社长审核"}
var StageTwoSuccess = InterViewEnum{step: 2, message: "审核通过，等待学生确认"}
var StageThreeSuccess = InterViewEnum{step: 3, message: "加入成功"}
var StageTwoFailed = InterViewEnum{step: 4, message: "审核未通过"}
var StageThreeFailed = InterViewEnum{step: 5, message: "放弃加入"}

func (interviewEnum *InterViewEnum) GetStep() int {
	return interviewEnum.step
}

func (interviewEnum InterViewEnum) GetMessageByStage(num int) string {
	return InterViewEnum{}.GetByStep(num).message
}

func (interviewEnum InterViewEnum) GetByStep(step int) InterViewEnum {
	switch step {
	case 1:
		return StageOneWaitB
	case 2:
		return StageTwoSuccess
	case 3:
		return StageThreeSuccess
	case 4:
		return StageTwoFailed
	case 5:
		return StageThreeFailed
	}
	return InterViewEnum{}
}

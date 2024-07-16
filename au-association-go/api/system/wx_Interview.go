package system

import (
	"au-golang/global"
	"au-golang/model/WxBean"
	"au-golang/model/common/response"
	"au-golang/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InterViewApi struct{}

type RequestInfo struct {
	StudentId   string `json:"studentId" form:"studentId"`
	Sex         string `json:"sex" form:"sex"`
	Description string `json:"description" form:"description"`
	Name        string `json:"name" form:"name"`
	PhoneNum    string `json:"phoneNum" form:"phoneNum"`
	WxNum       string `json:"wxNum" form:"wxNum"`
	AssId       int    `json:"assId" form:"assId"`
}

// AddInterviewUser 添加面试用户
func (interViewApi *InterViewApi) AddInterviewUser(c *gin.Context) {
	var requestInfo RequestInfo
	c.ShouldBind(&requestInfo)

	//1.没有姓名或者没有学号的错误
	if requestInfo.Name == "" || requestInfo.StudentId == "" {
		response.FailWithDetailed(gin.H{"status": 0}, "出现未知错误,请重新进入系统", c)
		return
	}

	//2.检测时间
	if !utils.CheckTime(global.GVA_CONFIG.Interview.OpenTime, global.GVA_CONFIG.Interview.EndTime) {
		Msg := "报名时间为:" + global.GVA_CONFIG.Interview.OpenTime + "到" + global.GVA_CONFIG.Interview.EndTime
		response.FailWithDetailed(gin.H{"status": 0}, Msg, c)
		return
	}

	//3.社团已满的错误
	/*	wxUser := wxUserService.FindByStudentId(requestInfo.StudentId)
		var ass1 system.Ass
		var ass2 system.Ass

		ass1 = wxUserService.FindAssByAssId(wxUser.AssEntityFirst)
		ass2 = wxUserService.FindAssByAssId(wxUser.AssEntitySecond)

		if (ass1.Assname != "" && ass1.Assid == requestInfo.AssId) || (ass2.Assname != "" && ass2.Assid == requestInfo.AssId) {
			response.FailWithDetailed(gin.H{"status": 0}, "您已经加入此社团了", c)
			return
		}
		if ass1.Assname != "" && ass2.Assname != "" {
			response.FailWithDetailed(gin.H{"status": 0}, "您已经加入了2个社团哦", c)
			return
		}*/

	//4.重复提交报名的错误
	interViewAss := interviewAssService.FindByAssId(requestInfo.AssId)
	exist := interviewUserService.FindByStudentIdAndInterViewAss(requestInfo.StudentId, interViewAss.AssId)
	if exist {
		response.FailWithDetailed(gin.H{"status": 0}, "您已经报名了此社团，请勿重复提交", c)
		return
	}

	//报名数量限制
	num := interviewUserService.FindAlByInterviewAssCount(requestInfo.StudentId, requestInfo.AssId)
	if num >= 2 {
		response.FailWithMessage("最多只能报名两个社团哦", c)
		return
	}

	new_interViewUser := WxBean.InterViewUser{Id: uuid.NewString(), StudentId: requestInfo.StudentId, Name: requestInfo.Name, Sex: requestInfo.Sex,
		Description: requestInfo.Description, InterViewAss: interViewAss, PhoneNum: requestInfo.PhoneNum, WxNum: requestInfo.WxNum, SubmitAssId: requestInfo.AssId,
	}
	new_interViewUser.ButtonControl = 1 //默认为1，使用权在社团
	new_interViewUser.BackMessage = interViewAss.ShowMessage
	new_interViewUser.InterViewStatus = global.StageOneWaitB.GetStep()
	new_interViewUser.BackMessage = global.StageOneWaitB.GetMessageByStage(new_interViewUser.InterViewStatus)
	err := interviewUserService.Create(&new_interViewUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{"status": 1, "id": new_interViewUser.Id}, "报名成功!", c)
}

// StudentConfirm 学生确认加入社团
func (interViewApi *InterViewApi) StudentConfirm(c *gin.Context) {
	id := c.Query("id")

	interViewUser, _ := interviewUserService.FindById(id)

	interViewEnum := interViewUser.InterViewStatus

	wxUser := wxUserService.FindByStudentId(interViewUser.StudentId)

	//检测社团是否已满
	if wxUser.AssEntityFirst != 0 && wxUser.AssEntitySecond != 0 {
		response.FailWithDetailed(gin.H{"status": 0}, "最只能加入2个社团", c)
		return
	}
	//检查是否加入此社团
	if wxUser.AssEntityFirst == interViewUser.SubmitAssId || wxUser.AssEntitySecond == interViewUser.SubmitAssId {
		response.FailWithMessage("您已加入过此社团", c)
		return
	}

	//判断状态是否正确
	if interViewEnum == global.StageTwoSuccess.GetStep() {
		//更改并更新面试表中的状态
		interViewEnum = global.StageThreeSuccess.GetStep()
		interViewUser.InterViewStatus = interViewEnum
		// 设置控制按钮
		interViewUser.ButtonControl = 0
		interViewUser.BackMessage = interViewUser.InterViewAss.CodeUrl + "|" + interViewUser.InterViewAss.ConfirmJoinMessage
		interviewUserService.Save(interViewUser)
		//设置学生所在社团

		if wxUser.AssEntityFirst == 0 {
			wxUserService.Save("AssEntityFirst", wxUser, interViewUser)
		} else if wxUser.AssEntitySecond == 0 {
			wxUserService.Save("AssEntitySecond", wxUser, interViewUser)
		}
		response.OkWithDetailed(gin.H{"status": 1}, "成功加入"+interViewUser.InterViewAss.AssName+"!", c)
		return
	}
	return
}

// StudentRefuse 学生拒绝加入社团
func (interViewApi *InterViewApi) StudentRefuse(c *gin.Context) {
	id := c.Query("id")

	interviewUser, _ := interviewUserService.FindById(id)
	interviewEnum := interviewUser.InterViewStatus
	if interviewEnum == global.StageTwoSuccess.GetStep() {
		interviewEnum = global.StageThreeFailed.GetStep()
		//更新状态
		interviewUser.InterViewStatus = interviewEnum
		interviewUser.BackMessage = global.StageThreeFailed.GetMessageByStage(interviewUser.InterViewStatus)
		//设置控制按钮
		interviewUser.ButtonControl = 0
		interviewUserService.Save(interviewUser)
	}
	response.Ok(c)
}

// GetUserInterviews 获取学生的面试记录
func (interViewApi *InterViewApi) GetUserInterviews(c *gin.Context) {
	studentId := c.Query("studentId")
	if len(studentId) == 0 {
		response.FailWithMessage("查找失败，请重新进入系统", c)
		return
	}

	interViewUsers := interviewUserService.FindAlByStudentId(studentId)
	var res = make([]map[string]interface{}, 0)
	for _, user := range interViewUsers {
		resInfo := WxBean.InterViewUser{}.GetInfo(user)

		if user.InterViewStatus != global.StageThreeSuccess.GetStep() {
			resInfo["BackMessage"] = user.InterViewAss.ShowMessage
		} else {
			resInfo["BackMessage"] = user.InterViewAss.ConfirmJoinMessage
		}
		resInfo["assId"] = user.InterViewAss.AssId
		res = append(res, resInfo)
	}

	response.OkWithDetailed(res, "", c)
}

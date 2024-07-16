package system

import (
	"au-golang/global"
	"au-golang/model/WxBean"
	"au-golang/model/common/response"
	"au-golang/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SysInterviewApi struct{}

type UserInterviewsInfo struct {
	Item                []WxBean.InterViewUser `json:"item"`                //当前页面的面试者
	Total               int64                  `json:"total"`               //满足条件的面试者
	TotalInterviewUsers int64                  `json:"totalInterviewUsers"` //所有面试者
}

/**
 * 获取社团的面试学生
 * @param assId
 * @param step
 * @return
 */
func (i *SysInterviewApi) GetInterviewUsers(c *gin.Context) {
	step, _ := strconv.Atoi(c.Query("interviewStep"))

	assid, _ := strconv.Atoi(c.Query("assId"))

	page, _ := strconv.Atoi(c.Query("page"))
	//每页的条数
	limit, _ := strconv.Atoi(c.Query("limit"))
	interviewUsers, totalInterviewUsers, total, err := interviewAssService.GetInterViewUsers(assid, step, page, limit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data := UserInterviewsInfo{
		Item:                interviewUsers,
		Total:               total,
		TotalInterviewUsers: totalInterviewUsers,
	}

	response.OkWithDetailed(data, "ok", c)
}

/**
 * 社长审核通过
 * @param id
 * @return
 */
func (i *SysInterviewApi) AssConfirm(c *gin.Context) {

	id := c.Query("id")
	interviewUser, _ := interviewUserService.FindById(id)

	interviewEnum := interviewUser.InterViewStatus
	if interviewEnum == global.StageOneWaitB.GetStep() {
		interviewUser.InterViewStatus = global.StageTwoSuccess.GetStep()
		interviewUser.BackMessage = global.StageTwoSuccess.GetMessageByStage(interviewUser.InterViewStatus)
		//更新控制权按钮，使用权交给学生操作
		//按钮控制，1为社团操作，2为学生操作，0为停用
		interviewUser.ButtonControl = 2
		interviewUserService.UpdateInfo(interviewUser)
	}

	//发送微信订阅消息
	openid, _ := wxUserService.GetMpOpenId(interviewUser.StudentId)
	err := utils.SendAssMessage(openid, "通过", interviewUser)
	if err != nil {
		fmt.Printf("订阅消息-社团-推送出错: %v。学生:%v %v，面试状态: %v \n", interviewUser.StudentId, interviewUser.Name, err.Error(), interviewUser.InterViewStatus)
	}
	response.OkWithMessage("ok", c)
}

/**
 * 社长拒绝
 * @param id
 * @return
 */
func (i *SysInterviewApi) AssRefuse(c *gin.Context) {
	id := c.Query("id")
	interviewUser, _ := interviewUserService.FindById(id)
	interviewEnum := interviewUser.InterViewStatus
	if interviewEnum == global.StageOneWaitB.GetStep() {
		//更新状态
		interviewUser.InterViewStatus = global.StageTwoFailed.GetStep()
		interviewUser.BackMessage = global.StageTwoFailed.GetMessageByStage(interviewUser.InterViewStatus)
		//设置控制按钮
		interviewUser.ButtonControl = 0
		if err := interviewUserService.UpdateInfo(interviewUser); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}
	//发送微信订阅消息
	openid, _ := wxUserService.GetMpOpenId(interviewUser.StudentId)
	err := utils.SendAssMessage(openid, "不通过", interviewUser)
	if err != nil {
		fmt.Printf("订阅消息-社团-推送出错: %v。学生:%v %v，面试状态: %v \n", interviewUser.StudentId, interviewUser.Name, err.Error(), interviewUser.InterViewStatus)
	}
	response.OkWithMessage("ok", c)
}

func (i *SysInterviewApi) SetAssShowMessage(c *gin.Context) {
	assId, _ := strconv.Atoi(c.Query("assId"))
	interviewAss := interviewAssService.FindByAssId(assId)
	showMessage := c.Query("showMessage")
	confirmJoinMessage := c.Query("confirmJoinMessage")
	presidentName := c.Query("presidentName")
	presidentWechat := c.Query("presidentWechat")
	interviewAss.ConfirmJoinMessage = confirmJoinMessage
	interviewAss.ShowMessage = showMessage
	interviewAss.PresidentWechat = presidentWechat
	interviewAss.PresidentName = presidentName
	interviewAssService.UpdateInfo(interviewAss)
	response.OkWithMessage("ok", c)
}

// 返回协会信息
type assInfo struct {
	ShowMessage        string
	ConfirmJoinMessage string
	PresidentWechat    string
	PresidentName      string
}

func (i *SysInterviewApi) GetAssMessage(c *gin.Context) {
	assId, _ := strconv.Atoi(c.Query("assId"))
	interviewAss := interviewAssService.FindByAssId(assId)
	ass := assInfo{
		PresidentName:      interviewAss.PresidentName,
		PresidentWechat:    interviewAss.PresidentWechat,
		ShowMessage:        interviewAss.ShowMessage,
		ConfirmJoinMessage: interviewAss.ConfirmJoinMessage,
	}
	response.OkWithDetailed(ass, "ok", c)
}

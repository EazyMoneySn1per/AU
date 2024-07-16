package system

import (
	"au-golang/global"
	"au-golang/model/common/response"
	"au-golang/model/system"
	"au-golang/model/system/Request"
	"github.com/araddon/dateparse"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
)

type ActivityApi struct{}

const FinalStepActivity int = 6

/*
Submit 提交社团活动
@param files 文件
@param description 活动简介
@param date 活动日期
@param assid 社团ID
@param name 活动名称
@param filed 文件类型，1策划书，2财务报表
@param step 当前进度
*/
func (Activity *ActivityApi) Submit(c *gin.Context) {
	activityInfo := Request.ActivityInfo{}
	c.ShouldBind(&activityInfo)

	activity, err := activityService.FindByAssidAndStepIsNot(activityInfo.Assid, FinalStepActivity)
	if err == gorm.ErrRecordNotFound { //没有数据
	} else {
		if activity.File1 != "" && activityInfo.Filed == 1 {
			os.Remove(activity.File1)
		}
		if activity.File2 != "" && activityInfo.Filed == 2 {
			os.Remove(activity.File2)
		}
	}
	if activityInfo.Filed == 1 {
		activity.File1 = Upload(c)
		activity.Name = activityInfo.Name

		time, _ := dateparse.ParseLocal(activityInfo.Date)
		activity.Date = time
		activity.Description = activityInfo.Description
	}
	if activityInfo.Filed == 2 {
		activity.File2 = Upload(c)
	}
	//直接跳过工作人员审核，即步骤变为 1->3, 3->6
	if activityInfo.Step == 1 {
		activity.Step = activityInfo.Step + 2
	} else if activity.Step == 3 {
		activity.Step = activityInfo.Step + 3
	}
	//activity.Step = activityInfo.Step + 2
	activity.Backmsg = ""
	if activity.Id == "" {
		activity.Id = uuid.NewString()[0:31]
		activity.Assid, _ = strconv.Atoi(activityInfo.Assid)
		activityService.Create(&activity)
		return
	} else {
		activityService.Save(&activity)
	}
	response.Ok(c)
}

func Upload(c *gin.Context) string {
	files, _ := c.FormFile("files")
	if files == nil {
		return ""
	}
	uuid, _ := uuid.NewUUID()
	storagePath := global.GVA_CONFIG.StoragePath.RootActivityPath + "/" + uuid.String() + "_" + files.Filename
	err := c.SaveUploadedFile(files, storagePath)
	if err != nil {

	}
	return storagePath
}

// Getinfo 当社长进入页面的时候，获取当前活动进度和失败信息
func (Activity *ActivityApi) Getinfo(c *gin.Context) {
	assid := c.Query("assid")
	activity, err := activityService.FindByAssidAndStepIsNot(assid, FinalStepActivity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(activity, "ok", c)
}

// Nextstep 审核时候，审核成功的下一步
func (Activity *ActivityApi) Nextstep(c *gin.Context) {
	assid := c.Query("assid")
	activity, _ := activityService.FindByAssidAndStepIsNot(assid, FinalStepActivity)

	tempStep := activity.Step
	activity.Step = tempStep + 1
	activityService.Save(&activity)

	if activity.Step == FinalStepActivity {
		activityComplete := system.ActivityComplete{}
		activityComplete.ActivityName = activity.Name
		activityComplete.Twitter = ""
		activityComplete.Date = activity.Date
		activityComplete.File1 = activity.File1
		activityComplete.File2 = activity.File2
		activityComplete.Description = activity.Description
		activityComplete.AssStruct = assService.FindByAssid(activity.Assid)
		activityCompleteService.Save(&activityComplete)
	}

	response.OkWithMessage("ok", c)
}

// SetBackMsg 审核失败的时候设置返回信息
func (Activity *ActivityApi) SetBackMsg(c *gin.Context) {
	assid := c.Query("assid")
	backmsg := c.Query("backmsg")
	activity, err := activityService.FindByAssidAndStepIsNot(assid, FinalStepActivity)
	if err == gorm.ErrRecordNotFound { // 没有找到数据
		response.FailWithMessage("request fail", c)
		return
	}
	activity.Backmsg = backmsg
	if activity.Step == 4 {
		activity.Step = activity.Step - 1
	} else {
		activity.Step = 1
	}
	activityService.Save(&activity)
	response.OkWithMessage("ok", c)
}

// GetActivities 获取审核列表
func (Activity *ActivityApi) GetActivities(c *gin.Context) {
	activities := activityService.FindAllByStepNot(FinalStepActivity)
	response.OkWithDetailed(activities, "ok", c)
}

// DownFile 文件上传
func (Activity *ActivityApi) DownFile(c *gin.Context) {
	path := c.Query("path")
	_, errByOpenFile := os.Open(path)
	arr := strings.Split(path, "_")
	if errByOpenFile != nil {
		response.FailWithMessage("fail", c)
		return
	}
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", "attachment; filename="+arr[1])
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(path)
	return
}

func (a *ActivityApi) GetRecord(c *gin.Context) {

	page := c.Query("page")
	limit := c.Query("limit")
	assId, exist := c.GetQuery("assId")
	if !exist {
		response.FailWithMessage("传参出错", c)
		return
	}
	activity, count := activityService.GetRecord(cast.ToInt(assId), cast.ToInt(page), cast.ToInt(limit))

	response.OkWithDetailed(gin.H{
		"data":  activity,
		"count": count,
	}, "ok", c)
}

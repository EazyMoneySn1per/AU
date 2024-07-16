package system

import (
	"au-golang/model/common/response"
	"github.com/gin-gonic/gin"
)

type DataShowApi struct{}

// GetUnfinishedEvent 获取正在审核中的活动财务报表推文
func (dataShowApi *DataShowApi) GetUnfinishedEvent(c *gin.Context) {
	activities := activityService.FindAllByStepNot(FinalStepActivity)
	twitters := twitterService.FindAllByStepNot(FinalStepTwitter)
	outlays := outlayService.FindAllByStepNot(FinalStepOutlay)

	activitiesNum := len(activities)
	twittersNum := len(twitters)
	outlaysNum := len(outlays)

	response.OkWithDetailed(gin.H{"activitiesNum": activitiesNum, "twittersNum": twittersNum, "outlaysNum": outlaysNum}, "ok", c)
}

// GetFinishedEvent 获取所有活动财务报表推文
func (dataShowApi *DataShowApi) GetFinishedEvent(c *gin.Context) {
	activityCompletes := activityCompleteService.FindAll()
	twitters := twitterService.FindAllByStep(FinalStepTwitter)
	outlays := outlayService.FindAllByStep(FinalStepOutlay)
	activityCompletesNum := len(activityCompletes)
	twittersNum := len(twitters)
	outlaysNum := len(outlays)

	response.OkWithDetailed(gin.H{"activitiesNum": activityCompletesNum, "twittersNum": twittersNum, "outlaysNum": outlaysNum}, "ok", c)
}

// GetTodayEvent 获取今日提交的活动，财务报表，推文
func (dataShowApi *DataShowApi) GetTodayEvent(c *gin.Context) {
	activities := activityService.FindTodayAll()
	twitters := twitterService.FindTodayAll()
	outlays := outlayService.FindTodayAll()

	activitiesNum := len(activities)
	twittersNum := len(twitters)
	outlaysNum := len(outlays)

	response.OkWithDetailed(gin.H{"activitiesNum": activitiesNum, "twittersNum": twittersNum, "outlaysNum": outlaysNum}, "ok", c)
}

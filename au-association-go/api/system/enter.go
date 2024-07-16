package system

import (
	api2 "au-golang/api/migrate"
	"au-golang/service"
)

type ApiGroup struct {
	UserApi
	AccountApi
	ActivityApi
	TwitterApi
	OutlayApi
	DataShowApi
	InterViewApi
	WxUserApi
	SynthesizeApi
	AdminApi
	AdministerApi
	SysInterviewApi
	api2.DepartmentApiMigrate
	api2.UserApiMigrate
}

var (
	userService             = service.ServiceGroupApp.SystemServiceGroup.UserService
	accountService          = service.ServiceGroupApp.SystemServiceGroup.AccountService
	activityService         = service.ServiceGroupApp.SystemServiceGroup.ActivityService
	assService              = service.ServiceGroupApp.SystemServiceGroup.AssService
	activityCompleteService = service.ServiceGroupApp.SystemServiceGroup.ActivityCompleteService
	twitterService          = service.ServiceGroupApp.SystemServiceGroup.TwitterService
	outlayService           = service.ServiceGroupApp.SystemServiceGroup.OutlayService
	interviewAssService     = service.ServiceGroupApp.SystemServiceGroup.InterViewAssService
	wxUserService           = service.ServiceGroupApp.SystemServiceGroup.WxUserService
	interviewUserService    = service.ServiceGroupApp.SystemServiceGroup.InterViewUserService
	synthesizeSubmitService = service.ServiceGroupApp.SystemServiceGroup.SysnthesizeSubmitService
)

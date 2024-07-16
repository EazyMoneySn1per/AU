package system

import (
	"au-golang/router/migrate"
)

type RouterGroup struct {
	UserRouter
	AccountRouter
	ActivityRouter
	WxUserRouter
	migrate.UserRouterMigrate
	migrate.DepartmentRouterMigrate
	SynthesizeRouter
	SysInterviewRouter
}

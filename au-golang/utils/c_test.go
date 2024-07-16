package utils

import (
	"au-go/global"
	"testing"
)

func Test(t *testing.T) {
	CheckTime(global.GVA_CONFIG.Interview.OpenTime, global.GVA_CONFIG.Interview.EndTime)
}

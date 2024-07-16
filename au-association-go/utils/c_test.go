package utils

import (
	"regexp"
	"testing"
)

func Test(t *testing.T) {
	query := "廖展明"
	match, _ := regexp.MatchString(`^[0-9]+$`, query)
	println(match)
	//CheckTime(global.GVA_CONFIG.Interview.OpenTime, global.GVA_CONFIG.Interview.EndTime)
}

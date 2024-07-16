package utils

import (
	"time"
)

func CheckTime(timeOne string, timeTwo string) bool {
	var LOC, _ = time.LoadLocation("Asia/Shanghai")
	to, _ := time.ParseInLocation("2006-01-02 15:04:05", timeOne, LOC)
	tw, _ := time.ParseInLocation("2006-01-02 15:04:05", timeTwo, LOC)
	now := time.Now()
	if now.After(to) && now.Before(tw) {
		return true
	}
	return false
}

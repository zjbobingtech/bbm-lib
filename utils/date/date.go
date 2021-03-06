package utils

import (
	"time"
)

func GetDuration(start time.Time, end time.Time) time.Duration {
	return end.Sub(start)
}

func GetDurationMilliseconds(start time.Time, end time.Time) int64 {
	return end.Sub(start).Milliseconds()
}

func IsNowInTimeRange(startTimeStr, endTimeStr string) bool {
	//当前时间
	now := time.Now()
	//当前时间转换为"年-月-日"的格式
	format := now.Format("2006-01-02")
	//转换为time类型需要的格式
	layout := "2006-01-02 15:04:05"
	//将开始时间拼接“年-月-日 ”转换为time类型
	timeStart, _ := time.ParseInLocation(layout, format+" "+startTimeStr, time.Local)
	//将结束时间拼接“年-月-日 ”转换为time类型
	timeEnd, _ := time.ParseInLocation(layout, format+" "+endTimeStr, time.Local)
	//使用time的Before和After方法，判断当前时间是否在参数的时间范围
	return now.Before(timeEnd) && now.After(timeStart)
}

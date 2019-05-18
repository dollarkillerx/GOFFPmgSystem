package utils

import (
	"fmt"
	"strconv"
	"time"
)

var (
	// 日期str模板
	timeLayout = "2006-01-02"
	// 时区设置
	location, _ = time.LoadLocation("Asia/Shanghai")
)

// 获取当前时间
func GetCurrentTime() string {
	unix := time.Now().In(location).Unix()
	return fmt.Sprintf("%v",unix)
}

// 时间戳转日期str
func GetTimeToString(times string) (string,error) {
	i, err := strconv.ParseInt(times, 10, 64)
	format := time.Unix(i, 0).Format(timeLayout)
	return format,err
}

// 日期str转时间戳
func GetTimeStringToTime(times string) (string,error) {
	inLocation, e := time.ParseInLocation(timeLayout, times, location)
	unix := inLocation.Unix()
	return fmt.Sprintf("%v",unix),e
}
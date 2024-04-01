/**
 * @author:       wangxuebing
 * @fileName:     timeutil.go
 * @date:         2023/1/14 16:48
 * @description:
 */

package timeutil

import (
	"chai-comutils/common/utils/characterutil"
	"fmt"
	"strings"
	"time"
)

// TimeMilliFormatDateTime 将int64毫秒时间格式化为时间字符串"YYYY-MM-DD HH:mm:ss"
func TimeMilliFormatDateTime(t int64) string {
	var timeFormat string
	if t == 0 {
		timeFormat = "-"
	} else {
		tm := time.Unix(t/1e3, 0)
		timeFormat = tm.Format(time.DateTime)
	}
	return timeFormat
}

// TimeMilliFormatDate 将int64毫秒时间格式化为日期字符串"YYYY-MM-DD"
func TimeMilliFormatDate(t *int64) string {
	if t == nil || *t == 0 {
		return "--"
	}
	return time.Unix(*t/1e3, 0).Format(time.DateOnly)
}

// TimeFormat 根据time.Time时间类型格式化为时间字符串"YYYY-MM-DD HH:mm:ss"
func TimeFormat(t *time.Time) string {
	var timeFormat string
	if t == nil {
		timeFormat = "——"
	} else {
		timeFormat = t.Format(time.DateTime)
	}
	return timeFormat
}

// TimeFormatYMD 根据time.Time时间类型格式化为日期字符串"YYYY-MM-DD"
func TimeFormatYMD(n *time.Time) string {
	var timeFormat string
	if n == nil {
		timeFormat = "—-"
	} else {
		timeFormat = n.Format(time.DateOnly)
	}
	return timeFormat
}

func FormatMillisecondsToTime(ms int64) string {
	// 将毫秒转换为秒
	seconds := ms / 1000

	// 使用Unix函数创建一个时间对象
	t := time.Unix(seconds, 0)

	// 将时分秒设置为0
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)

	// 格式化为字符串
	return t.Format("2006-01-02 15:04:05")
}

// TimeMilliFormatDateToSeparator 将int64毫秒时间格式化为日期字符串"YYYY/MM/DD"
func TimeMilliFormatDateToSeparator(t *int64) string {
	if t == nil || *t == 0 {
		return "--"
	}
	return strings.Replace(time.Unix(*t/1e3, 0).Format(time.DateOnly), "-", "/", -1)
}

// TimeConvertSeparator 将日期字符串"YYYY-MM-DD"转换为"YYYY/MM/DD"
func TimeConvertSeparator(date string) string {
	if date == "" {
		return ""
	}
	return strings.Replace(date, "-", "/", -1)
}

// SetToEndOfDay 将给定的时间戳时分秒设置为23:59:59
func SetToEndOfDay(timestamp int64) int64 {
	t := time.Unix(timestamp/1e3, 0)

	// Set time to the end of the day (23:59:59)
	endOfDay := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	return endOfDay.UnixMilli()
}

// TimeCode 根据当前时间和编码生成日期编码(例如：2023011400001)
func TimeCode(num int64) string {
	code := characterutil.StitchingBuilderStr(time.Now().Format("20060102"), fmt.Sprintf("%05d", num))
	return code
}

// DateStampByDate 日期(年月日)的时间戳
func DateStampByDate(date string) int64 {
	t, _ := time.ParseInLocation("2006/01/02", date, time.Local)
	return t.UnixMilli()
}

// SubDate 日期减天(年月日时分秒)
func SubDate(day int) (string, string) {
	return time.Now().AddDate(0, 0, -day).Format("2006/01/02"), time.Now().Format("2006/01/02")
}

// CalculateStartAndEndDates 日期减天(年月日时分秒)输出开始时间和结束时间的毫秒时间戳
func CalculateStartAndEndDates(lastDays int) (int64, int64) {
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -lastDays+1)

	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 999999999, endDate.Location())

	startTimestamp := startDate.UnixNano() / int64(time.Millisecond)
	endTimestamp := endDate.UnixNano() / int64(time.Millisecond)

	return startTimestamp, endTimestamp
}

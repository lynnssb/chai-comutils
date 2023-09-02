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
	"time"
)

// TimeMilliFormat 将int64毫秒时间格式化为时间字符串"YYYY-MM-DD HH:mm:ss"
func TimeMilliFormat(t int64) string {
	var timeFormat string
	if t == 0 {
		timeFormat = "-"
	} else {
		tm := time.Unix(t/1e3, 0)
		timeFormat = tm.Format(time.DateTime)
	}
	return timeFormat
}

// TimeFormat 根据time.Time时间类型格式化为时间字符串"YYYY-MM-DD HH:mm:ss"
func TimeFormat(n *time.Time) string {
	var timeFormat string
	if n == nil {
		timeFormat = "——"
	} else {
		timeFormat = n.Format(time.DateTime)
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

// TimeCode 根据当前时间和编码生成日期编码(例如：2023011400001)
func TimeCode(num int64) string {
	code := characterutil.StitchingBuilderStr(time.Now().Format("20060102"), fmt.Sprintf("%05d", num))
	return code
}

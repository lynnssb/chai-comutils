package statisticsutil

import (
	"github.com/shopspring/decimal"
	"time"
)

func CalculatePercentage(numerator int64, denominator int64) float64 {
	if denominator == 0 {
		return 0.00
	}
	// 使用 decimal 类型进行计算和舍入
	numeratorDecimal := decimal.NewFromInt(numerator)
	denominatorDecimal := decimal.NewFromInt(denominator)
	resultDecimal := numeratorDecimal.Div(denominatorDecimal).Round(2)
	resultFloat, _ := resultDecimal.Float64()
	return resultFloat
}

func CalculatePercentageFloat(numerator float64, denominator float64) float64 {
	if denominator == 0 {
		return 0.00
	}
	// 使用 decimal 类型进行计算和舍入
	numeratorDecimal := decimal.NewFromFloat(numerator)
	denominatorDecimal := decimal.NewFromFloat(denominator)
	resultDecimal := numeratorDecimal.Div(denominatorDecimal).Round(2)
	resultFloat, _ := resultDecimal.Float64()
	return resultFloat
}

// CalculateMoM 计算环比时间区间
func CalculateMoM(startTime int64, endTime int64) (int64, int64) {
	previousMonthStart := millisToTime(startTime).AddDate(0, -1, 0).UnixNano() / int64(time.Millisecond)
	previousMonthEnd := millisToTime(endTime).AddDate(0, -1, 0).UnixNano() / int64(time.Millisecond)
	return previousMonthStart, previousMonthEnd
}

// CalculateYoY 计算同比时间区间
func CalculateYoY(startTime int64, endTime int64) (int64, int64) {
	previousYearStart := millisToTime(startTime).AddDate(-1, 0, 0).UnixNano() / int64(time.Millisecond)
	previousYearEnd := millisToTime(endTime).AddDate(-1, 0, 0).UnixNano() / int64(time.Millisecond)
	return previousYearStart, previousYearEnd
}

// 将毫秒时间戳转换为time.Time类型
func millisToTime(millis int64) time.Time {
	return time.Unix(0, millis*int64(time.Millisecond))
}

/**
 * @author:       wangxuebing
 * @fileName:     convertortype.go
 * @date:         2023/5/20 0:49
 * @description:
 */

package convertor

import (
	"fmt"
	"strconv"
	"strings"
)

func String(val string) *string {
	return &val
}

func StringVal(val *string) string {
	if val == nil {
		return ""
	}
	return *val
}

func Int(val int) *int {
	return &val
}

func IntVal(val *int) int {
	if val == nil {
		return 0
	}
	return *val
}

func Bool(val bool) *bool {
	return &val
}

func BoolVal(val *bool) bool {
	if val == nil {
		return false
	}
	return *val
}

// ParseFloat 字符串转Float64
func ParseFloat(str string) (float64, error) {
	floatValue, err := strconv.ParseFloat(str, 64)
	return floatValue, err
}

func InterfaceToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return fmt.Sprint(v)
	// 其他类型的处理方式可以根据需要添加
	default:
		return fmt.Sprintf("%v", value)
	}
}

func InterfaceToInt(value interface{}) (int, error) {
	if intValue, ok := value.(int); ok {
		// 转换成功，返回 int 类型的值
		return intValue, nil
	}
	if floatValue, ok := value.(float64); ok {
		// 转换成功，返回 int 类型的值
		return int(floatValue), nil
	}
	if str, ok := value.(string); ok {
		// 尝试将字符串转换为 int
		intValue, err := strconv.Atoi(str)
		if err == nil {
			// 字符串转换为 int 成功
			return intValue, nil
		}
		if str == "-" {
			return 0, nil
		}
		return 0, fmt.Errorf("Failed to convert %v to int: %v", value, err)
	}
	return 0, fmt.Errorf("Failed to convert %v to int", value)
}

func InterfaceToFloat64(value interface{}) (float64, error) {
	if floatValue, ok := value.(float64); ok {
		// 转换成功，返回 float64 类型的值
		return floatValue, nil
	}

	// 尝试将字符串转换为 float64
	if str, ok := value.(string); ok {
		floatValue, err := strconv.ParseFloat(str, 64)
		if err == nil {
			// 字符串转换为 float64 成功
			return floatValue, nil
		}
	}

	// 转换失败，返回错误
	return 0, fmt.Errorf("Failed to convert %v to float64", value)
}

// StringPercentageToFloat64 字符串百分比转Float64
// 例如：将 "12.34%" 转换为 12.34
func StringPercentageToFloat64(str string) (float64, error) {
	if str == "-" {
		return 0, nil
	} else if str == "持平" {
		return 0, nil
	}
	// 移除百分号
	str = strings.TrimSuffix(str, "%")

	// 将字符串转换为 float64
	floatValue, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, fmt.Errorf("Error converting string to float64: %v", err)
	}

	return floatValue, nil
}

// ConvertToPointerArray 将 []string 转换为 []*string
func ConvertToPointerArray(strings []string) []*string {
	var pointerArray []*string

	for _, str := range strings {
		// 创建一个新的变量，将字符串的地址赋值给该变量
		tempStr := str
		pointer := &tempStr
		pointerArray = append(pointerArray, pointer)
	}

	return pointerArray
}

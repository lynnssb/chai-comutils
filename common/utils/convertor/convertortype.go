/**
 * @author:       wangxuebing
 * @fileName:     convertortype.go
 * @date:         2023/5/20 0:49
 * @description:
 */

package convertor

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

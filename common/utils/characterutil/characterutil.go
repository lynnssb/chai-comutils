/**
 * @author:       wangxuebing
 * @fileName:     characterutil.go
 * @date:         2023/1/14 16:39
 * @description:
 */

package characterutil

import (
	"strings"
)

// StitchingBuilderStr 字符串拼接
//
//	使用 strings.Builder 比使用 + 拼接字符串要快得多
func StitchingBuilderStr(args ...string) string {
	var build strings.Builder
	for _, v := range args {
		build.WriteString(v)
	}
	return build.String()
}

// StringInArray 判断字符串是否在数组中
//
//	使用循环遍历数组
func StringInArray(target string, array []string) bool {
	for _, v := range array {
		if target == v {
			return true
		}
	}
	return false
}

// StringInArrayMap 判断字符串是否在数组中
//
//	将字符串数组中的每个字符串作为 map 的 key，值可以是任何类型。然后，我们可以使用 map[key] 的形式来检查 key 是否存在于 map 中，这比使用循环遍历数组要快得多。
func StringInArrayMap(target string, array []string) bool {
	set := make(map[string]bool)
	for _, v := range array {
		set[v] = true
	}
	return set[target]
}

// StringIntersection 求两个字符串数组的交集
func StringIntersection(arr1, arr2 []string) []string {
	set := make(map[string]bool)
	for _, s := range arr1 {
		set[s] = true
	}
	var result []string
	for _, s := range arr2 {
		if set[s] {
			result = append(result, s)
		}
	}
	return result
}

// StringUnion 求两个字符串数组的并集
func StringUnion(arr1, arr2 []string) []string {
	set := make(map[string]bool)
	for _, v := range arr1 {
		set[v] = true
	}
	for _, v := range arr2 {
		set[v] = true
	}
	var result []string
	for k := range set {
		result = append(result, k)
	}
	return result
}

// StringDifference 求两个字符串数组的差集
func StringDifference(arr1, arr2 []string) []string {
	setAMap := make(map[string]bool)
	for _, v := range arr1 {
		setAMap[v] = true
	}

	result := make([]string, 0)
	for _, v := range arr2 {
		if _, ok := setAMap[v]; !ok {
			result = append(result, v)
		}
	}

	return result
}

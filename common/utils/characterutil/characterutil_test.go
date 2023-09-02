/**
 * @author:       wangxuebing
 * @fileName:     characterutil_test.go
 * @date:         2023/3/29 18:12
 * @description:
 */

package characterutil

import "testing"

func TestStitchingBuilderStr(t *testing.T) {
	args := []string{"lynn", " is ", " wang ", " xue ", " bing."}
	t.Log(StitchingBuilderStr(args...))
}

func TestStringInArrayMap(t *testing.T) {
	args := "lynn"
	arrayt := []string{"lynn", "wang", "xue", "bing"}
	arrayf := []string{"wang", "xue", "bing"}
	t.Log(StringInArrayMap(args, arrayt))
	t.Log(StringInArrayMap(args, arrayf))
}

func TestIntersection(t *testing.T) {
	arr1 := []string{"lynn", "wang", "xue", "bing"}
	arr2 := []string{"lynn", "wang", "xue", "bing", "is", "my", "name"}
	t.Log(StringIntersection(arr1, arr2))
}

func TestStringUnion(t *testing.T) {
	arr1 := []string{"lynn", "wang", "xue", "bing"}
	arr2 := []string{"lynn", "wang", "xue", "bing", "is", "my", "name"}
	t.Log(StringUnion(arr1, arr2))
}

func TestStringDifference(t *testing.T) {
	arr1 := []string{"lynn", "wang", "xue", "bing"}
	arr2 := []string{"lynn", "wang", "xue", "bing", "is", "my", "name"}
	t.Log(StringDifference(arr1, arr2))
}

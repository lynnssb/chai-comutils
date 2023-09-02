/**
 * @author:       wangxuebing
 * @fileName:     randomutil_test.go
 * @date:         2023/5/5 0:59
 * @description:
 */

package randomutil

import (
	"sort"
	"testing"
)

func TestGetRandomNumStr(t *testing.T) {
	var ss []string
	for i := 0; i < 50000; i++ {
		ss = append(ss, GetRandomNumStr(12))
	}
	sort.Strings(ss)
	t.Log(ss)

}

func TestGetRandomStr(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(GetRandomStr(32))
	}
}

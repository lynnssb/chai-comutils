/**
 * @author:       wangxuebing
 * @fileName:     excelutil_test.go
 * @date:         2023/3/29 18:37
 * @description:
 */

package excelutil

import (
	"chai-comutils/common/utils/characterutil"
	"strconv"
	"testing"
)

func TestExportExcelByMap(t *testing.T) {
	var data []map[string]interface{}
	titles := []string{"序号", "姓名", "性别", "职务"}

	for i := 0; i <= 10; i++ {
		data = append(data, map[string]interface{}{
			"code":     strconv.Itoa(i),
			"name":     characterutil.StitchingBuilderStr("小兵", strconv.Itoa(i)),
			"gender":   characterutil.StitchingBuilderStr("男", strconv.Itoa(i)),
			"position": characterutil.StitchingBuilderStr("程序员程序序员程序员", strconv.Itoa(i)),
		})
		i++
	}

	p, err := ExportExcelByMap(titles, data, "Sheet1", "./excel_file/", "test.xlsx")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(p)
}

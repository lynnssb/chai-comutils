/**
 * @author:       wangxuebing
 * @fileName:     excelutil.go
 * @date:         2023/3/29 18:34
 * @description:
 */

package excelutil

import (
	"chai-comutils/common/utils/characterutil"
	"chai-comutils/common/utils/ioutil"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"sort"
)

// Excel表头标识
var excelCode = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I",
	"J", "K", "L", "M", "N", "O", "P", "Q", "R",
	"S", "T", "U", "V", "W", "X", "Y", "Z",
	"AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI",
	"AJ", "AK", "AL", "AM", "AN", "AO", "AP", "AQ", "AR",
	"AS", "AT", "AU", "AV", "AW", "AX", "AY", "AZ", "BA",
	"BB", "BC", "BD", "BE", "BF", "BG", "BH", "BI", "BJ",
	"BK", "BL", "BM", "BN", "BO", "BP", "BQ", "BR", "BS",
	"BT", "BU", "BV", "BW", "BX", "BY", "BZ"}

// ImportExcel 导入Excel读取数据
func ImportExcel(filePath string, sheets []string) ([]interface{}, error) {
	var dataList []interface{}
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("读取文件出错:%s", err.Error()))
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(fmt.Sprintf("读取文件出错:%s", err.Error()))
		}
	}()
	for _, sheet := range sheets {
		//获取sheet$$上所有单元格数据
		rows, err := f.GetRows(sheet)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("读取文件出错:%s", err.Error()))
		}
		dataList = append(dataList, rows)
	}

	return dataList, nil
}

// ExportExcelByMap 导出数据
func ExportExcelByMap(titles []string, data []map[string]interface{}, sheetName, filePath, fileName string) (*string, error) {
	f := excelize.NewFile()
	f.SetSheetName("Sheet1", sheetName)
	header := make([]string, 0)
	for _, v := range titles {
		header = append(header, v)
	}
	// 设置表头样式
	rowTitleStyleID, _ := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 2},
			{Type: "right", Color: "#000000", Style: 2},
			{Type: "bottom", Color: "#000000", Style: 2},
			{Type: "top", Color: "#000000", Style: 2},
		},
		Fill:      excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{"#808080"}, Shading: 1},
		Font:      &excelize.Font{Bold: true, Family: "宋体", Size: 14, Color: "#000000"},
		Alignment: &excelize.Alignment{Horizontal: "center" /*水平对齐*/, Vertical: "center" /*垂直对齐*/, WrapText: true},

		Protection: &excelize.Protection{
			Hidden: true,
			Locked: true,
		},
		Lang: "zh-cn",
	})
	// 设置表头
	_ = f.SetSheetRow(sheetName, "A1", &header)
	_ = f.SetRowHeight(sheetName, 1, 15)
	_ = f.SetCellStyle(sheetName, "A1", fmt.Sprintf("%s1", excelCode[len(header)-1]), rowTitleStyleID)

	// 设置表格数据样式
	rowStyleID, _ := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 2},
			{Type: "right", Color: "#000000", Style: 2},
			{Type: "bottom", Color: "#000000", Style: 2},
			{Type: "top", Color: "#000000", Style: 2},
		},
		Font:      &excelize.Font{Family: "宋体", Size: 12},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: false},
	})

	if err := f.SetColWidth(sheetName, "A", excelCode[len(header)-1], 25); err != nil {
		return nil, err
	}

	rowNum := 2
	for _, value := range data {
		row := make([]interface{}, 0)
		var dataSlice []string
		for key := range value {
			dataSlice = append(dataSlice, key)
		}
		sort.Strings(dataSlice)
		for _, v := range dataSlice {
			if val, ok := value[v]; ok {
				row = append(row, val)
			}
		}
		if err := f.SetSheetRow(sheetName, fmt.Sprintf("A%d", rowNum), &row); err != nil {
			return nil, err
		}

		if rowNum > 1 {
			if err := f.SetCellStyle(sheetName, fmt.Sprintf("A%d", rowNum), fmt.Sprintf("%s%d", excelCode[len(header)-1], rowNum), rowStyleID); err != nil {
				return nil, err
			}
		}
		rowNum++
	}
	if err := ioutil.PathExist(filePath); err != nil {
		return nil, err
	}
	fileFullPath := characterutil.StitchingBuilderStr(filePath, fileName)
	if err := f.SaveAs(fileFullPath); err != nil {
		return nil, err
	}

	result := fileFullPath[2:]

	return &result, nil
}

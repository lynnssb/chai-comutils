package fileutil

import (
	"os"
	"strings"
)

// PathExists 判断文件夹是否存在,不存在则创建
func PathExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err == nil {
			return nil
		} else {
			return err
		}
	}
	return err
}

// 检查文件格式
// fileName 文件名
// allowedFormats 允许的文件格式
// return bool
func checkFileFormat(fileName string, allowedFormats []string) bool {
	//获取文件后缀
	fileExt := strings.ToLower(fileName[strings.LastIndex(fileName, ".")+1:])
	for _, allowedFormat := range allowedFormats {
		if fileExt == allowedFormat {
			return true
		}
	}
	return false
}

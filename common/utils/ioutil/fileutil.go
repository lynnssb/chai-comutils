/**
 * @author:       wangxuebing
 * @fileName:     ioutil.go
 * @date:         2023/1/14 16:54
 * @description:
 */

package ioutil

import (
	"os"
)

// PathExist 判断文件夹是否存在,不存在则创建
func PathExist(path string) error {
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

// DelByFilePath 删除文件
func DelByFilePath(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}

// IsPathExist 判断路径文件/文件夹是否存在
func IsPathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// IsDir 判断路径是否是文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断路径是否是文件
func IsFile(path string) bool {
	return !IsDir(path)
}

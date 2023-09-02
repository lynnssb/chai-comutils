/**
 * @author:       wangxuebing
 * @fileName:     gormcurd.go
 * @date:         2023/5/24 0:15
 * @description:
 */

package gormtpl

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"text/template"
)

type TableInfo struct {
	Table       string //表名
	ProjectName string //项目名
}

func GormCurdGenerationCode(tmplPath string, tableInfos []TableInfo, outPutFilePath string) (bool, error) {
	for _, item := range tableInfos {
		var (
			resultBuf          bytes.Buffer
			tpl                *template.Template
			f                  *os.File
			outPutFileNamePath string
			err                error
		)
		tpl, err = template.ParseFiles(tmplPath)
		if err != nil {
			return false, errors.New("加载表失败...")
		}

		err = tpl.Execute(&resultBuf, item)
		if err != nil {
			return false, errors.New("模板生成失败...")
		}
		outPutFileNamePath = outPutFilePath + strings.ToLower(item.Table) + ".go"
		if _, err = os.Stat(outPutFileNamePath); err != nil {
			if os.IsNotExist(err) {
				if err = os.MkdirAll(outPutFilePath, os.ModePerm); err != nil {
					return false, errors.New("文件夹创建失败")
				}
				f, err = os.Create(outPutFileNamePath)
				if err != nil {
					return false, errors.New("文件创建失败")
				} else {
					if _, err := f.Write(resultBuf.Bytes()); err != nil {
						return false, errors.New("文件写入失败")
					}
					continue
				}
			}
		} else {
			return false, errors.New("文件已存在")
		}
	}

	return true, nil
}

func GormCurdCacheGenerationCode() {

}

/**
 * @author:       wangxuebing
 * @fileName:     gormcurd_test.go
 * @date:         2023/5/24 11:55
 * @description:
 */

package gormtpl

import (
	"log"
	"testing"
)

func TestGormCurdCacheGenerationCode(t *testing.T) {
	tmplPath := "./gorm_curd.tmpl"
	tables := []TableInfo{
		{Table: "NoticeUserLink", ProjectName: "chai-hotel"},
	}
	outPutFilePath := "./model/"
	success, err := GormCurdGenerationCode(tmplPath, tables, outPutFilePath)
	if success {
		log.Println("代码文件生成成功...")
	} else {
		log.Println(err.Error())
	}
}

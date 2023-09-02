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
		//{Table: "User", ProjectName: "chai-smart-pavilion"},
		//{Table: "ClientUser", ProjectName: "chai-smart-pavilion"},
		//{Table: "Role", ProjectName: "chai-smart-pavilion"},
		//{Table: "Menu", ProjectName: "chai-smart-pavilion"},
		//{Table: "UserRole", ProjectName: "chai-smart-pavilion"},
		//{Table: "RolePermission", ProjectName: "chai-smart-pavilion"},
		//{Table: "Area", ProjectName: "chai-smart-pavilion"},
		//{Table: "AreaDeviceLink", ProjectName: "chai-smart-pavilion"},
		//{Table: "Device", ProjectName: "chai-smart-pavilion"},
		//{Table: "Media", ProjectName: "chai-smart-pavilion"},
		//{Table: "DeviceMediaLink", ProjectName: "chai-smart-pavilion"},
		//{Table: "SerialDevice", ProjectName: "chai-smart-pavilion"},
		//{Table: "RelayDevice", ProjectName: "chai-smart-pavilion"},
		//{Table: "User", ProjectName: "chai-smart-touch-pavilion"},
		//{Table: "ClientUser", ProjectName: "chai-smart-touch-pavilion"},
		//{Table: "Device", ProjectName: "chai-smart-touch-pavilion"},
		//{Table: "Media", ProjectName: "chai-smart-touch-pavilion"},
		//{Table: "DeviceMediaLink", ProjectName: "chai-smart-touch-pavilion"},
		//{Table: "DeviceMediaGroup", ProjectName: "chai-smart-touch-pavilion"},
		//{Table: "DeviceMediaGroupLink", ProjectName: "chai-smart-touch-pavilion"},
		//{Table: "Dictionary", ProjectName: "chai-smart-touch-pavilion"},
		{Table: "User", ProjectName: "chai-smart-touch-pavilion"},
		{Table: "Role", ProjectName: "chai-smart-touch-pavilion"},
		{Table: "Menu", ProjectName: "chai-smart-touch-pavilion"},
		{Table: "UserRole", ProjectName: "chai-smart-touch-pavilion"},
	}
	outPutFilePath := "./model/"
	success, err := GormCurdGenerationCode(tmplPath, tables, outPutFilePath)
	if success {
		log.Println("代码文件生成成功...")
	} else {
		log.Println(err.Error())
	}
}

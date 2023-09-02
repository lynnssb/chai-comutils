/**
 * @Author:      wangxuebing
 * @FileName:    idcardutil_test.go
 * @Date         2023/7/14 14:49
 * @Description:
 **/
package idcardutil

import (
	"fmt"
	idvalidator "github.com/guanguans/id-validator"
	"testing"
)

func TestIsValid(t *testing.T) {
	idCard := "530322199208021098"
	fmt.Println(IsValid(idCard))
	fmt.Println(idvalidator.IsValid("500154199301135886", true))
	//if !IsValid(idCard) {
	//	t.Error("TestIsValid error")
	//}
}

func TestGetIdCardInfo(t *testing.T) {
	idCard := "530322199208021097"
	info, err := GetIdCardInfo(idCard)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(info.Sex)
}

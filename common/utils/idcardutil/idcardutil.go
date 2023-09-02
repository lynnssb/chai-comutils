/**
 * @Author:      wangxuebing
 * @FileName:    idcardutil.go
 * @Date         2023/7/14 14:46
 * @Description:
 **/
package idcardutil

import (
	idvalidator "github.com/guanguans/id-validator"
)

func IsValid(idCard string) bool {
	return idvalidator.IsValid(idCard, false)
}

func GetIdCardInfo(idCard string) (idInfo idvalidator.IdInfo, err error) {
	idInfo, err = idvalidator.GetInfo(idCard, false)
	return idInfo, err
}

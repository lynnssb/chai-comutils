/**
 * @author:       wangxuebing
 * @fileName:     comerror.go
 * @date:         2023/5/27 21:01
 * @description:
 */

package common

type ComError struct {
	ErrCode int    `json:"errcode"` //错误码
	ErrMsg  string `json:"errmsg"`  //错误描述
}

/**
 * @author:       wangxuebing
 * @fileName:     errorx.go
 * @date:         2023/1/14 17:10
 * @description:
 */

package errorx

import (
	"chai-comutils/common/utils/characterutil"
	"errors"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

const (
	defaultCode                   = 1001 //默认错误
	defaultGrpcCode               = 1002 //Grpc错误
	databaseErrCode               = 1003 //数据库查询错误
	databaseErrRecordNotFoundCode = 1004 //数据库查询数据为空错误
)

type CodeErr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) error {
	return &CodeErr{
		Code: code,
		Msg:  msg,
	}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func NewDatabaseError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NewCodeError(databaseErrRecordNotFoundCode, "查询[Database]数据为nil")
	}
	return NewCodeError(databaseErrCode, characterutil.StitchingBuilderStr("操作失败[Database Error]", err.Error()))
}

func NewGrpcDefaultError(err error) error {
	return NewCodeError(defaultGrpcCode, status.Convert(err).Message())
}

func (e *CodeErr) Error() string {
	return e.Msg
}

func (e *CodeErr) GetErrCode() int {
	return e.Code
}

func (e *CodeErr) Data() *CodeErr {
	return &CodeErr{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

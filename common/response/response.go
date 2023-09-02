/**
 * @author:       wangxuebing
 * @fileName:     response.go
 * @date:         2023/1/14 16:46
 * @description:
 */

package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Response 统一接口返回参数
func Response(w http.ResponseWriter, data interface{}, err error) {
	var resp Resp
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
	} else {
		resp.Msg = "OK"
		resp.Data = data
	}
	httpx.OkJson(w, resp)
}

//自定义错误处理
//httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
//switch e := err.(type) {
//case *errorx.CodeErr:
//return http.StatusOK, e.Data()
//default:
//return http.StatusInternalServerError, nil
//}
//})

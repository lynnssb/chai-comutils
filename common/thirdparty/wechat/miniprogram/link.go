/**
 * @author:       wangxuebing
 * @fileName:     link.go
 * @date:         2023/5/19 17:43
 * @description:
 */

package miniprogram

import (
	"chai-comutils/common/thirdparty/wechat/common"
	"chai-comutils/common/utils/netutil"
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
)

const (
	shortLinkUrl = "https://api.weixin.qq.com/wxa/genwxashortlink"
)

type (
	ShortLinkReq struct {
		PageUrl     string `json:"page_url"`     //通过 Short Link 进入的小程序页面路径，必须是已经发布的小程序存在的页面，可携带 query，最大1024个字符
		PageTitle   string `json:"page_title"`   //页面标题，不能包含违法信息，超过20字符会用... 截断代替
		IsPermanent bool   `json:"is_permanent"` //默认值false。生成的 Short Link 类型，短期有效：false，永久有效：true
	}
	ShortLinkResp struct {
		common.ComError
		Link string `json:"link"` //生成的小程序 Short Link
	}
)

/**
 * 获取小程序 Short Link
 * @param accessToken 接口调用凭证
 * @param reqParam ShortLinkReq
 * @return *ShortLinkResp
 * @return error
 */
func GetShortLink(accessToken string, reqParam *ShortLinkReq) (*ShortLinkResp, error) {
	reqFormParam := map[string]string{"access_token": accessToken}
	reqData, _ := json.Marshal(reqParam)
	apiUrl, err := netutil.EncodeURL(shortLinkUrl, reqFormParam)
	if err != nil {
		return nil, err
	}

	req := &fasthttp.Request{}
	req.SetRequestURI(apiUrl)
	req.SetBody(reqData)

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, errors.New("http status error")
	}

	res := new(ShortLinkResp)
	if err := json.Unmarshal(resp.Body(), res); err != nil {
		return nil, err
	}
	if res.ErrCode != 0 {
		return nil, errors.New(res.ErrMsg)
	}

	return res, nil
}

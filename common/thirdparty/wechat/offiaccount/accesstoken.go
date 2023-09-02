/**
 * @author:       wangxuebing
 * @fileName:     accesstoken.go
 * @date:         2023/5/27 21:30
 * @description:
 */

package offiaccount

import (
	"chai-comutils/common/thirdparty/wechat/common"
	"chai-comutils/common/utils/netutil"
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
)

const (
	tokenUrl       = "https://api.weixin.qq.com/cgi-bin/token"
	stableTokenUrl = "https://api.weixin.qq.com/cgi-bin/stable_token"
)

type (
	stableAccessTokenReq struct {
		GrantType    string `json:"grant_type"`    //填写 client_credential
		AppId        string `json:"appid"`         //小程序唯一凭证
		Secret       string `json:"secret"`        //小程序唯一凭证密钥
		ForceRefresh bool   `json:"force_refresh"` //是否强制刷新token
	}

	AccessTokenResp struct {
		common.ComError
		AccessToken string `json:"access_token,omitempty"` //获取到的凭证
		ExpiresIn   int    `json:"expires_in,omitempty"`   //凭证有效时间，单位：秒。目前是7200秒之内的值。
	}
)

/**
 * 获取调用凭据
 * @param appId  服务号唯一凭证
 * @param secret 服务号唯一凭证密钥
 * @return *AccessTokenResp
 * @return error
 */
func GetAccessToken(appId, secret string) (*AccessTokenResp, error) {
	req := map[string]string{
		"grant_type": "client_credential",
		"appid":      appId,
		"secret":     secret,
	}
	apiUrl, err := netutil.EncodeURL(tokenUrl, req)
	if err != nil {
		return nil, err
	}

	status, resp, err := fasthttp.Get(nil, apiUrl)
	if err != nil {
		return nil, err
	}
	if status != fasthttp.StatusOK {
		return nil, errors.New("Http请求出错")
	}

	res := new(AccessTokenResp)
	err = json.Unmarshal(resp, res)
	if err != nil {
		return nil, err
	}
	if res.ErrCode != 0 {
		return nil, errors.New(res.ErrMsg)
	}

	return res, nil
}

/*
 * 获取稳定的access_token
 * @param appId 小程序唯一凭证
 * @param secret 小程序唯一凭证密钥
 * @param forceRefresh 是否强制刷新token
 * @return *AccessTokenResp
 * @return error
 */
func GetStableAccessToken(appId, secret string, forceRefresh bool) (*AccessTokenResp, error) {
	reqParam := stableAccessTokenReq{
		GrantType:    "client_credential",
		AppId:        appId,
		Secret:       secret,
		ForceRefresh: forceRefresh,
	}
	reqData, _ := json.Marshal(reqParam)

	req := &fasthttp.Request{}
	req.SetRequestURI(stableTokenUrl)
	req.SetBodyRaw(reqData)

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, errors.New("Http请求出错")
	}

	res := new(AccessTokenResp)
	if err := json.Unmarshal(resp.Body(), res); err != nil {
		return nil, err
	}
	if res.ErrCode != 0 {
		return nil, errors.New(res.ErrMsg)
	}

	return res, nil
}

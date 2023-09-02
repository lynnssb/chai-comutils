/**
 * @author:       wangxuebing
 * @fileName:     code2session.go
 * @date:         2023/5/19 15:23
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
	loginUrl = "https://api.weixin.qq.com/sns/jscode2session"
)

type (
	CodeSessionResp struct {
		common.ComError
		SessionKey string `json:"session_key"` //会话密钥
		UnionId    string `json:"union_id"`    //用户在开放平台的唯一标识符
		OpenId     string `json:"openid"`      //用户唯一标识
	}

	/**
	 * 微信小程序接口登录凭据校验错误码
	 *  40029: code 无效                                                       js_code无效
	 *  45011: api minute-quota reach limit  mustslower  retry next minute    API 调用太频繁，请稍候再试
	 *  40226: code blocked                                                   高风险等级用户，小程序登录拦截 。风险等级详见用户安全解方案
	 *  -1:    system error                                                   系统繁忙，此时请开发者稍候再试
	 *
	 */
)

/**
 * 登录凭据校验
 * @param appid 小程序唯一凭证
 * @param secret 小程序唯一凭证密钥
 * @param jsCode 登录时获取的 code
 * @return *LoginResp
 * @return error
 */
func Login(appid, secret, jsCode string) (*CodeSessionResp, error) {
	req := map[string]string{
		"appid":      appid,
		"secret":     secret,
		"js_code":    jsCode,
		"grant_type": "authorization_code",
	}
	apiUrl, err := netutil.EncodeURL(loginUrl, req)
	if err != nil {
		return nil, err
	}
	status, resp, err := fasthttp.Get(nil, apiUrl)
	if err != nil {
		return nil, err
	}
	if status != fasthttp.StatusOK {
		return nil, errors.New("http status error")
	}

	res := new(CodeSessionResp)
	err = json.Unmarshal(resp, res)
	if err != nil {
		return nil, err
	}
	if res.ErrCode != 0 {
		return nil, errors.New(res.ErrMsg)
	}
	return res, nil
}

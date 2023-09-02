/**
 * @author:       wangxuebing
 * @fileName:     openplatform.go
 * @date:         2023/5/27 21:00
 * @description:  微信开放平台相关API接口
 */

package openplatform

import (
	"chai-comutils/common/thirdparty/wechat/common"
	"chai-comutils/common/utils/netutil"
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
)

const (
	tokenUrl            = "https://api.weixin.qq.com/sns/oauth2/access_token"
	refreshTokenUrl     = "https://api.weixin.qq.com/sns/oauth2/refresh_token"
	checkAccessTokenUrl = "https://api.weixin.qq.com/sns/auth"
	userInfoUrl         = "https://api.weixin.qq.com/sns/userinfo"
)

type (
	AccessTokenResp struct {
		common.ComError
		AccessToken  string `json:"access_token,omitempty"`  //获取到的凭证
		ExpiresIn    int    `json:"expires_in,omitempty"`    //凭证有效时间，单位：秒。目前是7200秒之内的值。
		RefreshToken string `json:"refresh_token,omitempty"` //用户刷新access_token
		OpenID       string `json:"openid,omitempty"`        //授权用户唯一标识
		Scope        string `json:"scope,omitempty"`         //用户授权的作用域,使用(,)分割
		unionID      string `json:"union_id,omitempty"`      //用户统一标，针对一个微信开放平台帐号下的应用，同一用户的 unionid 是唯一的
	}

	UserInfoResp struct {
		common.ComError
		OpenID     string   `json:"openid,omitempty"`     //普通用户标识
		NickName   string   `json:"nickname,omitempty"`   //用户昵称
		Sex        int      `json:"sex,omitempty"`        //用户性别[1:男;2:女]
		Province   string   `json:"province,omitempty"`   //省份
		City       string   `json:"city,omitempty"`       //城市
		Country    string   `json:"country,omitempty"`    //国家
		HeadImgUrl string   `json:"headimgurl,omitempty"` //头像
		Privilege  []string `json:"privilege,omitempty"`  //用户特权信息JSON数组
		UnionID    string   `json:"unionid,omitempty"`    //用户统一标识
	}
)

/**
 * 根据code获取access_token
 * @param appId 小程序唯一凭证
 * @param secret 小程序唯一凭证密钥
 * @param code
 * @return *AccessTokenResp
 * @return error
 */
func GetAccessToke(appID, secret string, code string) (*AccessTokenResp, error) {
	req := map[string]string{
		"appid":      appID,
		"secret":     secret,
		"code":       code,
		"grant_type": "authorization_code",
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

/**
 * 刷新或续期access_token使用
 * @param appID
 * @param refreshToken
 * @return *AccessTokenResp
 * @return error
 */
func GetRefreshToken(appID string, refreshToken string) (*AccessTokenResp, error) {
	req := map[string]string{
		"appid":         appID,
		"refresh_token": refreshTokenUrl,
		"grant_type":    "refresh_token",
	}

	apiUrl, err := netutil.EncodeURL(refreshTokenUrl, req)
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

/**
 * 检验授权凭证（access_token）是否有效
 * @param accessToken
 * @openID
 * @return bool (true:有效;false:无效)
 * @return error
 */
func CheckAccessToken(accessToken, openID string) (bool, error) {
	req := map[string]string{
		"access_token": accessToken,
		"openid":       openID,
	}

	apiUrl, err := netutil.EncodeURL(checkAccessTokenUrl, req)
	if err != nil {
		return false, err
	}

	status, resp, err := fasthttp.Get(nil, apiUrl)
	if err != nil {
		return false, err
	}
	if status != fasthttp.StatusOK {
		return false, errors.New("Http请求出错")
	}

	res := new(common.ComError)
	err = json.Unmarshal(resp, res)
	if err != nil {
		return false, err
	}
	if res.ErrMsg == "ok" {
		return true, nil
	} else {
		return false, errors.New(res.ErrMsg)
	}
}

/**
 * 获取用户个人信息（UnionID机制）
 * @param accessToken
 * @param openID
 * @param lang (zh_CN:简体;zh_TW:繁体;en:英语)
 * @return *UserInfoResp
 * @return error
 */
func GetUserInfo(accessToken string, openID string, land string) (*UserInfoResp, error) {
	req := map[string]string{
		"access_token": accessToken,
		"openid":       openID,
		"lang":         land,
	}

	apiUrl, err := netutil.EncodeURL(userInfoUrl, req)
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

	res := new(UserInfoResp)
	err = json.Unmarshal(resp, res)
	if err != nil {
		return nil, err
	}
	if res.ErrCode != 0 {
		return nil, errors.New(res.ErrMsg)
	}
	return res, nil
}

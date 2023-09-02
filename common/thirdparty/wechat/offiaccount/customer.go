/**
 * @author:       wangxuebing
 * @fileName:     customer.go
 * @date:         2023/5/27 21:35
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
	kfListUrl       = "https://api.weixin.qq.com/cgi-bin/customservice/getkflist"
	onlineKfListUrl = "https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist"
	kfAddUrl        = "https://api.weixin.qq.com/customservice/kfaccount/add"
	inviteKfUrl     = "https://api.weixin.qq.com/customservice/kfaccount/inviteworker"
	updateKfUrl     = "https://api.weixin.qq.com/customservice/kfaccount/update"
	kfHeadImgUrl    = "https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg"
	delKfUrl        = "https://api.weixin.qq.com/customservice/kfaccount/del"

	createSessionUrl  = "https://api.weixin.qq.com/customservice/kfsession/create"
	closeSessionUrl   = "https://api.weixin.qq.com/customservice/kfsession/close"
	getSessionUrl     = "https://api.weixin.qq.com/customservice/kfsession/getsession"
	getSessionListUrl = "https://api.weixin.qq.com/customservice/kfsession/getsessionlist"
	getWaitCaseUrl    = "https://api.weixin.qq.com/customservice/kfsession/getwaitcase"

	getMsgUrl = "https://api.weixin.qq.com/customservice/msgrecord/getmsglist"
)

type (
	KfListData struct {
		KfAccount        string `json:"kf_account"`         //完整客服帐号
		KfNick           string `json:"kf_nick"`            //客服昵称
		KfId             string `json:"kf_id"`              //客服编号
		KfHeadimgurl     string `json:"kf_headimgurl"`      //客服头像
		KfWx             string `json:"kf_wx"`              //如果客服帐号已绑定了客服人员微信号， 则此处显示微信号
		InviteWx         string `json:"invite_wx"`          //如果客服帐号尚未绑定微信号，但是已经发起了一个绑定邀请， 则此处显示绑定邀请的微信号
		InviteExpireTime int64  `json:"invite_expire_time"` //如果客服帐号尚未绑定微信号，但是已经发起过一个绑定邀请， 邀请的过期时间，为unix 时间戳
		InviteStatus     string `json:"invite_status"`      //邀请的状态，有等待确认“waiting”，被拒绝“rejected”， 过期“expired”
	}
	KfListResp struct {
		common.ComError
		KfList []KfListData `json:"kf_list,omitempty"`
	}
	OnlineKfListData struct {
		KfAccount    string `json:"kf_account"`    //完整客服帐号，格式为：帐号前缀@公众号微信号
		Status       int    `json:"status"`        //客服在线状态，目前为：1、web 在线
		KfId         string `json:"kf_id"`         //客服编号
		AcceptedCase int    `json:"accepted_case"` //客服当前正在接待的会话数
	}
	OnlineKfListResp struct {
		common.ComError
		KfOnlineList []OnlineKfListData `json:"kf_online_list,omitempty"`
	}

	KfAddReq struct {
		KfAccount string `json:"kf_account"` //完整客服帐号，格式为：帐号前缀@公众号微信号，帐号前缀最多10个字符，必须是英文、数字字符或者下划线，后缀为公众号微信号，长度不超过30个字符
		Nickname  string `json:"nickname"`   //客服昵称，最长16个字
	}
)

/**
 * 获取客服基本信息
 * @param accessToken
 * @return *KfListResp
 * @return error
 */
func GetKfList(accessToken string) (*KfListResp, error) {
	req := map[string]string{
		"access_token": accessToken,
	}
	apiUrl, err := netutil.EncodeURL(kfListUrl, req)
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

	res := new(KfListResp)
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
 * 获取在线客服基本信息
 * @param accessToken
 * @return *OnlineKfListResp
 * @return error
 */
func GetOnlineKfList(accessToken string) (*OnlineKfListResp, error) {
	req := map[string]string{
		"access_token": accessToken,
	}
	apiUrl, err := netutil.EncodeURL(onlineKfListUrl, req)
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

	res := new(OnlineKfListResp)
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
 * 添加客服
 * @param accessToken
 * @param kfAccount
 * @param nickName
 * @return error
 */
func AddKf(accessToken string, kfAccount, nickName string) error {
	reqQueryParam := map[string]string{
		"access_token": accessToken,
	}
	reqParam := KfAddReq{
		KfAccount: kfAccount,
		Nickname:  nickName,
	}

	reqData, _ := json.Marshal(reqParam)

	apiUrl, err := netutil.EncodeURL(kfAddUrl, reqQueryParam)
	if err != nil {
		return err
	}

	req := &fasthttp.Request{}
	req.SetRequestURI(apiUrl)
	req.SetBodyRaw(reqData)

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if err := fasthttp.Do(req, resp); err != nil {
		return err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return errors.New("Http请求出错")
	}

	res := new(common.ComError)
	if err := json.Unmarshal(resp.Body(), res); err != nil {
		return err
	}
	if res.ErrCode != 0 {
		return errors.New(res.ErrMsg)
	}

	return nil
}

/**
 * 邀请绑定客服账号
 * @param accessToken
 * @param kfAccount
 * @param inviteWx
 * @return error
 */
func InviteKf(accessToken string, kfAccount, inviteWx string) error {
	return nil
}

/**
 * 设置客服信息
 * @param accessToken
 * @param kfAccount
 * @param nickName
 * @return error
 */
func SetKfAccount(accessToken string, kfAccount, nickName string) error {
	return nil
}

/**
 * 上传客服头像
 * @param accessToken
 * @param media
 * @return error
 */
func UploadHeadImg(accessToken string, media string) error {
	return nil
}

/**
 * 删除客服账号
 * @param accessToken
 * @param kfAccount
 * @return error
 */
func DeleteKf(accessToken string, kfAccount string) error {
	req := map[string]string{
		"access_token": accessToken,
		"kf_account":   kfAccount,
	}
	apiUrl, err := netutil.EncodeURL(onlineKfListUrl, req)
	if err != nil {
		return err
	}

	status, resp, err := fasthttp.Get(nil, apiUrl)
	if err != nil {
		return err
	}
	if status != fasthttp.StatusOK {
		return errors.New("Http请求出错")
	}

	res := new(common.ComError)
	err = json.Unmarshal(resp, res)
	if err != nil {
		return err
	}
	if res.ErrCode != 0 {
		return errors.New(res.ErrMsg)
	}

	return nil
}

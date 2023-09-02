/**
 * @author:       wangxuebing
 * @fileName:     contentseccheck.go
 * @date:         2023/5/19 17:49
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
	msgSecCheckUrl   = "https://api.weixin.qq.com/wxa/msg_sec_check"
	mediaSecCheckUrl = "https://api.weixin.qq.com/wxa/media_check_async"
)

type (
	MsgSecCheckReq struct {
		Content   string `json:"content"`             //需检测的文本内容，文本字数的上限为2500字，需使用UTF-8编码
		Version   string `json:"version"`             //接口版本号，2.0版本为固定值2
		Scene     int    `json:"scene"`               //场景枚举值（1 资料；2 评论；3 论坛；4 社交日志）
		OpenId    string `json:"openid"`              //用户的openid（用户需在近两小时访问过小程序）
		Title     string `json:"title,omitempty"`     //文本标题，需使用UTF-8编码
		NickName  string `json:"nickname,omitempty"`  //用户昵称，需使用UTF-8编码
		Signature string `json:"signature,omitempty"` //个性签名，该参数仅在资料类场景有效(scene=1)，需使用UTF-8编码
	}

	MsgSecCheckDetailData struct {
		Strategy string `json:"strategy"` //策略类型
		ErrCode  int    `json:"errcode"`  //错误码，仅当该值为0时，该项结果有效
		Suggest  string `json:"suggest"`  //建议，有risky、pass、review三种值
		Label    int    `json:"label"`    //命中标签枚举值，100 正常；10001 广告；20001 时政；20002 色情；20003 辱骂；20006 违法犯罪；20008 欺诈；20012 低俗；20013 版权；21000 其他
		Keyword  string `json:"keyword"`  //命中的自定义关键词
		Prob     int    `json:"prob"`     //0-100，代表置信度，越高代表越有可能属于当前返回的标签（label）
	}
	MsgSecCheckResultData struct {
		Suggest string `json:"suggest"` //建议，有risky、pass、review三种值
		Label   int    `json:"label"`   //命中标签枚举值，100 正常；10001 广告；20001 时政；20002 色情；20003 辱骂；20006 违法犯罪；20008 欺诈；20012 低俗；20013 版权；21000 其他
	}
	MsgSecCheckResp struct {
		common.ComError
		Detail  []MsgSecCheckDetailData `json:"detail"`   //检测结果详情
		TraceId string                  `json:"trace_id"` //唯一请求标识，标记单次请求
		Result  MsgSecCheckResultData   `json:"result"`   //综合结果
	}

	MediaSecCheckReq struct {
		MediaUrl  string `json:"media_url"`  //要检测的图片或音频的url，支持图片格式包括jpg, jepg, png, bmp, gif（取首帧），支持的音频格式包括mp3, aac, ac3, wma, flac, vorbis, opus, wav
		MediaType int    `json:"media_type"` //1:音频;2:图片
		Version   int    `json:"version"`    //接口版本号，2.0版本为固定值2
		Scene     int    `json:"scene"`      //场景枚举值（1 资料；2 评论；3 论坛；4 社交日志）
		OpenId    string `json:"openid"`     //用户的openid（用户需在近两小时访问过小程序）
	}

	MediaSecCheckResp struct {
		common.ComError
		TraceId string `json:"trace_id"` //唯一请求标识，标记单次请求，用于匹配异步推送结果
	}

	// MediaSecCheckAsyncResp 异步检测结果推送
	MediaSecCheckAsyncResp struct {
		ToUserName   string                  `json:"ToUserName"`   //小程序的username
		FromUserName string                  `json:"FromUserName"` //平台推送服务UserName
		CreateTime   int64                   `json:"CreateTime"`   //发送时间
		MsgType      string                  `json:"MsgType"`      //默认为：event
		Event        string                  `json:"Event"`        //默认为：wxa_media_check
		Appid        string                  `json:"appid"`        //小程序的appid
		Version      string                  `json:"version"`      //可用于区分接口版本
		Detail       []MsgSecCheckDetailData `json:"detail"`       //检测结果详情
		TraceId      string                  `json:"trace_id"`     //任务ids
		Result       MsgSecCheckResultData   `json:"result"`       //综合结果
	}
)

/**
 * 文本内容安全识别
 * @param accessToken 接口调用凭证
 * @param reqParam 请求参数
 * @return *MsgSecCheckResp
 * @return error
 */
func MsgSecCheck(accessToken string, reqParam *MsgSecCheckReq) (*MsgSecCheckResp, error) {
	reqFormParam := map[string]string{"access_token": accessToken}
	reqData, _ := json.Marshal(reqParam)
	apiUrl, err := netutil.EncodeURL(msgSecCheckUrl, reqFormParam)
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

	res := new(MsgSecCheckResp)
	if err := json.Unmarshal(resp.Body(), res); err != nil {
		return nil, err
	}
	if res.ErrCode != 0 {
		return nil, errors.New(res.ErrMsg)
	}

	return res, nil
}

/**
 * 音视频内容安全识别
 * @param accessToken 接口调用凭证
 * @param reqParam 请求参数
 * @return *MediaSecCheckResp
 * @return error
 */
func MediaSecCheck(accessToken string, reqParam *MediaSecCheckReq) (*MediaSecCheckResp, error) {
	reqFormParam := map[string]string{"access_token": accessToken}
	reqData, _ := json.Marshal(reqParam)
	apiUrl, err := netutil.EncodeURL(mediaSecCheckUrl, reqFormParam)
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

	res := new(MediaSecCheckResp)
	if err := json.Unmarshal(resp.Body(), res); err != nil {
		return nil, err
	}
	if res.ErrCode != 0 {
		return nil, errors.New(res.ErrMsg)
	}

	return res, nil
}

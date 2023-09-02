/**
 * @author:       wangxuebing
 * @fileName:     phonenumber.go
 * @date:         2023/5/19 15:40
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
	phoneNumberUrl = "https://api.weixin.qq.com/wxa/business/getuserphonenumber"
)

type (
	WatermarkData struct {
		Timestamp int64  `json:"timestamp"` //时间戳
		AppId     string `json:"appid"`     //小程序APPID
	}
	PhoneNumberInfoData struct {
		PhoneNumber     string        `json:"phonenumber"`     //用户绑定的手机号（国外手机号会有区号）
		PurePhoneNumber string        `json:"purePhoneNumber"` //没有区号的手机号
		CountryCode     int           `json:"countrycode"`     //区号
		Watermark       WatermarkData `json:"watermark"`       //数据水印
	}
	phoneNumberReq struct {
		Code string `json:"code"` //手机号获取凭证
	}
	PhoneNumberResp struct {
		common.ComError
		PhoneInfo PhoneNumberInfoData `json:"phone_info"` //用户手机号信息
	}
)

/**
 * 获取用户手机号
 * @param accessToken 接口调用凭证
 * @param code 手机号获取凭证
 * @return *PhoneNumberResp
 * @return error
 */
func GetPhoneNumber(accessToken, code string) (*PhoneNumberResp, error) {
	reqParam := phoneNumberReq{Code: code}
	reqFormParam := map[string]string{"access_token": accessToken}

	reqData, _ := json.Marshal(reqParam)
	apiUrl, err := netutil.EncodeURL(phoneNumberUrl, reqFormParam)
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

	res := new(PhoneNumberResp)
	if err := json.Unmarshal(resp.Body(), res); err != nil {
		return nil, err
	}
	if res.ErrCode != 0 {
		return nil, errors.New(res.ErrMsg)
	}

	return res, nil
}

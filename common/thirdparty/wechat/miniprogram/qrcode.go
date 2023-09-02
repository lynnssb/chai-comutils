/**
 * @author:       wangxuebing
 * @fileName:     qrcode.go
 * @date:         2023/5/19 15:51
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
	miniQrCodeUrl          = "https://api.weixin.qq.com/wxa/getwxacode"
	miniUnlimitedQRCodeUrl = "https://api.weixin.qq.com/wxa/getwxacodeunlimit"
	qrCodeUrl              = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode"
)

type (
	LineColorData struct {
		R int `json:"r"`
		G int `json:"g"`
		B int `json:"b"`
	}
	// MiniQrCodeReq 获取小程序码
	MiniQrCodeReq struct {
		Path       string        `json:"path"`                  //扫码进入的小程序页面路径，最大长度 1024 字节，不能为空；对于小游戏，可以只传入 query 部分
		With       int           `json:"with,omitempty"`        //二维码的宽度，单位 px。默认值为430，最小 280px，最大 1280px
		AutoColor  bool          `json:"auto_color,omitempty"`  //默认值false；自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
		LineColor  LineColorData `json:"line_color,omitempty"`  //默认值{"r":0,"g":0,"b":0} ；auto_color 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"} 十进制表示
		IsHyaline  bool          `json:"is_hyaline,omitempty"`  //默认值false；是否需要透明底色，为 true 时，生成透明底色的小程序码
		EnvVersion string        `json:"env_version,omitempty"` //要打开的小程序版本。正式版为 "release"，体验版为 "trial"，开发版为 "develop"。默认是正式版
	}

	// MiniUnlimitedQRCodeReq 获取不限制的小程序码
	MiniUnlimitedQRCodeReq struct {
		Screen     string        `json:"screen"`                //最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式）
		Page       string        `json:"page,omitempty"`        //默认是主页，页面 page，例如 pages/index/index，根路径前不要填加 /，不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面。
		CheckPath  bool          `json:"check_path,omitempty"`  //默认是true，检查page 是否存在，为 true 时 page 必须是已经发布的小程序存在的页面（否则报错）；为 false 时允许小程序未发布或者 page 不存在， 但page 有数量上限（60000个）请勿滥用。
		EnvVersion string        `json:"env_version,omitempty"` //要打开的小程序版本。正式版为 "release"，体验版为 "trial"，开发版为 "develop"。默认是正式版。
		Width      int           `json:"width,omitempty"`       //默认430，二维码的宽度，单位 px，最小 280px，最大 1280px
		AutoColor  bool          `json:"auto_color,omitempty"`  //自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调，默认 false
		LineColor  LineColorData `json:"line_color,omitempty"`  //默认是{"r":0,"g":0,"b":0} 。auto_color 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"} 十进制表示
		IsHyaline  bool          `json:"is_hyaline,omitempty"`  //默认是false，是否需要透明底色，为 true 时，生成透明底色的小程序
	}

	// QrCodeReq 获取小程序二维码
	QrCodeReq struct {
		Path  string `json:"path"`            //扫码进入的小程序页面路径，最大长度 128 字节，不能为空
		Width int    `json:"width,omitempty"` //二维码的宽度，单位 px。最小 280px，最大 1280px;默认是430
	}

	QrCodeResp struct {
		common.ComError
		Buffer []byte `json:"buffer"`
	}
)

/**
 * 获取小程序码
 * @param accessToken
 * @param reqParam
 * @return []byte
 * @return error
 */
func GetMiniQrCode(accessToken string, reqParam *MiniQrCodeReq) ([]byte, error) {
	reqFormParam := map[string]string{"access_token": accessToken}
	reqData, _ := json.Marshal(reqParam)
	apiUrl, err := netutil.EncodeURL(miniQrCodeUrl, reqFormParam)
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

	res := new(QrCodeResp)
	if err := json.Unmarshal(resp.Body(), res); err != nil {
		return nil, err
	}
	if res.ErrCode != 0 {
		return nil, errors.New(res.ErrMsg)
	}

	return res.Buffer, nil
}

/**
 * 获取不限制的小程序码
 * @param accessToken
 * @param reqParam
 * @return []byte
 * @return error
 */
func GetMiniUnlimitedQRCode(accessToken string, reqParam *MiniUnlimitedQRCodeReq) ([]byte, error) {
	reqFormParam := map[string]string{"access_token": accessToken}
	reqData, _ := json.Marshal(reqParam)
	apiUrl, err := netutil.EncodeURL(miniUnlimitedQRCodeUrl, reqFormParam)
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

	res := new(QrCodeResp)
	if err := json.Unmarshal(resp.Body(), res); err != nil {
		return nil, err
	}
	if res.ErrCode != 0 {
		return nil, errors.New(res.ErrMsg)
	}

	return res.Buffer, nil
}

/**
 * 获取小程序二维码
 * @param accessToken
 * @param reqParam
 * @return []byte
 * @return error
 */
func QrCode(accessToken string, reqParam *QrCodeReq) ([]byte, error) {
	reqFormParam := map[string]string{"access_token": accessToken}
	reqData, _ := json.Marshal(reqParam)
	apiUrl, err := netutil.EncodeURL(qrCodeUrl, reqFormParam)
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

	res := new(QrCodeResp)
	if err := json.Unmarshal(resp.Body(), res); err != nil {
		return nil, err
	}
	if res.ErrCode != 0 {
		return nil, errors.New(res.ErrMsg)
	}

	return res.Buffer, nil
}

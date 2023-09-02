/**
 * @author:       wangxuebing
 * @fileName:     captchautil.go
 * @date:         2023/2/23 18:07
 * @description:
 */

package captchautil

import "github.com/mojocn/base64Captcha"

var store = base64Captcha.DefaultMemStore

type (
	CaptchaTypeEnum string //枚举
	// ConfigParam 参数
	ConfigParam struct {
		CaptchaType   CaptchaTypeEnum
		DriverAudio   *base64Captcha.DriverAudio
		DriverString  *base64Captcha.DriverString
		DriverChinese *base64Captcha.DriverChinese
		DriverMath    *base64Captcha.DriverMath
		DriverDigit   *base64Captcha.DriverDigit
	}
	CaptchaResp struct {
		CaptchaId  string
		Base64Data string
	}
)

const (
	Audio   CaptchaTypeEnum = "Audio"
	String  CaptchaTypeEnum = "String"
	Math    CaptchaTypeEnum = "Math"
	Chinese CaptchaTypeEnum = "Chinese"
	None    CaptchaTypeEnum = "None"
)

// GenerateCaptcha 生成图形验证码
func GenerateCaptcha(param ConfigParam) (*CaptchaResp, error) {
	var driver base64Captcha.Driver
	switch param.CaptchaType {
	case Audio:
		driver = param.DriverAudio
	case String:
		driver = param.DriverString
	case Math:
		driver = param.DriverMath
	case Chinese:
		driver = param.DriverChinese
	default:
		driver = base64Captcha.DefaultDriverDigit
	}
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		return nil, err
	}
	return &CaptchaResp{CaptchaId: id, Base64Data: b64s}, nil
}

// VerifyCaptcha 验证图形验证码
func VerifyCaptcha(id, answer string, clear bool) bool {
	if store.Verify(id, answer, clear) {
		return true
	}
	return false
}

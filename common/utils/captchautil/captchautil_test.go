/**
 * @author:       wangxuebing
 * @fileName:     captchautil_test.go
 * @date:         2023/3/29 18:09
 * @description:
 */

package captchautil

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
	"log"
	"testing"
)

func TestGenerateCaptcha(t *testing.T) {
	var (
		ic  *CaptchaResp
		err error
	)
	param := ConfigParam{
		CaptchaType: String,
		DriverAudio: nil,
		DriverString: &base64Captcha.DriverString{
			Height:          40,
			Width:           125,
			NoiseCount:      0,
			ShowLineOptions: 2 | 2,
			Length:          5,
			Source:          "0123456789",
			BgColor: &color.RGBA{
				R: 139,
				G: 105,
				B: 20,
				A: 80,
			},
			Fonts: []string{"Comismsh.ttf"},
		},
	}
	ic, err = GenerateCaptcha(param)
	if err != nil {
		log.Println("图形验证码获取失败")
	}
	log.Println(ic.CaptchaId)
	log.Println(ic.Base64Data)
}

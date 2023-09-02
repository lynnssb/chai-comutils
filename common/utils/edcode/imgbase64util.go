/**
 * @author:       wangxuebing
 * @fileName:     imgbase64util.go
 * @date:         2023/3/29 18:31
 * @description:
 */

package edcode

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
)

// ImgToBase64 图片转Base64
func ImgToBase64(img image.Image) (string, error) {
	// 将图像编码为PNG格式的字节切片
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		return "", err
	}

	// 将字节切片编码为base64字符串
	encodedStr := base64.StdEncoding.EncodeToString(buffer.Bytes())
	return encodedStr, nil
}

// Base64ToImg Base64转图片
func Base64ToImg(base64Str string) (image.Image, error) {
	// 将base64字符串解码为字节切片
	decoded, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}

	// 将字节切片解码为图像
	img, err := png.Decode(bytes.NewReader(decoded))
	if err != nil {
		return nil, err
	}

	return img, nil
}

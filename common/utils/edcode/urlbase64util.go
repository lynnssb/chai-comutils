/**
 * @author:       wangxuebing
 * @fileName:     urlbase64util.go
 * @date:         2023/3/29 18:31
 * @description:
 */

package edcode

import "encoding/base64"

// Base64URLEncode Base64URL编码
func Base64URLEncode(data string) string {
	result := base64.URLEncoding.EncodeToString([]byte(data))
	return result
}

// Base64URLDecode Base64URL解码
func Base64URLDecode(data string) string {
	result, _ := base64.URLEncoding.DecodeString(data)
	return string(result)
}

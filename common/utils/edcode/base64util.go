/**
 * @author:       wangxuebing
 * @fileName:     base64util.go
 * @date:         2023/1/14 16:57
 * @description:
 */

package edcode

import (
	"encoding/base64"
)

// Base64Encode Base64编码
func Base64Encode(data string) string {
	result := base64.StdEncoding.EncodeToString([]byte(data))
	return result
}

// Base64Decode Base64解码
func Base64Decode(data string) string {
	result, _ := base64.StdEncoding.DecodeString(data)
	return string(result)
}

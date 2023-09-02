/**
 * @author:       wangxuebing
 * @fileName:     urlencodeutil.go
 * @date:         2023/3/29 18:32
 * @description:
 */

package netutil

import (
	"errors"
	"net/url"
)

// EncodeURL URL编码
func EncodeURL(api string, params map[string]string) (string, error) {
	reqUrl, err := url.Parse(api)
	if err != nil {
		return "", errors.New("url parse error")
	}
	query := reqUrl.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	reqUrl.RawQuery = query.Encode()
	result, _ := url.QueryUnescape(reqUrl.String())
	return result, nil
}

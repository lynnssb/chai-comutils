/**
 * @author:       wangxuebing
 * @fileName:     netutil_test.go
 * @date:         2023/4/10 23:46
 * @description:
 */

package netutil

import (
	"log"
	"testing"
)

func TestEncodeURL(t *testing.T) {
	url := "http://www.baidu.com"

	params := map[string]string{
		"username": "wangxuebing",
		"password": "123456",
	}

	result, err := EncodeURL(url, params)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(result)
	// http://www.baidu.com?password=123456&username=wangxuebing
}

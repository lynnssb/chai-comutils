/**
 * @author:       wangxuebing
 * @fileName:     token_test.go
 * @date:         2023/5/19 14:43
 * @description:
 */

package miniprogram

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	appid := "wx08f5f37dfcfdac3e"
	secret := "ac3df89151f54d7ca612e0fc453e48b7"
	resp, err := GetAccessToken(appid, secret)
	if err != nil {
		log.Fatal(err)
	}
	result, _ := json.Marshal(resp)
	log.Fatal(string(result))
}

func TestGetStableAccessToken(t *testing.T) {
	appid := "wx08f5f37dfcfdac3e"
	secret := "ac3df89151f54d7ca612e0fc453e48b7"
	resp, err := GetStableAccessToken(appid, secret, true)
	if err != nil {
		log.Fatal(err)
	}
	result, _ := json.Marshal(resp)
	log.Fatal(string(result))
}

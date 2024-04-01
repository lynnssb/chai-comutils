package netutil

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

// FastGET 发送GET请求
func FastGET(url string, headerArgs map[string]interface{}, queryArgs map[string]interface{}) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	req.SetRequestURI(url)
	req.Header.SetMethod("GET")

	//设置请求头
	for key, value := range headerArgs {
		req.Header.Set(key, fmt.Sprint(value))
	}

	//设置查询参数
	query := req.URI().QueryArgs()
	for key, value := range queryArgs {
		query.Add(key, fmt.Sprint(value))
	}

	//发送请求
	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() == fasthttp.StatusFound {
		location := resp.Header.Peek("Location")
		if len(location) > 0 {
			req.SetRequestURIBytes(location)
			if err := fasthttp.Do(req, resp); err != nil {
				return nil, err
			}
		}
	}
	//返回响应内容
	return resp.Body(), nil
}

func FastPOST(url string, headerArgs map[string]interface{}, queryArgs map[string]interface{}, bodyArgs []byte) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	//设置请求方法和URL
	req.SetRequestURI(url)
	req.Header.SetMethod("POST")

	//设置查询参数
	if queryArgs != nil {
		for key, value := range queryArgs {
			req.URI().QueryArgs().Add(key, fmt.Sprint(value))
		}
	}

	//设置请求头
	for key, value := range headerArgs {
		req.Header.Set(key, fmt.Sprint(value))
	}

	//设置请求体
	if bodyArgs != nil {
		req.SetBody(bodyArgs)
	}

	//发送请求
	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}

	// 检查状态码，如果是302，再次发送请求
	if resp.StatusCode() == fasthttp.StatusFound {
		location := resp.Header.Peek("Location")
		if len(location) > 0 {
			req.SetRequestURIBytes(location)
			if err := fasthttp.Do(req, resp); err != nil {
				return nil, err
			}
		}
	}

	//返回响应内容
	return resp.Body(), nil
}

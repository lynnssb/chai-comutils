/**
 * @author:       wangxuebing
 * @fileName:     ipaddress.go
 * @date:         2023/1/16 23:09
 * @description:
 */

package netutil

import (
	"net"
	"net/http"
)

const (
	xForwardedFor = "X-Forwarded-For"
	xRealIP       = "X-Real-IP"
)

// RemoteIp 获取客户端IP
func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(xRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(xForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

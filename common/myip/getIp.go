package myip

import (
	"net/http"
	"strings"
)

func GetCurrentIP(r *http.Request) string {
	// 这里也可以通过X-Forwarded-For请求头的第一个值作为用户的ip
	// 但是要注意的是这两个请求头代表的ip都有可能是伪造的
	ip := r.Header.Get("X-Forwarded-For")
	ip = strings.TrimSpace(strings.Split(ip, ",")[0])
	if ip == "" {
		ip = r.Header.Get("X-Real-Ip")
		if ip == "" {
			// 当请求头不存在即不存在代理时直接获取ip
			ip = strings.TrimSpace(strings.Split(r.RemoteAddr, ":")[0])
		}
	}
	return ip
}

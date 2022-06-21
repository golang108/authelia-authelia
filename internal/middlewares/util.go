package middlewares

import (
	"net"
	"strings"

	"github.com/valyala/fasthttp"
)

// GetRemoteIP returns the remote IP for a connection taking into consideration the X-Forwarded-For header.
func GetRemoteIP(ctx *fasthttp.RequestCtx) (ip net.IP) {
	if hdr := ctx.Request.Header.PeekBytes(headerXForwardedFor); hdr != nil {
		ips := strings.Split(string(hdr), ",")

		if len(ips) > 0 {
			ip = net.ParseIP(strings.Trim(ips[0], " "))
			if ip != nil {
				return ip
			}
		}
	}

	return ctx.RemoteIP()
}

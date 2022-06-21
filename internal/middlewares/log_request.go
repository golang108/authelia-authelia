package middlewares

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"github.com/authelia/authelia/v4/internal/logging"
)

// LogRequest provides trace logging for all requests.
func LogRequest(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		autheliaCtx := &AutheliaCtx{RequestCtx: ctx}
		logger := NewRequestLogger(autheliaCtx)

		logger.Trace("Request hit")
		next(ctx)
		logger.Tracef("Replied (status=%d)", ctx.Response.StatusCode())
	}
}

func LogRequestDetailed(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ip := GetRemoteIP(ctx)

		logging.Logger().WithFields(logrus.Fields{
			"ip":     ip.String(),
			"method": string(ctx.Method()),
			"path":   string(ctx.Path()),
			"status": strconv.Itoa(ctx.Response.StatusCode()),
			"uri":    string(ctx.RequestURI()),
			"body":   string(ctx.PostBody()),
		}).Debug("Request Hit")

		next(ctx)

		location := ctx.Response.Header.Peek("Location")

		if location == nil {
			return
		}

		fmt.Printf("%s\n", ctx.Response.Header.Header())

		logging.Logger().WithFields(logrus.Fields{
			"ip":       ip.String(),
			"method":   string(ctx.Method()),
			"path":     string(ctx.Path()),
			"status":   strconv.Itoa(ctx.Response.StatusCode()),
			"uri":      string(ctx.RequestURI()),
			"location": string(location),
		}).Debug("Response Written")

	}
}

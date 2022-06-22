package middlewares

import (
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

func LogRequestHeaders(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		fields := logrus.Fields{}

		ctx.Request.Header.VisitAll(func(key, value []byte) {
			fields[string(key)] = string(value)
		})

		logging.Logger().WithFields(fields).Debug("Verify Request")

		next(ctx)

		fields = logrus.Fields{}

		ctx.Response.Header.VisitAll(func(key, value []byte) {
			fields[string(key)] = string(value)
		})

		logging.Logger().WithFields(fields).Debug("Verify Response")
	}
}

package middleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
)

func LogMiddlewareBuilder(ctx context.Context) endpoint.Middleware {
	logger := ctx.Value(endpoint.CtxLoggerKey).(klog.FormatLogger) // get the logger

	return func(next endpoint.Endpoint) endpoint.Endpoint { // middleware
		return func(ctx context.Context, request, response interface{}) error {
			logger.Debugf("request is: %+v", request)
			err := next(ctx, request, response)
			return err
		}
	}
}

package log

import (
	"context"

	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
)

var _ kitlog.Logger = (*ctxLogger)(nil)

// DefaultMappingCtxKeys 默认上下文日志键值对
// 可以在业务代码中，进行覆写
var DefaultMappingCtxKeys = map[string]any{
	"x_request_id": kithttp.ContextKeyRequestXRequestID,
	"method":       kithttp.ContextKeyRequestMethod,
}

type ctxLogger struct {
	logger kitlog.Logger
	ctx    context.Context
}

func newContextLogger(logger kitlog.Logger) *ctxLogger {
	if c, ok := logger.(*ctxLogger); ok {
		return c
	}
	return &ctxLogger{logger: logger}
}

func WithCtx(ctx context.Context, logger kitlog.Logger) kitlog.Logger {

	l := newContextLogger(logger)

	return &ctxLogger{
		logger: l.logger,
		ctx:    ctx,
	}
}

func (l *ctxLogger) Log(kvs ...interface{}) error {

	for logKey, ctxKey := range DefaultMappingCtxKeys {
		if v, ok := l.ctx.Value(ctxKey).(string); ok {
			kvs = append(kvs, logKey, v)
		}
	}

	return l.logger.Log(kvs...)
}

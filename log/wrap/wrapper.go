// Package wrap 日志包装器，可用于打印上下文信息 (链路ID...)
package wrap

import (
	"context"

	kithttp "github.com/go-kit/kit/transport/http"

	kitlog "github.com/go-kit/log"
)

// DefaultContextKVs 默认上下文键值对
var DefaultContextKVs = map[string]any{
	"method":     kithttp.ContextKeyRequestMethod,
	"request_id": kithttp.ContextKeyRequestXRequestID,
}

// log 上下文日志打印
func log(ctx context.Context, logger kitlog.Logger, level string) kitlog.Logger {
	keyvals := []any{"level", level}
	for loggerKey, ctxKey := range DefaultContextKVs {
		if ctxVal, ok := ctx.Value(ctxKey).(string); ok {
			keyvals = append(keyvals, loggerKey, ctxVal)
		}
	}
	return kitlog.With(logger, keyvals...)
}

func Error(ctx context.Context, logger kitlog.Logger) kitlog.Logger {
	return log(ctx, logger, "error")
}

func Warn(ctx context.Context, logger kitlog.Logger) kitlog.Logger {
	return log(ctx, logger, "warn")
}

func Info(ctx context.Context, logger kitlog.Logger) kitlog.Logger {
	return log(ctx, logger, "info")
}

func Debug(ctx context.Context, logger kitlog.Logger) kitlog.Logger {
	return log(ctx, logger, "debug")
}

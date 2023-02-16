package grpc

import (
	"context"
	"errors"
	"fmt"
	"runtime"

	kitlog "github.com/go-kit/kit/log"
	"google.golang.org/grpc"
)

// ErrorHandleInterceptor GRPC 异常处理拦截器
func ErrorHandleInterceptor(logger kitlog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("unknown error")
				}

				// 错误日志
				if logger != nil {
					logger.Log("errors", printStackTrace(err))
				}
			}
		}()
		return handler(ctx, req)
	}
}

// 打印堆栈信息
func printStackTrace(err interface{}) []string {
	lines := make([]string, 0)
	lines = append(lines, fmt.Sprintf("%v", err))
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		lines = append(lines, fmt.Sprintf("%s:%d (0x%x)", file, line, pc))
	}
	return lines
}

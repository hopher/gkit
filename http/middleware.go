package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime"

	kitlog "github.com/go-kit/log"
)

type Middleware func(http.Handler) http.Handler

// CORS 跨域中间件
func CORS(headers map[string]string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if headers == nil {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
			} else {
				for k, v := range headers {
					w.Header().Set(k, v)
				}
			}

			if r.Method == "OPTIONS" {
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// Recover 异常处理
func Recover(logger kitlog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				rec := recover()
				if rec != nil {
					var err error
					switch t := rec.(type) {
					case string:
						err = errors.New(t)
					case error:
						err = t
					default:
						err = errors.New("unknown error")
					}

					// 记录错误日志
					if logger != nil {
						logger.Log("errors", printStackTrace(err))
					}

					res := map[string]interface{}{
						"code": 1,
						"msg":  fmt.Sprintf("服务器内部错误: %s", err.Error()),
					}
					bytes, _ := json.Marshal(res)
					w.Header().Set("Content-Type", "application/json; charset=utf-8")
					w.Write(bytes)
				}
			}()

			next.ServeHTTP(w, r)
		})
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

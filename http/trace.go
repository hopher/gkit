package http

import (
	"net/http"

	"github.com/google/uuid"
)

// Trace 链路ID，每个请求中增加 requestID，方便日志查询，链路追踪
func Trace(secret string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			requestID := r.Header.Get("X-Request-Id")

			// 验证密钥，防止 X-Request-Id 伪造
			// 密钥正确时，则采用上游 X-Request-Id
			if secret == "" || r.Header.Get("X-Request-Secret") != secret || requestID == "" {
				requestID = uuid.NewString()
			}

			r.Header.Set("X-Request-Id", requestID)
			w.Header().Set("X-Request-ID", requestID)

			next.ServeHTTP(w, r)
		})
	}
}

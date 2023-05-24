package http

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
)

// Trace 链路ID，每个请求中增加 requestID，方便日志查询，链路追踪
func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			requestID := r.Header.Get("X-Request-Id")

			if requestID == "" {
				requestID = uuid.NewString()
				requestID = strings.Replace(requestID, "-", "", -1)
			}

			r.Header.Set("X-Request-Id", requestID)
			w.Header().Set("X-Request-Id", requestID)

			next.ServeHTTP(w, r)
		})
	}
}

package http

import (
	"fmt"
	"net/http"
	"time"
)

func Example_basic() {

	errChan := make(chan error)

	r := NewRouter()

	// HTTP 服务
	go (func() {

		var handler http.Handler
		{
			handler = CORS(nil)(r)          // 支持跨域
			handler = Recover(nil)(handler) // 错误处理
		}

		hs := &http.Server{
			Handler:      handler,
			Addr:         fmt.Sprintf(":%d", 8080),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
		errChan <- hs.ListenAndServe()
	})()

}

func ExampleTrace() {
	var handler http.Handler
	{
		handler = CORS(nil)(handler)           // 支持跨域
		handler = Recover(nil)(handler)        // 错误处理
		handler = Trace("secret KEY")(handler) // 链路ID
	}
}

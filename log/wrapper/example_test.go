package wrapper

import (
	"context"
	"os"

	kithttp "github.com/go-kit/kit/transport/http"

	kitlog "github.com/go-kit/log"
)

func Example_basic() {

	var logger kitlog.Logger
	{
		logger = kitlog.NewJSONLogger(os.Stdout)
		//logger = kitlog.NewLogfmtLogger(os.Stdout)
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, kithttp.ContextKeyRequestXRequestID, "1111")
	ctx = context.WithValue(ctx, kithttp.ContextKeyRequestMethod, "POST")

	Error(ctx, logger).Log("aaa", "bbb")
	Warn(ctx, logger).Log("aaa", "bbb")
	Info(ctx, logger).Log("aaa", "bbb")
	Debug(ctx, logger).Log("aaa", "bbb")
	// Output:
	// {"aaa":"bbb","level":"error","method":"POST","request_id":"1111"}
	// {"aaa":"bbb","level":"warn","method":"POST","request_id":"1111"}
	// {"aaa":"bbb","level":"info","method":"POST","request_id":"1111"}
	// {"aaa":"bbb","level":"debug","method":"POST","request_id":"1111"}

}

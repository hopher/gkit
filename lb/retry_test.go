package lb

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-kit/kit/sd/lb"
)

func TestRetryMaxTotalFail(t *testing.T) {
	var (
		e = func(context.Context, interface{}) (interface{}, error) {
			t.Log("run")
			time.Sleep(1 * time.Second)
			return nil, errors.New("error one")
		}

		retry = lb.Retry(999, 10*time.Second, BalancerFunc(e)) // 999次重试，10秒钟超时
		ctx   = context.Background()
	)

	// 999次重试中，每一次都应该失败
	if _, err := retry(ctx, struct{}{}); err == nil {
		t.Errorf("expected error, got none") // should fail
	}
}

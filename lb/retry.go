// Package lb 负载平衡器
package lb

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd/lb"
)

var _ lb.Balancer = (BalancerFunc)(nil)

type BalancerFunc endpoint.Endpoint

func (f BalancerFunc) Endpoint() (endpoint.Endpoint, error) {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return f(ctx, request)
	}, nil
}

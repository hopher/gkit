package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Options = redis.Options

func NewClient(ctx context.Context, o *redis.Options) (*redis.Client, error) {
	client := redis.NewClient(o)

	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, fmt.Errorf("redis ping err: %w", err)
	}
	return client, nil
}

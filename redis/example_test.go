package redis

import (
	"context"
)

func ExampleNewClient() {

	var ctx = context.Background()

	rdb, err := NewClient(ctx, &Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		DB:       0,
	})
	if err != nil {
		panic(err)
	}

	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
}

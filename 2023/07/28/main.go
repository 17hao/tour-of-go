package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
	})

	ctx := context.Background()

	v0, err := rdb.Set(ctx, "k1", "v1", 0*time.Second).Result()
	if err != nil {
		logrus.Errorf("%v", err)
	}
	fmt.Printf("%s\n", v0)

	v1, err := rdb.Get(ctx, "k1").Result()
	if err != nil {
		logrus.Errorf("%v", err)
	}
	fmt.Printf("%s\n", v1)

	v2, err := rdb.ZAdd(ctx, "zset1", redis.Z{Member: "z1", Score: 1}, redis.Z{Member: "z2", Score: 2}, redis.Z{Member: "z3", Score: 3}).Result()
	if err != nil {
		logrus.Errorf("%v", err)
	}
	fmt.Printf("%v\n", v2)

	v3, err := rdb.ZRem(ctx, "zset1", "z2").Result()
	if err != nil {
		logrus.Errorf("%v", err)
	}
	fmt.Printf("%v\n", v3)

	v4, err := rdb.ZRangeByScore(ctx, "zset1", &redis.ZRangeBy{Min: "-inf", Max: "+inf"}).Result()
	if err != nil {
		logrus.Errorf("%v", err)
	}
	fmt.Printf("%v\n", v4)

	v5, err := rdb.ZRange(ctx, "zset1", 0, -1).Result()
	if err != nil {
		logrus.Errorf("%v", err)
	}
	fmt.Printf("%v\n", v5)
}

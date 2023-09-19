package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
	})
	strs := []string{"v1", "v2"}
	if err := rdb.RPush(context.Background(), "list-k1", strs).Err(); err != nil {
		fmt.Printf("%+v", err)
		return
	}
	res, err := rdb.LRange(context.Background(), "list-k1", 0, -1).Result()
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	fmt.Printf("res=%+v", res)
}

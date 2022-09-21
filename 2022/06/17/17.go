package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "k1", "v1")
	ctx = context.WithValue(ctx, "k2", "v2")
	fmt.Println(ctx.Value("k1"))
	fmt.Println(ctx.Value("k2"))
}

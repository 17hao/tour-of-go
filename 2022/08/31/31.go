package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "k1", "v1")
	ctx = context.WithValue(ctx, "k2", 2)
	ctx = context.WithValue(ctx, "k3", "3")

	fmt.Printf("ctx: %v", ctx)
}

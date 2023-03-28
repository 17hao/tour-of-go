package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func main() {
	ctxDone()
}

func ctxValue() {
	start := time.Now()
	ctx := context.Background()
	for i := 0; i < 1000; i++ {
		key := "k" + strconv.Itoa(i)
		value := "v" + strconv.Itoa(i)
		ctx = context.WithValue(ctx, key, value)
	}

	// f1
	for i := 0; i < 100; i++ {
		f1(ctx)
	}
	now := time.Now()
	duration := now.Sub(start)
	fmt.Printf("[f1] exec time: %+v\n", duration)

	// f2
	for i := 0; i < 100; i++ {
		f2(ctx)
	}
	now = time.Now()
	duration = now.Sub(start)
	fmt.Printf("[f2] exec time: %+v\n", duration)
}

func f1(ctx context.Context) {
	fmt.Print(ctx.Value("k0"), "\n")
}

func f2(ctx context.Context) {
	fmt.Print(ctx.Value("k999"), "\n")
}

func ctxDone() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	f3(ctx)
	defer cancel()
}

func f3(ctx context.Context) {
	select {
	case <- ctx.Done():
		fmt.Println("context done")
	default:
		fmt.Println("default")
	}
}

package main

import (
	"context"
	"fmt"
	"net/http"
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
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

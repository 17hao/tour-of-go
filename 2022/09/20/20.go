package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//printBool(true, false)

	mirrors := make(chan string, 3)

	go func() {
		rand.Seed(time.Now().Unix() + 1)
		n := rand.Intn(3)
		fmt.Printf("a:%d\n", n)
		time.Sleep(time.Duration(n) * time.Second)
		mirrors <- "a.shiqihao.xyz"
	}()

	go func() {
		rand.Seed(time.Now().Unix() + 2)
		n := rand.Intn(3)
		fmt.Printf("h:%d\n", n)
		time.Sleep(time.Duration(n) * time.Second)
		mirrors <- "h.shiqihao.xyz"
	}()

	go func() {
		rand.Seed(time.Now().Unix() + 3)
		n := rand.Intn(3)
		fmt.Printf("x:%d\n", n)
		time.Sleep(time.Duration(n) * time.Second)
		mirrors <- "x.shiqihao.xyz"
	}()

	println(<-mirrors)
}

func printBool(b ...bool) {
	fmt.Printf("%t\n", b)
	fmt.Printf("%v", b)
}

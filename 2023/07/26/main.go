package main

import (
	"errors"
	"time"
)

func f2() error {
	return errors.New("f2")
}

func f1() error {
	return f2()
}

func main() {
	now := time.Now()
	println(now.Unix() - now.UnixNano()/1000000000)

	// err := f1()
	// if err != nil {
	// 	fmt.Printf("%s", string(debug.Stack()))
	// }
}

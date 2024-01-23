package main

import (
	"fmt"
	"time"
)

func sendToChannel(c chan int) {
	tick := time.Tick(1 * time.Second)
	timeout := time.After(2 * time.Second)
	for {
		select {
		case t := <-tick:
			fmt.Printf("insert to c: %d\n", t.Unix())
			c <- int(t.Unix())
		case <-timeout:
			close(c)
			fmt.Println("timed out, close c")
			return
		default:
			// no-op
		}
	}
}

func main() {
	c := make(chan int)

	go func() {
		fmt.Println("consume c")
		for {
			select {
			// 注释掉第一个case语句，channel c只有发送端，没有接收端，发送第二条消息开始会被阻塞
			//case n, ok := <-c:
			//	if !ok {
			//		fmt.Println("c is closed")
			//		return;
			//	}
			//	fmt.Printf("receive from c: %d\n", n)
			default:
				// no-op
			}
		}
	}()

	sendToChannel(c)
}

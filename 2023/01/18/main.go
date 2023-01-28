package main

import "fmt"

func main() {
	c := make(chan string, 1)
	//c := make(chan string)
	done := make(chan bool)
	go func() {
		ping(c, "ping")
	}()

	go func() {
		pong(c, done)
	}()

	<-done
}

func ping(c chan<- string, s string) {
	fmt.Print("ping\n")
	c <- s

}

func pong(c <-chan string, done chan<- bool) {
	//select {
	//case <-c:
	//	fmt.Printf("pong\n")
	//}
	<-c
	fmt.Println("pong")
	done <- true
}

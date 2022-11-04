package main

import "time"

func main() {
	println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		println(countdown)
		<-tick
	}
}

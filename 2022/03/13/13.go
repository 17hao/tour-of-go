package main

import "fmt"

func main() {
	c := Configuration{
		value: "test",
		proxy: &Proxy{
			address: "localhost",
			port:    8888,
		},
	}
	fmt.Println(c.proxy.address)
}

type Configuration struct {
	value string
	proxy *Proxy
}

type Proxy struct {
	address string
	port    int
}

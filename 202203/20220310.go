package main

import "fmt"

type Client interface {
	sum(i int, j int) int
}

type WrapperClient struct {
	c *Client
}

type ServiceClient struct {
	w *WrapperClient
}

func (c ServiceClient) sum(i int, j int) int {
	return i + j
}

func main() {
	client := &ServiceClient{}
	fmt.Println(client.sum(1, 2))
}

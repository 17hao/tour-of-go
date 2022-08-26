package main

import (
	"context"
	"log"
	"time"

	"github.com/17hao/echo_server/kitex_gen/api"
	"github.com/17hao/echo_server/kitex_gen/api/echoservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

func main() {
	cli, err := echoservice.NewClient("example", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		panic(err)
	}
	req := &api.EchoRequest{Message: "my message"}
	resp, err := cli.Echo(context.Background(), req, callopt.WithConnectTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}

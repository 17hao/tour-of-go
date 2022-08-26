package main

import (
	"github.com/apache/thrift/lib/go/thrift"
	"shiqihao.xyz/tour-of-go/rpc/thrift/tutorial"
)

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	host := "localhost:9090"
	tutorial.RunServer(transportFactory, protocolFactory, host, false)
}

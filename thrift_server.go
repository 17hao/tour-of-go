package main

import (
	"github.com/apache/thrift/lib/go/thrift"
	"shiqihao.xyz/tour-of-go/thrift_tutorial"
)

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	host := "localhost:9090"
	thrift_tutorial.RunServer(transportFactory, protocolFactory, host, false)
}

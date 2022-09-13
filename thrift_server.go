package main

import (
	"log"

	"github.com/17hao/tour-of-go/rpc/thrift/tutorial/server"
	"github.com/17hao/tour-of-go/thrift-gen/tutorial"
	"github.com/apache/thrift/lib/go/thrift"
)

func main() {
	host := "localhost:9090"
	runServer(host)
}

func runServer(addr string) {
	handler := server.NewCalculatorHandler()
	processor := tutorial.NewCalculatorProcessor(handler)

	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		log.Fatal(err)
	}

	transportFactory := thrift.NewTFramedTransportFactoryConf(thrift.NewTTransportFactory(), &thrift.TConfiguration{})
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(&thrift.TConfiguration{})

	s := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	log.Printf("Starting the tutorial server on %s\n", addr)
	err = s.Serve()
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/17hao/tour-of-go/thrift-gen/tutorial"
	"github.com/apache/thrift/lib/go/thrift"
)

func main() {
	host := "localhost:9090"
	client, err := newClient(host)
	if err != nil {
		panic(err)
	}

	handleClient(client)
}

func newClient(addr string) (*tutorial.CalculatorClient, error) {
	//transportFactory := thrift.NewTFramedTransportFactoryConf(thrift.NewTTransportFactory(), &thrift.TConfiguration{})
	//protocolFactory := thrift.NewTBinaryProtocolFactoryConf(&thrift.TConfiguration{})
	//socket := thrift.NewTSocketConf(addr, &thrift.TConfiguration{})

	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	socket, err := thrift.NewTSocket(addr)
	if err != nil {
		log.Fatal(err)
	}
	transport, err := transportFactory.GetTransport(socket)
	if err != nil {
		return nil, err
	}
	if err = transport.Open(); err != nil {
		return nil, err
	}

	iport := protocolFactory.GetProtocol(transport)
	oport := protocolFactory.GetProtocol(transport)

	return tutorial.NewCalculatorClient(thrift.NewTStandardClient(iport, oport)), nil
}

func handleClient(client *tutorial.CalculatorClient) (err error) {
	ctx := context.Background()
	client.Ping(ctx)
	fmt.Println("ping()")

	sum, _ := client.Add(ctx, 1, 1)
	fmt.Print("1+1=", sum, "\n")

	work := tutorial.NewWork()
	work.Op = tutorial.Operation_DIVIDE
	work.Num1 = 1
	work.Num2 = 0
	quotient, err := client.Calculate(ctx, 1, work)
	if err != nil {
		switch v := err.(type) {
		case *tutorial.InvalidOperation:
			fmt.Println("Invalid operation:", v)
		default:
			fmt.Println("Error during operation:", err)
		}
	} else {
		fmt.Println("Whoa we can divide by 0 with new value:", quotient)
	}

	work.Op = tutorial.Operation_SUBTRACT
	work.Num1 = 15
	work.Num2 = 10
	diff, err := client.Calculate(ctx, 1, work)
	if err != nil {
		switch v := err.(type) {
		case *tutorial.InvalidOperation:
			fmt.Println("Invalid operation:", v)
		default:
			fmt.Println("Error during operation:", err)
		}
		return err
	} else {
		fmt.Print("15-10=", diff, "\n")
	}

	log, err := client.GetStruct(ctx, 1)
	if err != nil {
		fmt.Println("Unable to get struct:", err)
		return err
	} else {
		fmt.Println("Check log:", log.Value)
	}
	return err
}

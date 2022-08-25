package main

import (
	"fmt"
	"log"

	"github.com/bytedance/sonic"
	"shiqihao.xyz/tour-of-go/greetings"
)

func main() {
	//fmt.Printf("hello, world")
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//message, err := greetings.Hello("sqh")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf(message)

	names := []string{"sqh", "lyj"}
	for idx, name := range names {
		s := fmt.Sprint(idx) + ": " + name
		fmt.Println(s)
	}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	var res map[string]string
	sonic.Unmarshal(messages, &res)
	for k, v := range res {
		fmt.Println(k, " ", v)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", info)
}

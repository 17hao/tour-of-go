package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("/Users/bytedance/a.txtt")
	if err != nil {
		log.Fatal(err)
	}
	b, _ := ioutil.ReadAll(f)
	fmt.Println(string(b))
}

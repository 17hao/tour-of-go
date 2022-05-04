package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Foo28() int {
	return 0x7fffffff
}

func main() {
	type User struct {
		Name string `json:"name"`
		Age int `json:"age"`
		Address string `json:"address,omitempty"`
	}

	u := User{}
	dir, _ := os.Getwd()
	file, _ := os.Open(dir + "/conf/test.json")
	content, _ := ioutil.ReadAll(file)
	json.Unmarshal(content, &u)
	fmt.Printf("%+v\n", u)
}
package main

import (
	"encoding/json"
	"fmt"
)

type s1 struct {
	ID       int           `json:"id"`
	Age      int           `json:"age"`
	Accounts []interface{} `json:"accounts"`
}

type s2 struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Accounts []struct {
		Key string `json:"key"`
	} `json:"accounts"`
}

func main() {
	i := s1{ID: 1, Age: 10}
	j := s2{
		ID:   2,
		Name: "aaa",
		Accounts: []struct {
			Key string `json:"key"`
		}([]struct {
			Key string
		}{
			{
				Key: "key_xxx",
			},
		}),
	}

	bs, _ := json.Marshal(i)
	println(string(bs))

	bs, _ = json.Marshal(j)
	println(string(bs))

	x := s1{}
	err := json.Unmarshal(bs, &x)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	fmt.Printf("%+v\n", x)

	y := s1{}
	_ = json.Unmarshal([]byte(""), &y)
	fmt.Printf("%+v\n", y)
}

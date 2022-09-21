package main

import (
	"encoding/json"
	"fmt"
)

type val struct {
	V string `json:"v"`
	T string `json:"t"`
}

func main() {
	v1 := val{V: "v_v", T: "t_t"}
	b, _ := json.Marshal(v1)
	fmt.Println(string(b))

	v2 := val{}
	_ = json.Unmarshal(b, &v2)
	fmt.Printf("%v\n", v2)
	fmt.Printf("%+v\n", v2)
}

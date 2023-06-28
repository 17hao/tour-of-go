package main

import (
	"encoding/json"
	"fmt"
)

type body struct {
	F2 struct {
		F3 struct {
		} `json:"f3"`
	} `json:"f2"`
}

func main() {
	m := map[string]string{"k1": "v1"}
	f1(m)
	fmt.Printf("%+v\n", m)

	m1 := make(map[string]interface{})
	f2(m1)
	fmt.Printf("%+v", m1)

	bytes, _ := json.Marshal(m1)
	fmt.Println(string(bytes))
	var b body
	_ = json.Unmarshal(bytes, &b)
	fmt.Printf("%+v\n", b)
}

func f1(m map[string]string) {
	m["k2"] = "v2"
}

func f2(m map[string]interface{}) {
	m["f2"] = make(map[string]interface{})
	f3(m["f2"].(map[string]interface{}))
}

func f3(m map[string]interface{}) {
	m["f2"] = make(map[string]interface{})
}

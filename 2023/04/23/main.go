package main

import (
	"encoding/json"
	"fmt"
)

type model1 struct {
	Id int64
}

type model2 struct {
	ID int64
}

func main() {
	m1 := model1{Id: 100}
	bytes, _ := json.Marshal(m1)
	fmt.Println(string(bytes))

	var m2 model2
	_ = json.Unmarshal(bytes, &m2)
	fmt.Println(m2)
}

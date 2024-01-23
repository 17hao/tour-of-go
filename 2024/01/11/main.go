package main

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

type request struct {
	Title string
}

func main() {
	//str := "{\"title\":\"aaa\\nbbb\"}"
	str := `{"title":"aaa"bbb"}`
	var req request
	if err := json.Unmarshal([]byte(str), &req); err != nil {
		logrus.Error(err)
		return
	}
	fmt.Printf("%+v", req)
}

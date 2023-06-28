package main

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	str := "{\"code\":0,\"msg\":\"mock msg\",\"data\":{}}"
	resp := struct {
		Code int32    `json:"code"`
		Msg  string   `json:"msg"`
		Data struct{} `json:"data"`
	}{}
	err := json.Unmarshal([]byte(str), &resp)
	if err != nil {
		logrus.Error(err)
	}
	fmt.Printf("%+v", resp)
}

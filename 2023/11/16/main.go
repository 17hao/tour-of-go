package main

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	type request struct {
		Body string `json:"body"`
	}
	type requestBody struct {
		TenantID string `json:"TenantId"`
	}

	// var reqBody = requestBody{TenantID: "123"}
	// b, err := json.Marshal(&reqBody)
	// if err != nil {
	// 	return
	// }
	// var req = request{Body: string(b)}
	// b, err = json.Marshal(&req)
	// fmt.Println(string(b))

	str := `{"body":"{\"IdempotentId\":\"7301152442590560258\",\"MessagePriority\":1,\"TenantId\":\"6968351207937130497\"}"}`

	var req request
	if unmarshalErr := json.Unmarshal([]byte(str), &req); unmarshalErr != nil {
		logrus.Errorf("unmarshal req.Request failed, req.Request=%s, err=%+v", str, unmarshalErr)
		return
	}

	var reqBody requestBody
	if unmarshalErr := json.Unmarshal([]byte(req.Body), &reqBody); unmarshalErr != nil {
		logrus.Errorf("unmarshal request.Body failed, request.Body=%s, err=%+v", req.Body, unmarshalErr)
		return
	}
	fmt.Printf("%s", reqBody.TenantID)
}

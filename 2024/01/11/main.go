package main

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

type request struct {
	Text string `json:"text"`
}

func main() {
	//str := "{\"title\":\"aaa\\nbbb\"}"
	//s = strconv.QuoteToASCII(s)
	//str := fmt.Sprintf(`{"text":"%s"}`, s)

	s := `aaa"bbb@😀`
	r1 := request{}
	r1.Text += s
	b, err := json.Marshal(r1)
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Printf("r=%s\n", string(b))

	var r2 request
	if err := json.Unmarshal(b, &r2); err != nil {
		logrus.Error(err)
		return
	}
	fmt.Printf("%+v", r2)
}

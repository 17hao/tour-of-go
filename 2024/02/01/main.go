package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type s struct {
	builder strings.Builder
}

func main()  {
	s1 := s{}
	s1.builder.WriteString("foo")
	s1.builder.WriteString(`aaa"bbb\n😀`)
	m := map[string]string{"text": s1.builder.String()}
	b, err := json.Marshal(m)
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Print(string(b))
}

package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	str := "/"

	fmt.Printf("%v\n", strings.Split(str, "/"))
	fmt.Printf("%d\n", len(strings.Split(str, "/")))

	regex := regexp.MustCompile(`{{flow-migrate-\w+}}`)
	input := "123{{flow-migrate-0}}{{flow-migrate-number}}"
	//input := "123{{flow-migrate-0}}abc{{flow-migrate-number}}{{flow-migrate-__body__}}abc"
	//input := "{{flow-migrate-0}}abc{{flow-migrate-number}}{{flow-migrate-__body__}}"
	indexes := regex.FindAllIndex([]byte(input), -1)

	for i, idx := range indexes {
		if i == 0 && idx[0] > 0 {
			println(input[0:idx[0]])
		}
		println(input[idx[0]:idx[1]])
		if i == len(indexes)-1 && idx[1] < len(input) {
			println(input[idx[1]:])
		}
	}
	matches := regex.FindAllString(input, -1)
	for _, m := range matches {
		println(m)
	}

	str = `{\"0\":\"123\",\"2\":\"\"}`
	str, _ = strconv.Unquote(str)
	m := map[string]string{}
	_ = json.Unmarshal([]byte(str), &m)
	fmt.Printf("%+v", m)
}

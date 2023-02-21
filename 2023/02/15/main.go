package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(uuid.New().String())

	str := fmt.Sprintf("{\"resp\": \"%s\"}", "data\n")
	fmt.Println(strings.Contains(str, "\n"))
	str = strings.ReplaceAll(str, "\n", "\\n")
	fmt.Printf("%s\n", str)
	fmt.Println(str == "{\"resp\": \"data\\n\"}")
	m := map[string]string{}
	_ = json.Unmarshal([]byte(str), &m)
	fmt.Printf("%v", m)
	//str := `{"resp": "data\n"}`

}

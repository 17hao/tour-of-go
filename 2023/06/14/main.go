package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "queryJobKeys 1,2,3"
	substr := strings.ReplaceAll(str, "queryJobKeys ", "")
	fmt.Printf("%s\n", substr)

	keys := map[string][]string{}
	keys["k1"] = append(keys["k1"], "v1")
	keys["k1"] = append(keys["k1"], "v2")
	fmt.Printf("%+v\n", keys)
}

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	//str := "@$#$#@var1@$#$#@@$#$#@var2@$#$#@"
	str := "@$#$#@var1@$#$#@@$#$#@var2@$#$#@text"
	fmt.Printf("len(str)=%d\n", len(str))
	regex := regexp.MustCompile(`@\$#\$#@`)
	indexes := regex.FindAllIndex([]byte(str), -1)
	fmt.Printf("%+v\n", indexes)
	for _, index := range indexes {
		fmt.Println(str[index[0]:index[1]])
	}

	varIndexes := make([][]int, 0)
	for i := 0; i < len(indexes)-1; i += 2 {
		varIndexes = append(varIndexes, []int{indexes[i][0], indexes[i+1][1]})
	}
	for _, varIndex := range varIndexes {
		s := strings.ReplaceAll(str[varIndex[0]:varIndex[1]], "@$#$#@", "")
		fmt.Println(s)
	}
	fmt.Printf("%+v\n", varIndexes)

	constIndexes := make([][]int, 0)
	if varIndexes[0][0] > 0 {
		constIndexes = append(constIndexes, []int{0, varIndexes[0][0]})
	}
	for i := 0; i < len(varIndexes)-1; i++ {
		if varIndexes[i][1] == varIndexes[i+1][0] { // 连续2个变量
			continue
		}
		constIndexes = append(constIndexes, []int{varIndexes[i][1], varIndexes[i+1][0]})
	}
	if varIndexes[len(varIndexes)-1][1] < len(str) {
		constIndexes = append(constIndexes, []int{varIndexes[len(varIndexes)-1][1], len(str)})
	}
	for _, constIndex := range constIndexes {
		fmt.Println(str[constIndex[0]:constIndex[1]])
	}
	fmt.Printf("%+v\n", constIndexes)
}

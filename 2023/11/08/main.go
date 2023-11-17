package main

import "fmt"

func f1(set map[string]struct{}) {
	if _, ok := set["a"]; ok {
		fmt.Printf("ok")
	} else {
		fmt.Printf("!ok")
	}
}

func f2(mapping map[string]struct{}, nums []string) *string {
	for _, num := range nums {
		if _, ok := mapping[num]; ok {
			return &num
		}
	}
	return nil
}

func main() {
	// f1(nil)
	mapping := map[string]struct{}{"1": {}, "2": {}}
	fmt.Printf("%+v\n", *f2(mapping, []string{"1"}))
	fmt.Printf("%+v\n", f2(mapping, []string{"3"}))
}

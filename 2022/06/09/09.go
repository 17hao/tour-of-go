package main

import "fmt"

type user struct {
	id   int
	name string
}

func main() {
	args := make(map[string]interface{}, 0)
	u1 := user{
		id:   1,
		name: "test",
	}
	args["u1"] = u1
	fmt.Printf("%+v\n", args)
}

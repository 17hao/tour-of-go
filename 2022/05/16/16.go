package main

import "encoding/json"

type user struct {
	Name string
}

func main() {
	m := map[string]string{"key": "value"}
	modifyMap(m)
	b, _ := json.Marshal(m)
	println(string(b))

	u := user{Name: "u1"}
	modifyUser(u)
	println(u.Name)

	x := 1
	foo16(&x)
	println(x)
}

func modifyMap(m map[string]string) {
	m["key"] = "update"
}

func modifyUser(u user) {
	u.Name = "update u"
}

func foo16(x *int) {
	*x = 2
	y := 7
	x = &y
}

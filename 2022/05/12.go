package main

func main() {
	var i interface{}

	list := make([]interface{}, 4)
	list[0] = "0"
	i = list
	switch t := i.(type) {
	case []interface{}:
		println("[]interface{}: ", t)
	case []string:
		println("[]string: ", t)
	default:
		println(t)
	}
}

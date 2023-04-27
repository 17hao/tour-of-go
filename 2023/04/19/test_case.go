package main

// go tool compile -W test_case.go
func f() {
	a := 1
	b := 1
	if true {
		add(a, b)
	}
}

func add(i, j int) {
	println(i + j)
}

package main

type point struct {
	x int
	y int
}

// pass by value
func main() {
	p := point{x: 1, y: 1}
	foo(p)

	bar(p)
	foo(p)
}

func foo(p point) {
	println(p.x)
}

func bar(p point) {
	p.x = 2
}

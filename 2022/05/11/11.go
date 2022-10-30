package main

type Arg struct {
	isEmpty bool
	//IsEmpty func() bool
}

func (a Arg) IsEmpty() bool {
	return a.isEmpty
}

type Func struct {
	Arg
	stmt string
}

func main() {
	f := Func{Arg{isEmpty: false}, ""}
	println(f.IsEmpty())
}

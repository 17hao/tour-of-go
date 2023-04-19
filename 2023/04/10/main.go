package main

type Outer struct {
	In  *Inner
	Val *int
}

func (t *Outer) Empty() int {
	return 1
}

func (t *Outer) GetVal() int {
	return *t.Val
}

type Inner struct {
	Val *int
}

//go:noinline
func (t *Inner) GetVal() int {
	return *t.Val
}

func main() {
	var out *Outer
	println(out.Empty())
	println(out.In.GetVal())
}

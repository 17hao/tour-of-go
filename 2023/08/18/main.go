package main

type ihandler interface {
	f1() string

	f2() int64
}

type defaultF2 struct{}

func (f2 defaultF2) f2() int64 {
	return 1
}

type handler struct {
	defaultF2
}

func (h handler) f1() string {
	return "str"
}

func main() {

}

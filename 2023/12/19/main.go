package main

type iS interface {
	f1() int64
}

type abstractS struct {}

func (*abstractS) f1() int64 {
	return 1
}

type s1 struct {
	*abstractS
}

func (*s1) f1() int64 {
	return 2
}

func (*s1) f2() int64 {
	return 2
}

var s s1

func main() {
	s.abstractS.f1()
}

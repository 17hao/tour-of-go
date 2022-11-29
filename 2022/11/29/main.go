package main

type s struct {
	b bool
}

func (v s) valueReceiverM() {
	v.b = true
}

func (p *s) pointerReceiverM() {
	p.b = true
}

func main() {
	i := s{}
	i.valueReceiverM()
	println(i.b)

	i.pointerReceiverM()
	println(i.b)
}

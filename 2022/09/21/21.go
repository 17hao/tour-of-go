package main

func main() {
	e := hire()
	e.(*salesperson).sales()
}

type iEmployee interface {
	report() string

	retire()
}

type salesperson struct {
	id   int
	name string
}

func (s *salesperson) report() string {
	return "report"
}

func (s *salesperson) retire() {
	// do nothing
}

func (s *salesperson) sales() {
	print("sales")
}

func hire() iEmployee {
	return &salesperson{
		id:   1,
		name: "Alice",
	}
}

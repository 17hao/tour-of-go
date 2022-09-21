package main

type address struct {
	state string
	city  *string
}

func main() {
	city := "c"
	a := address{state: "a", city: &city}
	b := address{state: "b", city: &city}
	c := address{state: "c", city: &city}
	d := address{state: "d", city: &city}
	addresses := []address{a, b, c, d}
	for _, addr := range addresses {
		modifyProperty(&addr.state)
		modifyProperty(addr.city)
	}
	for _, a := range addresses {
		println(a.state)
		println(*a.city)
	}
}

func modifyProperty(p *string) {
	*p = "updated"
}

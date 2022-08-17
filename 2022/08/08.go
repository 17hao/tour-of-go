package main

func main() {
	m := map[string]string{"k": "v"}
	delete(m, "key")
	for k, v := range m {
		println(k, v)
	}
}

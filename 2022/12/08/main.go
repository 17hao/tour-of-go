package main

func main() {
	s1 := myStruct{id: 1}
	s2 := myStruct{id: 2}
	s3 := myStruct{id: 3}
	s4 := myStruct{id: 4}
	s5 := myStruct{id: 5}
	variableFunc(s1, s2, s3, s4, s5)
}

type myStruct struct {
	id int
}

func variableFunc(myStructs ...myStruct) {
	for _, s := range myStructs {
		println(s.id)
	}
}

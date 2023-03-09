package main

import (
	"fmt"
	"log"
	"os"
)

type myStruct struct {
	pstr *string
	i1   int
	i2   int
}

func main() {
	//str := "str"
	//s := myStruct{
	//	pstr: &str,
	//	i1:   1,
	//	i2:   2,
	//}
	//fmt.Printf("%+v\n", s)
	//
	//newS := f1(s)
	//fmt.Printf("%+v\n", newS)

	//s1 := []int{0, 1, 2, 3, 4, 5}
	//s2 := s1[1:3]
	//fmt.Printf("len(s2)=%d cap(s2)=%d", len(s2), cap(s2))

	path := []byte("/usr/bin/bash")
	toUpper(path)
	fmt.Printf("%s\n", string(path))

	fmt.Printf("Before len(path)=%d cap(path)=%d\n", len(path), cap(path))
	newPath := subtractOneFromLength(path)
	fmt.Printf("After len(path)=%d cap(path)=%d\n", len(path), cap(path))
	fmt.Printf("After len(newPath)=%d cap(newPath)=%d\n", len(newPath), cap(newPath))

	f1()
}

func toUpper(s []byte) {
	for i, b := range s {
		if 'a' <= b && b <= 'z' {
			s[i] = b + 'A' - 'a'
		}
	}
}

func subtractOneFromLength(s []byte) []byte {
	s = s[0 : len(s)-1]
	return s
}

func f1() {
	f, err := os.Open("/Users/bytedance/test")
	if err != nil {
		log.Fatalf("open err:%+v", err)
	}
	buf := make([]byte, 1, 32)
	count := 0
	for i := 0; i < 32; i++ {
		n, err := f.Read(buf[i : i+1])
		fmt.Printf("buf=%s cap(buf)=%d len(buf)=%d\n", string(buf), cap(buf), len(buf))
		if err != nil {
			log.Fatal(err)
		}
		count += n
	}
	fmt.Printf("count=%d\n", count)
}

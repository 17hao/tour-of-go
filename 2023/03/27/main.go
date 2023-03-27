package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	ss := make([]string, 0, 5)
	ss = append(ss, "0")
	ss = append(ss, "1")
	ss = append(ss, "2")
	f1(ss)
}

func f1(ss []string) {
	fmt.Println("before f2")

	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&ss))
	fmt.Printf("[f1] len=%d cap=%d hdr=%+v %+v\n", len(ss), cap(ss), hdr, ss)

	f2(ss)
	fmt.Println("after f2")

	hdr = (*reflect.SliceHeader)(unsafe.Pointer(&ss))
	fmt.Printf("[f1] len=%d cap=%d hdr=%+v %+v\n", len(ss), cap(ss), hdr, ss)

	reverse(ss)
	hdr = (*reflect.SliceHeader)(unsafe.Pointer(&ss))
	fmt.Printf("[f1] len=%d cap=%d hdr=%+v %+v\n", len(ss), cap(ss), hdr, ss)
}

func f2(ss []string) {
	ss = append(ss, "100")
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&ss))
	fmt.Printf("[f2] len=%d cap=%d hdr=%+v %+v\n", len(ss), cap(ss), hdr, ss)
}

func reverse(ss []string) {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&ss))
	for i := 0; i < len(ss)/2; i++ {
		tmp := ss[i]
		ss[i] = ss[len(ss)-1-i]
		ss[len(ss)-1-i] = tmp
	}
	fmt.Printf("[reverse] len=%d cap=%d hdr=%+v %+v\n", len(ss), cap(ss), hdr, ss)
}

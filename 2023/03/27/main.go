package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	ss := make([]string, 0, 3)
	ss = append(ss, "0")
	ss = append(ss, "1")
	ss = append(ss, "2")
	f1(ss)
}

func f1(ss []string) {
	fmt.Println("before f2")

	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&ss))
	fmt.Printf("[f1] len=%d cap=%d hdr=%+v %+v\n", len(ss), cap(ss), hdr, ss)

	fmt.Println("after f2")
	f2(ss)

	hdr = (*reflect.SliceHeader)(unsafe.Pointer(&ss))
	fmt.Printf("[f1] len=%d cap=%d hdr=%+v %+v\n", len(ss), cap(ss), hdr, ss)

	f3(ss)
	hdr = (*reflect.SliceHeader)(unsafe.Pointer(&ss))
	fmt.Printf("[f1] len=%d cap=%d hdr=%+v %+v\n", len(ss), cap(ss), hdr, ss)

	f4(ss)
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

func f3(ss []string) {
	news := ss[0:]
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&news))
	fmt.Printf("[f3] len=%d cap=%d hdr=%+v %+v\n", len(news), cap(news), hdr, news)
}

func f4(ss []string) {
	news := make([]string, 0)
	for _, s := range ss {
		news = append(news, s)
	}
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&news))
	fmt.Printf("[f4] len=%d cap=%d hdr=%+v %+v\n", len(news), cap(news), hdr, news)
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

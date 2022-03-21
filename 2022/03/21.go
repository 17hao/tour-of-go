package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	testSlice()
	//test()
	//fmt.Println(math.Ceil(1.5))
	fmt.Println(unsafe.Sizeof(int(1)))
}

func test() {
	//imem := uintptr(1000)
	fmt.Println(^uintptr(0) >> 63)
	fmt.Println(4 << (^uintptr(0) >> 63))
}

func testSlice() {
	slice := make([]int, 0)
	//fmt.Printf("%+v\n", (*reflect.SliceHeader)(unsafe.Pointer(&slice)))

	//slice = append(slice, 1) //①
	//slice = append(slice, 1, 2) //②
	//slice = append(slice, 1, 2, 3) //③
	//slice = append(slice, 1, 2, 3, 4) //④
	//slice = append(slice, 1, 2, 3, 4, 5) //⑤
	//slice = append(slice, 1, 2, 3, 4, 5, 6) //⑥
	//slice = append(slice, 1, 2, 3, 4, 5, 6, 7) //⑥
	//slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9) //⑥
	//slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11) //⑥
	slice = append(
		slice,
		1, 2, 3, 4, 5, 6, 7, 8,
		1, 2, 3, 4, 5, 6, 7, 8,
		1, 2, 3, 4, 5, 6, 7, 8,
		1, 2, 3, 4, 5, 6, 7, 8,
		1, 2, 3, 4,
	)
	fmt.Printf("%+v\n", (*reflect.SliceHeader)(unsafe.Pointer(&slice)))

	slice2 := append(slice, 1)
	fmt.Printf("%+v\n", (*reflect.SliceHeader)(unsafe.Pointer(&slice2)))
	slice3 := append(slice, 2)
	fmt.Printf("%+v\n", (*reflect.SliceHeader)(unsafe.Pointer(&slice3)))

	for _, i := range slice2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
	for _, i := range slice3 {
		fmt.Print(i, " ")
	}

	//slice := make([]int, 0)
	//for i := 0; i < 6 ; i++ {
	//	slice = append(slice, i)
	//	fmt.Printf("%+v\n", (*reflect.SliceHeader)(unsafe.Pointer(&slice)))
	//}
}

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	/*
		type slice struct {
			array unsafe.Pointer
			len   int
			cap   int
		}
	*/

	intSlice := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		//intSlice[i] = i panic
		intSlice = append(intSlice, i)
	}

	fmt.Println("-------------------------------------------------------")
	fmt.Printf("| intSlice slice %v, len: %d, cap: %d |\n", intSlice, len(intSlice), cap(intSlice))
	fmt.Println("-------------------------------------------------------")

	cut := intSlice[2:4]
	fmt.Printf("| cut slice %v, len: %d, cap: %d |\n", cut, len(cut), cap(cut))

	fmt.Println("-------------------------------------------------------")
	cut = cut[:cap(cut)]
	fmt.Printf("| cut slice %v, len: %d, cap: %d |\n", cut, len(cut), cap(cut))
	fmt.Println("-------------------------------------------------------")

	fmt.Printf(
		"| intSlice: %d, cut: %d |\n"+
			"| intSlice: %v, cut: %v |\n",
		reflect.ValueOf(intSlice).Pointer(),
		reflect.ValueOf(cut).Pointer(),
		*(*reflect.SliceHeader)(unsafe.Pointer(&intSlice)),
		*(*reflect.SliceHeader)(unsafe.Pointer(&cut)),
	)
	fmt.Println("-------------------------------------------------------")

	cut[0] = 1 << 32 // 2**32
	fmt.Printf("| intSlice slice %v, len: %d, cap: %d |\n", intSlice, len(intSlice), cap(intSlice))
	fmt.Printf("| cut slice %v, len: %d, cap: %d |\n", cut, len(cut), cap(cut))
	fmt.Printf(
		"| intSlice: %d, cut: %d |\n"+
			"| intSlice: %v, cut: %v |\n",
		reflect.ValueOf(intSlice).Pointer(),
		reflect.ValueOf(cut).Pointer(),
		*(*reflect.SliceHeader)(unsafe.Pointer(&intSlice)),
		*(*reflect.SliceHeader)(unsafe.Pointer(&cut)),
	)
	fmt.Println("-------------------------------------------------------")

	cut = append(cut, (1<<32)+1)
	cut[0] = 1 << 10
	fmt.Printf("| intSlice slice %v, len: %d, cap: %d |\n", intSlice, len(intSlice), cap(intSlice))
	fmt.Printf("| cut slice %v, len: %d, cap: %d |\n", cut, len(cut), cap(cut))
	fmt.Printf(
		"| intSlice: %d, cut: %d |\n"+
			"| intSlice: %v, cut: %v |\n",
		reflect.ValueOf(intSlice).Pointer(),
		reflect.ValueOf(cut).Pointer(),
		*(*reflect.SliceHeader)(unsafe.Pointer(&intSlice)),
		*(*reflect.SliceHeader)(unsafe.Pointer(&cut)),
	)
	fmt.Println("-------------------------------------------------------")
}

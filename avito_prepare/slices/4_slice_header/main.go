package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	slice := make([]int, 0, 3)
	slice = append(slice, 1<<10)
	fmt.Println(slice)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	fmt.Println(reflect.TypeOf(slice), reflect.TypeOf(sliceHeader))
	fmt.Println("-------------------------------------------------------")
	fmt.Printf("| slice: %v, size: %d, sliceHeader: %v |\n", slice, unsafe.Sizeof(slice), sliceHeader)
	fmt.Println("-------------------------------------------------------")

	func(inLambdaSlice []int) {
		lambdaHeader := (*reflect.SliceHeader)(unsafe.Pointer(&inLambdaSlice))
		fmt.Printf("| in lambda: sliceHeader was copied BY VALUE: %v |\n", lambdaHeader)

		inLambdaSlice[0] = 1 << 11
		inLambdaSlice = append(inLambdaSlice, 1<<16)
		inLambdaSlice = append(inLambdaSlice, 1<<32)
		inLambdaSlice = append(inLambdaSlice, 1<<32)
		inLambdaSlice = append(inLambdaSlice, 1<<32)

		fmt.Printf("| in lambda: sliceHeader after appending: %v |\n", lambdaHeader)
		fmt.Printf("| in lambda: slice: %v |\n", inLambdaSlice)
		fmt.Println("-------------------------------------------------------")
	}(slice)

	fmt.Printf("| slice: %v, size: %d, sliceHeader: %v |\n", slice, unsafe.Sizeof(slice), sliceHeader)
	slice = slice[:cap(slice)]
	fmt.Printf("| slice: %v, size: %d, sliceHeader: %v |\n", slice, unsafe.Sizeof(slice), sliceHeader)

	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	sliceCopy[0] = 0
	fmt.Println(sliceCopy, slice)
}

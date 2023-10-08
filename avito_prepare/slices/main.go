package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 257)
	fmt.Println(len(slice), cap(slice))
	slice = append(slice, 1)
	fmt.Println(len(slice), cap(slice))
	fmt.Println(257 * 1.25)
}

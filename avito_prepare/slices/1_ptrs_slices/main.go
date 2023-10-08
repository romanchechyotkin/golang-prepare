package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 1, 2)
	fmt.Println(nums, len(nums), cap(nums)) // [0] 1 2

	appendSlice(nums, 1024)
	fmt.Println(nums, len(nums), cap(nums)) // [0] 1 2

	changeSLice(nums, 0, 2) // panic out of range
	fmt.Println(nums, len(nums), cap(nums))
}

func appendSlice(arr []int, val int) {
	arr = append(arr, val)
}

func changeSLice(arr []int, idx, val int) {
	arr[idx] = val
}

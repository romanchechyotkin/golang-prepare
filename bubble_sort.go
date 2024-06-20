package main

import "log"

func main() {
	a := []int{3, 4, 5, 1, 2}

	log.Println(sort(a))
}

func sort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-(i+1); j++ {
			left := nums[j]
			right := nums[j+1]

			if left > right {
				nums[j] = right
				nums[j+1] = left
			}
		}
	}

	return nums
}

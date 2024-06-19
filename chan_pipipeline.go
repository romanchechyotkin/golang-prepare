package main

import "fmt"

func main() {
	for n := range square(pipe(1, 2, 3, 4, 5)) {
		fmt.Println(n)
	}
}

func pipe(nums ...int) <-chan int {
	res := make(chan int)

	go func() {
		defer close(res)

		for _, num := range nums {
			res <- num
		}

	}()

	return res
}

func square(in <-chan int) <-chan int {
	res := make(chan int)

	go func() {
		defer close(res)

		for n := range in {
			res <- n * n
		}
	}()

	return res
}

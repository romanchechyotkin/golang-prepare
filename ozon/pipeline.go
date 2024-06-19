package main

import "log"

func main() {
	for i := range square(gen(1, 2, 3, 4, 5)) {
		log.Println(i)
	}
}

func gen(num ...int) chan int {
	res := make(chan int)

	go func() {
		defer close(res)

		for _, n := range num {
			res <- n
		}
	}()

	return res
}

func square(in chan int) chan int {
	res := make(chan int)

	go func() {
		defer close(res)

		for i := range in {
			res <- i * i
		}
	}()

	return res
}

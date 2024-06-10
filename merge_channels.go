package main

import (
	"fmt"
	"sync"
)

func merge(channels ...<-chan int) <-chan int {
	res := make(chan int)

	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for n := range ch {
				res <- n
			}
		}()
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func main() {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	ch2 <- 4
	ch2 <- 5

	close(ch1)
	close(ch2)

	for res := range merge(ch1, ch2) {
		fmt.Println(res)
	}
}

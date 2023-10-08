package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)
	ch1 <- 2
	ch2 <- 1
	ch1 <- 31
	ch2 <- 0
	ch1 <- 3
	ch2 <- 2
	close(ch1)
	close(ch2)

	ch3 := merge[int](ch1, ch2)

	for v := range ch3 {
		fmt.Println(v)
	}
}

func merge[T any](chans ...chan T) chan T {
	result := make(chan T)
	wg := sync.WaitGroup{}
	for _, singleChan := range chans {
		wg.Add(1)
		singleChan := singleChan
		go func() {
			defer wg.Done()
			for val := range singleChan {
				result <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

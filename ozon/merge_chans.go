package main

import (
	"log"
	"sync"
)

func main() {
	a := make(chan int)
	b := make(chan int)

	go func() {
		defer close(a)
		defer close(b)

		for i := 0; i < 5; i++ {
			a <- i
			b <- i * 2
		}
	}()

	for n := range merge(a, b) {
		log.Println(n)
	}

}

func merge(chans ...chan int) chan int {
	res := make(chan int)

	var wg sync.WaitGroup

	for _, ch := range chans {
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

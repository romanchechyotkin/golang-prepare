package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 20
	wg := sync.WaitGroup{}

	for i := 1; i <= counter; i++ {
		i := i
		wg.Add(1)
		go func() {
			fmt.Println(i * i)
			wg.Done()
		}()
	}

	wg.Wait()
}

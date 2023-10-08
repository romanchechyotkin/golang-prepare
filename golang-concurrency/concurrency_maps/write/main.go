package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	writes := 1000
	storage := make(map[int]int, writes)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	now := time.Now()
	wg.Add(writes)

	for i := 0; i < writes; i++ {
		i := i
		go func() {
			mu.Lock()
			storage[i] = i
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(storage)
	fmt.Println(time.Now().Sub(now))
}

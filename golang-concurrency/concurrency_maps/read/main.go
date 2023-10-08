package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ops := 1000
	storage := make(map[int]int, ops)

	now := time.Now()

	mu := sync.RWMutex{}
	wg := sync.WaitGroup{}

	wg.Add(ops)
	for i := 0; i < ops; i++ {
		i := i
		go func() {
			mu.Lock()
			storage[i] = i
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Add(ops)
	for i := 0; i < ops; i++ {
		i := i
		go func() {
			mu.RLock()
			fmt.Println(storage[i])
			mu.RUnlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(time.Now().Sub(now))
}

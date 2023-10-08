package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

func main() {
	alreadyStored := make(map[int]struct{})
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	capacity := 1000

	doubles := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10))
	}
	log.Println(doubles)
	// 1, 3, 3, 6, 7, 8, 9, 9, 6, 7, 1

	uniqIDs := make(chan int, capacity)

	for i := 0; i < capacity; i++ {
		i := i
		wg.Add(1)

		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			if _, ok := alreadyStored[doubles[i]]; !ok {
				alreadyStored[doubles[i]] = struct{}{}
				uniqIDs <- doubles[i]
			}
		}()
	}

	close(uniqIDs)
	wg.Wait()

	for id := range uniqIDs {
		fmt.Println(id)
	}
}

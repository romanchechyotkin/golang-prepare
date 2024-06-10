package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	for k := range m {
		wg.Add(1)

		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			mu.Lock()
			m[k]++
			mu.Unlock()
		}()
	}

	wg.Wait()

	log.Println(m)
}

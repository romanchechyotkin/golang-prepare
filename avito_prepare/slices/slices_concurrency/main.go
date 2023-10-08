package main

import (
	"fmt"
	"sync"
)

func main() {
	var links = [10000]string{}
	var counter int
	ch := make(chan string, len(links))

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, link := range links {
		ch <- link
	}
	close(ch)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for l := range ch {
				if err := checkURL(l); err == nil {
					mu.Lock()
					counter++
					mu.Unlock()
				}
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

func checkURL(url string) error {
	return nil
}

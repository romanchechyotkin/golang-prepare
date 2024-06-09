package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	res := make(chan int)

	for i := range 5 {
		_ = i
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(ctx, res)
		}()
	}

	for {
		select {
		case num, ok := <-res:
			if !ok {
				log.Println("closed")
				return
			}
			log.Println("got", num)

			if num == 5 {
				cancel()
				close(res)
			}
		}
	}

	wg.Wait()

}

func worker(ctx context.Context, ch chan int) {
	rand.Seed(time.Now().Unix())
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ctx.Done():
			log.Println("done")
			return
		case <-ticker.C:
			num := rand.Intn(25)
			ch <- num
		}
	}

}

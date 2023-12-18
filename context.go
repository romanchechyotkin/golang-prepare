package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	//context.Background() same with context.TODO()

	ctx := context.Background()
	fmt.Println(ctx.Err())
	deadline, ok := ctx.Deadline()
	fmt.Println(deadline, ok)

	ctx = context.TODO()
	fmt.Println(ctx.Err())
	deadline, ok = ctx.Deadline()
	fmt.Println(deadline, ok)

	// ------------------------------------------------------------------------------------------

	// values in context.WithValue() not stored in map, this is TREE, check debugger
	ctx = context.Background()
	ctx = context.WithValue(ctx, "1", 1)
	ctx = context.WithValue(ctx, "2", 2)

	fmt.Println(ctx.Value("1"))
	fmt.Println(ctx.Value("2"))
	fmt.Println(ctx.Value("0"))

	// ------------------------------------------------------------------------------------------

	ctx = context.Background()
	ctx, cancel := context.WithCancel(ctx)
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			work(ctx, i)
		}()
	}

	time.AfterFunc(time.Second, func() {
		cancel()
	})
	wg.Wait()

	log.Println("completed")
}

func work(ctx context.Context, i int) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	slowFn(ctx, i)
}

func slowFn(ctx context.Context, i int) {
	ctx = context.WithValue(ctx, "one", 1)
	ctx = context.WithValue(ctx, "two", 2)

	log.Printf("slow function %d started\n", i)

	select {
	case <-time.Tick(3 * time.Second):
		log.Printf("slow function %d finished\n", i)
	case <-ctx.Done():
		log.Printf("slow function %d too slow, err: %s\n", i, ctx.Err())
	}
}

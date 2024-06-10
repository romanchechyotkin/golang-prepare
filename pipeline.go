package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, num := range nums {
			out <- num
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for num := range in {
			out <- num * num
		}
	}()

	return out
}

func worker(id int, jobQueue <-chan int, res chan<- int) {
	log.Println("worker", id, "starting")

	for num := range jobQueue {
		log.Println("worker", id, "processing job", num)
		res <- num * num
		time.Sleep(10 * time.Millisecond * time.Duration(num))
	}
}

func main() {
	log.Println("pipeline")

	for num := range square(gen(1, 2, 3, 4, 5)) {
		fmt.Println(num)
	}

	log.Println("-----------------------------------")
	log.Println("worker pool")

	//var wg sync.WaitGroup
	const workerNum = 5
	const jobNums = 25

	var wg sync.WaitGroup
	jobQueue := make(chan int, jobNums)
	resultQueue := make(chan int, jobNums)

	for a := 0; a < workerNum; a++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(a, jobQueue, resultQueue)
		}()
	}

	for a := 1; a <= jobNums; a++ {
		jobQueue <- a
	}
	close(jobQueue)

	go func() {
		wg.Wait()
		close(resultQueue)
	}()

	for a := range resultQueue {
		fmt.Println(a)
	}

	//for a := 1; a < jobNums; a++ {
	//	<-resultQueue
	//}

}

package main

import (
	"log"
	"net/http"
)

func main() {
	const workerNum = 5
	const jobNums = 25

	jobsQueue := make(chan int, workerNum)
	res := make(chan int)

	for i := 0; i < workerNum; i++ {
		go worker(i, jobsQueue, res)
	}

	go func() {
		defer close(jobsQueue)
		for i := 0; i < jobNums; i++ {
			jobsQueue <- i
		}
	}()

	for i := 0; i < jobNums; i++ {
		log.Println(<-res)
	}

	close(res)
}

func worker(id int, jobsQueue <-chan int, res chan<- int) {
	log.Printf("worker %d started\n", id)

	for j := range jobsQueue {
		log.Printf("worker %d; job %d\n", id, j)
		res <- download()
	}
}

func download() int {
	resp, err := http.Get("https://example.com")
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	return resp.StatusCode
}

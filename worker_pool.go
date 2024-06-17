package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type WorkerPool interface {
	Start()
	Stop()
}

type Pool struct {
	workerNum uint

	jobQueue chan int
	res      chan int
}

func NewPool(workerNum uint) *Pool {
	jobQueue := make(chan int, workerNum*2)
	res := make(chan int)

	return &Pool{
		workerNum: workerNum,
		jobQueue:  jobQueue,
		res:       res,
	}
}

func (p *Pool) worker(id int) {
	for j := range p.jobQueue {
		log.Printf("worker %d job %d\n", id, j)

		p.download()
	}
}

func (p *Pool) download() {
	response, err := http.Get("https://example.com")
	if err != nil {
		return
	}
	defer response.Body.Close()

	p.res <- response.StatusCode
}

func (p *Pool) Start() {
	for i := 0; i < int(p.workerNum); i++ {
		go p.worker(i)
	}
}

func (p *Pool) Stop() {
	close(p.jobQueue)
	close(p.res)
}

func main() {
	pool := NewPool(4)
	pool.Start()

	go func() {
		for i := 0; i < 12; i++ {
			pool.jobQueue <- i
		}
	}()

	for i := 0; i < 12; i++ {
		log.Println(i, <-pool.res)
	}

	log.Println("----------------------")

	a := make(chan int, 5)
	b := make(chan int, 5)

	go func() {
		for i := range mergeChans(a, b) {
			log.Println(i)
		}
	}()

	for i := 0; i < 4; i++ {
		a <- i
		b <- i
		time.Sleep(time.Second * 1 * time.Duration(i))
		a <- 12
	}

	close(a)
	close(b)

}

func mergeChans(chans ...<-chan int) <-chan int {
	res := make(chan int)

	var wg sync.WaitGroup

	for _, ch := range chans {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for i := range ch {
				res <- i
			}
		}()
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

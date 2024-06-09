package main

import (
	"log"
	"net/http"
)

func main() {
	workersNum := 100
	itarations := 1000
	ch := make(chan int, 100)
	res := make(chan int)

	for i := 0; i < workersNum; i++ {
		go worker(i, ch, res)
	}

	go func() {
		for i := 0; i < itarations; i++ {
			ch <- i
		}
	}()

	for i := 0; i < itarations; i++ {
		r := <-res
		log.Println(r)
	}

	close(res)
	close(ch)

}

func worker(id int, jobs <-chan int, res chan<- int) {
	for j := range jobs {
		log.Println("worker", id, "start", j)
		resp, err := http.Get("https://www.example.com")
		if err != nil {
			log.Println(err)
			continue
		}
		defer resp.Body.Close()

		res <- resp.StatusCode
	}
}

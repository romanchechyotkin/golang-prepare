package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	websites := []string{
		"https://hackernoon.com/",
		"https://github.com/",
		"https://apple.com/",
		"https://google.com/",
		"https://youtube.com/",
		"https://www.udemy.com/",
		"https://netflix.com/",
		"https://www.coursera.org/",
		"https://facebook.com/",
		"https://microsoft.com",
		"https://wikipedia.org",
		"https://educative.io",
		"https://acloudguru.com",
	}

	now := time.Now()

	resources := make(chan string, 5)
	results := make(chan string)

	for i := 0; i < 5; i++ {
		go worker(resources, results)
	}

	go func() {
		for _, website := range websites {
			resources <- website
		}
	}()

	for i := 0; i < len(websites); i++ {
		fmt.Println(<-results)
	}

	fmt.Println(time.Now().Sub(now))
}

func worker(resources, results chan string) {
	for resource := range resources {
		if res, err := http.Get(resource); err != nil {
			results <- resource + " is down"
		} else {
			results <- fmt.Sprintf("[%d] %s is up", res.StatusCode, resource)
		}
	}
}

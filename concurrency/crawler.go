package main

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type SiteContent struct {
	Title   string
	Date    time.Time
	Content string
}

func worker(urlChan <-chan string, resultChan chan<- SiteContent) {
	for url := range urlChan {
		content := DownloadSiteContent(context.Background(), url)
		resultChan <- content
	}
}

func DownloadSiteContent(ctx context.Context, url string) SiteContent {
	rand.Seed(time.Now().UnixNano())
	sleepTime := time.Duration(rand.Intn(6)+5) * time.Second

	time.Sleep(sleepTime)

	return SiteContent{
		Title:   "Заголовок" + url,
		Date:    time.Now(),
		Content: "Контент" + url,
	}
}

func ParallelDownload(ctx context.Context, urls <-chan string, numWorkers int) map[string]SiteContent {
	var wg sync.WaitGroup
	var mu sync.Mutex

	m := make(map[string]SiteContent)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					log.Println(ctx.Err())
					return

				case url := <-urls:
					sc := DownloadSiteContent(ctx, url)
					mu.Lock()
					m[url] = sc
					mu.Unlock()
				}
			}
		}()
	}

	wg.Wait()

	return m
}

func main() {
	//bufSize := 3
	//urls := make(chan string, bufSize)
	//resultChan := make(chan SiteContent)
	//
	//for i := 0; i < 10; i++ {
	//	go worker(urls, resultChan)
	//}
	//
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		urls <- "https://www.baidu.com"
	//	}
	//}()
	//
	//for i := 0; i < 10; i++ {
	//	res := <-resultChan
	//	fmt.Printf("%d %s %s %s \n", i, res.Title, res.Content, res.Date.String())
	//}
	//
	//close(urls)
	//close(resultChan)

	bufSize := 3
	urls := make(chan string, bufSize)
	//resultChan := make(chan SiteContent)

	go func() {
		for i := range 5 {
			urls <- "https://example.com" + strconv.Itoa(i)
		}
		close(urls)
	}()

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()

	results := ParallelDownload(ctx, urls, bufSize)

	time.AfterFunc(time.Second*5, func() {
		cancelFunc()
	})

	for url, content := range results {
		log.Println(url, content.Title, content.Date, content.Content)
	}
}

package main

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	MaxGoroutines = 1
	Timeout       = 2 * time.Second
)

type SiteStatus struct {
	Name           string
	StatusCode     int
	TimeOfRequest  time.Time
	RequestLatency time.Duration
}

type Monitor struct {
	StatusMap        map[string]SiteStatus
	Sites            []string
	RequestFrequency time.Duration
	G                errgroup.Group
	Mu               *sync.Mutex
}

func main() {
	wg := sync.WaitGroup{}

	var sites = []string{
		"https://youtube.com",
		"https://habr.com",
		"https://twitch.com",
		"https://vk.com",
		"https://ya.ru",
	}

	monitor := NewMonitor(sites, 5*time.Second)
	err := monitor.Run(context.Background(), &wg)
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}

func NewMonitor(sites []string, frequency time.Duration) *Monitor {
	return &Monitor{
		StatusMap:        make(map[string]SiteStatus),
		Sites:            sites,
		RequestFrequency: frequency,
		Mu:               &sync.Mutex{},
	}
}

// Run printStatuses and checkSite in diff goroutines
func (m *Monitor) Run(ctx context.Context, wg *sync.WaitGroup) error {
	ticker := time.NewTicker(m.RequestFrequency)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				for _, site := range m.Sites {
					err := m.checkSite(ctx, site)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				m.printStatuses()
			}
		}
	}()

	return nil
}

// checkSite makes request to site and write result of request to StatusMap
func (m *Monitor) checkSite(ctx context.Context, site string) error {
	now := time.Now()

	resp, err := http.Get(site)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	m.Mu.Lock()
	m.StatusMap[site] = SiteStatus{
		Name:           site,
		StatusCode:     resp.StatusCode,
		RequestLatency: time.Since(now),
		TimeOfRequest:  time.Now(),
	}
	m.Mu.Unlock()

	return nil
}

// printStatuses iterate over map and print results
func (m *Monitor) printStatuses() {
	m.Mu.Lock()
	for k, v := range m.StatusMap {
		log.Printf("site %s; status code %d; latency %s; time %s\n", k, v.StatusCode, v.RequestLatency, v.TimeOfRequest)
	}
	m.Mu.Unlock()
}

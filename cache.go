package main

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

type Cache struct {
	storage map[string]*item
	mu      sync.RWMutex
}

type item struct {
	val            any
	expirationTime int64
}

func NewCache(ctx context.Context, expirePeriod time.Duration) (*Cache, error) {
	if expirePeriod == 0 {
		return nil, errors.New("wrong expiration period")
	}

	c := &Cache{
		storage: make(map[string]*item),
	}

	go c.cleanCache(ctx, expirePeriod)
	return c, nil
}

func (c *Cache) cleanCache(ctx context.Context, period time.Duration) {
	ticker := time.NewTicker(period)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			now := time.Now().UnixNano()

			c.mu.Lock()
			for k, v := range c.storage {
				if v.expirationTime < now {
					delete(c.storage, k)
				}
			}

			c.mu.Unlock()
		}
	}

}

func (c *Cache) Set(key string, val any, expirationTime time.Duration) error {
	if expirationTime <= 0 {
		return errors.New("ttl must be >= 0")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	expire := time.Now().Add(expirationTime).UnixNano()
	c.storage[key] = &item{
		val:            val,
		expirationTime: expire,
	}

	return nil
}

func (c *Cache) Get(key string) (any, error) {
	c.mu.RLock()

	val, ok := c.storage[key]
	if !ok {
		c.mu.Unlock()
		return nil, errors.New("not found")
	}

	if val.expirationTime <= time.Now().UnixNano() {
		c.mu.RUnlock()

		c.mu.Lock()
		delete(c.storage, key)
		c.mu.Unlock()
		return nil, errors.New("not found")
	}

	c.mu.RUnlock()
	return val.val, nil
}

func main() {
	c, err := NewCache(context.Background(), 5*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	c.Set("test", 123, time.Second)
	c.Set("test2", 123, 10*time.Second)
}

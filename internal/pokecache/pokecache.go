package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]cacheEntry
	mu   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		data: make(map[string]cacheEntry),
		mu:   &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, value []byte) {
	entry := cacheEntry{createdAt: time.Now().UTC(), val: value}
	c.mu.Lock()
	c.data[key] = entry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	value, ok := c.data[key]
	c.mu.Unlock()

	return value.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		for key, val := range c.data {
			if time.Since(val.createdAt) > interval {
				c.mu.Lock()
				delete(c.data, key)
				c.mu.Unlock()
			}
		}
	}
}

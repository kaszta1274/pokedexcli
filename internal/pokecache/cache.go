package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu           sync.Mutex
	interval     time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		cacheEntries: map[string]cacheEntry{},
		interval:     interval,
	}

	go c.reapLoop()

	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheEntries[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exitst := c.cacheEntries[key]
	if !exitst {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	for range time.NewTicker(c.interval).C {
		c.mu.Lock()
		for key, entry := range c.cacheEntries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cacheEntries, key)
			}
		}
		c.mu.Unlock()
	}
}

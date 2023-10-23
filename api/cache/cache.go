package cache

import (
	"sync"
	"time"
)

type entry struct {
	val []byte
	timeCached time.Time 
}

type Cache struct {
	cache map[string]entry
	mux *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c :=  Cache {
		cache: make(map[string]entry),
		mux: &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Set(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = entry{
		val: val,
		timeCached: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mux.Lock()
	defer c.mux.Unlock()
	ent, ok := c.cache[key]
	return ent.val, ok
}

func (c *Cache) reapLoop(interval time.Duration){
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}
func (c *Cache) reap(interval time.Duration){
	c.mux.Lock()
	defer c.mux.Unlock()
	timeBeforeInterval := time.Now().UTC().Add(-interval)
	for key, val := range c.cache {
		if val.timeCached.Before(timeBeforeInterval) {
			delete(c.cache, key)
		}
	}
}
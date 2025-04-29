package pokecache

import (
	"sync"
	"time"
)


type Cache struct {
	items map[string]cacheEntry
	mu sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		items: make(map[string]cacheEntry),
	}

	go cache.reapLoop(interval)

	return cache
}

func (cache *Cache) Add(key string, val []byte)  {
	cache.mu.Lock();
	defer cache.mu.Unlock();

	cache.items[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	entry, ok := cache.items[key]
	cache.mu.Unlock()
	
	if !ok {
		return []byte{}, false
	}

	return entry.val, true
}

func (cache *Cache) reapLoop(d time.Duration)  {
	ticker := time.NewTicker(d)

		select {
			case t := <- ticker.C:
				for key, value := range cache.items {
					if value.createdAt.Add(d).Before(t) {
						cache.mu.Lock()
						delete(cache.items, key)
						cache.mu.Unlock()
					}
				}
		}
}


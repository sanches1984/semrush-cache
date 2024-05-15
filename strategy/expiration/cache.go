package expiration

import (
	"github.com/sanches1984/semrush-cache/model"
	"sync"
	"time"
)

type cache struct {
	items      map[string]model.CacheItem
	mutex      sync.Mutex
	capacity   int
	expiration time.Duration
}

func NewCache(capacity int, expiration time.Duration) *cache {
	return &cache{
		items:      make(map[string]model.CacheItem),
		mutex:      sync.Mutex{},
		capacity:   capacity,
		expiration: expiration,
	}
}

func (c *cache) Set(key string, value interface{}) {
	if c.Size() >= c.capacity {
		c.cleanup()
	}

	var expiration int64
	if c.expiration > 0 {
		expiration = time.Now().Add(c.expiration).UnixNano()
	}

	c.mutex.Lock()
	c.items[key] = model.CacheItem{
		Value:      value,
		Expiration: expiration,
	}
	c.mutex.Unlock()
}

// Get получает значение из кэша по ключу
func (c *cache) Get(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	item, found := c.items[key]
	return item.Value, found
}

// Delete удаляет значение из кэша по ключу
func (c *cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.items, key)
}
func (c *cache) Size() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return len(c.items)
}

func (c *cache) cleanup() {
	now := time.Now().UnixNano()
	c.mutex.Lock()
	for key, item := range c.items {
		if item.Expiration > 0 && now > item.Expiration {
			delete(c.items, key)
		}
	}
	c.mutex.Unlock()
}

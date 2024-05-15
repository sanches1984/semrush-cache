package semrush_cache

import (
	"math"
	"sync"
	"time"
)

type StrategyType string
type Option func(c *cache)

const (
	StrategyLeastRecentlyUsed   StrategyType = "recently_used"
	StrategyLeastFrequentlyUsed StrategyType = "frequently_used"
	StrategyTimeBasedExpiration StrategyType = "time_based_expiration"
)

type Cache interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
	Delete(key string)
	Size() int
}

type cache struct {
	strategy   StrategyType
	maxSize    int
	expiration time.Duration

	items map[string]CacheItem
	mutex sync.RWMutex
}

func New(strategy StrategyType, options ...Option) Cache {
	c := &cache{
		strategy: strategy,
		maxSize:  math.MaxInt,
		mutex:    sync.RWMutex{},
	}

	for _, opt := range options {
		opt(c)
	}

	c.items = make(map[string]CacheItem, c.maxSize)
	return c
}

func (c *cache) Set(key string, value interface{}) {
	if c.Size() == c.maxSize {
		c.removeItem()
	}

	c.mutex.Lock()
	c.items[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(c.expiration),
	}
	c.mutex.Unlock()
}

func (c *cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	item, found := c.items[key]
	c.mutex.RUnlock()
	return item, found
}

func (c *cache) Delete(key string) {
	c.mutex.Lock()
	delete(c.items, key)
	c.mutex.Unlock()
}

func (c *cache) Size() int {
	c.mutex.RLock()
	size := len(c.items)
	c.mutex.RUnlock()
	return size
}

func (c *cache) removeItem() {
	switch c.strategy {
	case StrategyTimeBasedExpiration:
	// todo
	case StrategyLeastFrequentlyUsed:
	// todo
	case StrategyLeastRecentlyUsed:
		// todo
	}
}

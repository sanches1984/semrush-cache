package semrush_cache

import (
	"github.com/sanches1984/semrush-cache/strategy/expiration"
	"github.com/sanches1984/semrush-cache/strategy/frequently_used"
	"github.com/sanches1984/semrush-cache/strategy/recently_used"
	"log"
	"math"
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
	strategy   Cache
	capacity   int
	expiration time.Duration
	logger     *log.Logger
}

func New(strategy StrategyType, options ...Option) Cache {
	c := &cache{
		capacity: math.MaxInt,
		logger:   log.Default(),
	}

	for _, opt := range options {
		opt(c)
	}

	c.logger.Printf("init strategy: %s", strategy)
	c.initCacheByStrategy(strategy)
	return c
}

func (c *cache) Set(key string, value interface{}) {
	c.logger.Printf("set key: %s", key)
	c.strategy.Set(key, value)
}

func (c *cache) Get(key string) (interface{}, bool) {
	c.logger.Printf("get key: %s", key)
	return c.strategy.Get(key)
}

func (c *cache) Delete(key string) {
	c.logger.Printf("delete key: %s", key)
	c.strategy.Delete(key)
}

func (c *cache) Size() int {
	return c.strategy.Size()
}

func (c *cache) initCacheByStrategy(strategy StrategyType) {
	switch strategy {
	case StrategyTimeBasedExpiration:
		c.strategy = expiration.NewCache(c.capacity, c.expiration)
	case StrategyLeastFrequentlyUsed:
		c.strategy = frequently_used.NewCache(c.capacity)
	case StrategyLeastRecentlyUsed:
		c.strategy = recently_used.NewCache(c.capacity)
	default:
		// todo
		c.logger.Println("unknown cache strategy")
	}
}

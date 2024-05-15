package semrush_cache

import (
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
	maxSize    int
	expiration time.Duration
	logger     *log.Logger
}

func New(strategy StrategyType, options ...Option) Cache {
	c := &cache{
		maxSize: math.MaxInt,
		logger:  log.Default(),
	}

	for _, opt := range options {
		opt(c)
	}

	c.initCacheByStrategy(strategy)
	return c
}

func (c *cache) Set(key string, value interface{}) {
	c.strategy.Set(key, value)
}

func (c *cache) Get(key string) (interface{}, bool) {
	return c.strategy.Get(key)
}

func (c *cache) Delete(key string) {
	c.strategy.Delete(key)
}

func (c *cache) Size() int {
	return c.strategy.Size()
}

func (c *cache) initCacheByStrategy(strategy StrategyType) {
	switch strategy {
	case StrategyTimeBasedExpiration:
		// todo
	case StrategyLeastFrequentlyUsed:
		// todo
	case StrategyLeastRecentlyUsed:
		c.strategy = recently_used.NewCache(c.maxSize)
	default:
		// todo
		c.logger.Println("unknown cache strategy")
	}
}

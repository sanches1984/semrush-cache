package semrush_cache

import "time"

// MaxSize set max cache size
func MaxSize(size int) Option {
	return func(c *cache) {
		c.maxSize = size
	}
}

// Expiration set default item expiration, by default = 0
func Expiration(duration time.Duration) Option {
	return func(c *cache) {
		c.expiration = duration
	}
}

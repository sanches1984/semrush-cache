package semrush_cache

import "time"

type CacheItem struct {
	Value      interface{}
	Expiration time.Time
}

# Semrush Cache

In-memory cache with capacity. Implements strategies:
- Least Recently Used
- Least Frequently Used
- Time-Based Expiration

## Example
```go
import (
	cache "github.com/sanches1984/semrush-cache"
	"time"
)

func main() {
    c := cache.New(StrategyTimeBasedExpiration, MaxSize(3), Expiration(time.Second))
    c.Set("key1", 1)
    value, ok := c.Get("key1")
}
```

## TODO
- implement context
- benchmark tests
- proper tests, more concurrent cases
- fix issue with expiration strategy when all events have not expired but capacity is full
- process error on unknown strategy
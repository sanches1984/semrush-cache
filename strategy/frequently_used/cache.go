package frequently_used

import (
	"container/heap"
	"github.com/sanches1984/semrush-cache/model"
	"sync"
	"time"
)

type entry struct {
	key       string
	value     model.CacheItem
	frequency int
	index     int
	timestamp int64
}

type cache struct {
	items    map[string]*entry
	pq       Queue
	mutex    sync.Mutex
	capacity int
}

func NewCache(capacity int) *cache {
	return &cache{
		items:    make(map[string]*entry),
		pq:       make(Queue, 0, capacity),
		capacity: capacity,
	}
}

func (c *cache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if item, found := c.items[key]; found {
		item.value = model.CacheItem{Value: value}
		c.increment(item)
		return
	}

	if len(c.items) >= c.capacity {
		c.removeLessUsed()
	}

	item := &entry{
		key:       key,
		value:     model.CacheItem{Value: value},
		timestamp: time.Now().UnixNano(),
	}
	heap.Push(&c.pq, item)
	c.items[key] = item
}

func (c *cache) Get(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if item, found := c.items[key]; found {
		c.increment(item)
		return item.value.Value, true
	}
	return nil, false
}

func (c *cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if item, found := c.items[key]; found {
		c.remove(item)
	}
}

func (c *cache) Size() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return len(c.items)
}

func (c *cache) increment(item *entry) {
	item.frequency++
	item.timestamp = time.Now().UnixNano()
	heap.Fix(&c.pq, item.index)
}

func (c *cache) removeLessUsed() {
	item := heap.Pop(&c.pq).(*entry)
	delete(c.items, item.key)
}

func (c *cache) remove(item *entry) {
	heap.Remove(&c.pq, item.index)
	delete(c.items, item.key)
}

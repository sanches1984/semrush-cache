package recently_used

import (
	"container/list"
	"github.com/sanches1984/semrush-cache/model"
	"sync"
)

type cache struct {
	items    map[string]*list.Element
	order    *list.List
	mutex    sync.Mutex
	capacity int
}

type entry struct {
	key   string
	value model.CacheItem
}

func NewCache(capacity int) *cache {
	return &cache{
		items:    make(map[string]*list.Element),
		order:    list.New(),
		capacity: capacity,
		mutex:    sync.Mutex{},
	}
}

func (c *cache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, found := c.items[key]; found {
		c.order.MoveToFront(elem)
		elem.Value.(*entry).value = model.CacheItem{
			Value: value,
		}
		return
	}

	if c.order.Len() >= c.capacity {
		c.deleteRecentlyUsed()
	}

	elem := c.order.PushFront(&entry{key, model.CacheItem{Value: value}})
	c.items[key] = elem
}

func (c *cache) Get(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, found := c.items[key]; found {
		item := elem.Value.(*entry).value
		c.order.MoveToFront(elem)
		return item.Value, true
	}
	return nil, false
}

func (c *cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, found := c.items[key]; found {
		c.order.Remove(elem)
		delete(c.items, key)
	}
}

func (c *cache) Size() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.order.Len()
}

func (c *cache) deleteRecentlyUsed() {
	elem := c.order.Back()
	if elem != nil {
		c.order.Remove(elem)
		entry := elem.Value.(*entry)
		delete(c.items, entry.key)
	}
}

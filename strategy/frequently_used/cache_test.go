package frequently_used

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCache(t *testing.T) {
	c := NewCache(3)
	c.Set("key1", 1)
	c.Set("key2", 2)
	c.Set("key3", 3)

	c.Get("key1")
	c.Get("key1")
	c.Get("key1")
	c.Get("key2")
	c.Get("key2")
	c.Get("key3")

	c.Set("key4", 4)

	_, ok := c.Get("key1")
	assert.True(t, ok)

	_, ok = c.Get("key2")
	assert.True(t, ok)

	_, ok = c.Get("key3")
	assert.False(t, ok)

	_, ok = c.Get("key4")
	assert.True(t, ok)
}

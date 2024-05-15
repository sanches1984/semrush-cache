package expiration

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := NewCache(5, time.Second)
	c.Set("key1", 1)
	time.Sleep(100 * time.Millisecond)
	c.Set("key2", 2)
	time.Sleep(100 * time.Millisecond)
	c.Set("key3", 3)
	time.Sleep(100 * time.Millisecond)
	c.Set("key4", 4)
	time.Sleep(100 * time.Millisecond)
	c.Set("key5", 5)
	time.Sleep(100 * time.Millisecond)

	time.Sleep(650 * time.Millisecond)

	c.Set("key6", 6)
	require.Equal(t, 4, c.Size())

	_, ok := c.Get("key1")
	assert.False(t, ok)

	_, ok = c.Get("key2")
	assert.False(t, ok)

	_, ok = c.Get("key3")
	assert.True(t, ok)

	_, ok = c.Get("key4")
	assert.True(t, ok)

	_, ok = c.Get("key5")
	assert.True(t, ok)

	_, ok = c.Get("key5")
	assert.True(t, ok)
}

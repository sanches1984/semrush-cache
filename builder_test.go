package semrush_cache

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestBuilder(t *testing.T) {
	c := New(StrategyTimeBasedExpiration, MaxSize(3), Expiration(time.Second))

	c.Set("key1", 1)
	time.Sleep(100 * time.Millisecond)
	c.Set("key2", 2)
	time.Sleep(100 * time.Millisecond)
	c.Set("key3", 3)
	time.Sleep(100 * time.Millisecond)

	time.Sleep(850 * time.Millisecond)

	c.Set("key4", 4)
	require.Equal(t, 2, c.Size())

	_, ok := c.Get("key1")
	assert.False(t, ok)

	_, ok = c.Get("key2")
	assert.False(t, ok)

	_, ok = c.Get("key3")
	assert.True(t, ok)

	_, ok = c.Get("key4")
	assert.True(t, ok)
}

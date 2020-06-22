package chapter16

import (
	"sync"
	"testing"

	"syreclabs.com/go/faker"

	"github.com/stretchr/testify/assert"
)

func TestNewLRUCache(t *testing.T) {
	assert.IsType(t, &LRUCache{}, NewLRUCache(3))
}

func TestLRUCache_SetKeyValue(t *testing.T) {
	t.Run("Setting in go routines", func(t *testing.T) {
		c := NewLRUCache(3)
		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				c.SetKeyValue(i, faker.RandomString(faker.RandomInt(0, 100)))
				wg.Done()
			}(i)
		}

		wg.Wait()

		assert.Equal(t, 3, len(c.cacheMap))
	})
}

func TestLRUCache_GetValue(t *testing.T) {
	t.Run("cache miss", func(t *testing.T) {
		c := NewLRUCache(3)

		_, err := c.GetValue(4)
		assert.EqualError(t, err, errCacheMiss.Error())
	})

	t.Run("cache get", func(t *testing.T) {
		c := NewLRUCache(3)
		val := faker.RandomString(42)

		c.SetKeyValue(50, val)
		result, err := c.GetValue(50)
		assert.NoError(t, err)

		assert.Equal(t, val, result)
	})
}

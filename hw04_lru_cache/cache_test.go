package hw04lrucache

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("clear test", func(t *testing.T) {
		c := NewCache(5)

		for i := 1; i <= 5; i++ {
			key := Key(fmt.Sprint(i))
			c.Set(key, i)
		}

		c.Clear()

		for i := 1; i <= 5; i++ {
			val, ok := c.Get(Key(fmt.Sprint(i)))
			require.False(t, ok)
			require.Nil(t, val)
		}

		for i := 101; i <= 105; i++ {
			key := Key(fmt.Sprint(i))
			c.Set(key, i)
		}

		for i := 101; i <= 105; i++ {
			val, ok := c.Get(Key(fmt.Sprint(i)))
			require.True(t, ok)
			require.Equal(t, val, i)
		}
	})

	t.Run("rollout test", func(t *testing.T) {
		c := NewCache(3)

		wasInCache := c.Set("a", 10)
		require.False(t, wasInCache)

		wasInCache = c.Set("b", 20)
		require.False(t, wasInCache)

		wasInCache = c.Set("c", 30)
		require.False(t, wasInCache)

		wasInCache = c.Set("d", 40)
		require.False(t, wasInCache)

		val, ok := c.Get("a")
		require.False(t, ok)
		require.Nil(t, val)

		wasInCache = c.Set("b", 200)
		require.True(t, wasInCache)

		wasInCache = c.Set("e", 50)
		require.False(t, wasInCache)

		val, ok = c.Get("b")
		require.True(t, ok)
		require.Equal(t, val, 200)
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}

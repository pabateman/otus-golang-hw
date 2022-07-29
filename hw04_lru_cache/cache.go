package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	locker   sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.locker.Lock()
	defer c.locker.Unlock()

	cacheItemValue := &cacheItem{
		key:   key,
		value: value,
	}

	if c.queue.Len() == c.capacity {
		if backItem, ok := c.queue.Back().Value.(*cacheItem); ok {
			delete(c.items, backItem.key)
			c.queue.Remove(c.queue.Back())
		}
	}

	if item, ok := c.items[key]; ok {
		item.Value = cacheItemValue
		c.queue.MoveToFront(item)
		c.items[key] = c.queue.Front()
		return true
	}
	c.queue.PushFront(cacheItemValue)
	c.items[key] = c.queue.Front()
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.locker.Lock()
	defer c.locker.Unlock()

	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		listItem := item.Value

		if listItemValue, ok := listItem.(*cacheItem); ok {
			return listItemValue.value, true
		}
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.locker.Lock()
	defer c.locker.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

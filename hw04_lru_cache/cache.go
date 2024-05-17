package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
	GetQueue() List
	GetItems() map[Key]*ListItem
}

type lruCache struct {
	capacity int
	Queue    List
	Items    map[Key]*ListItem
	mu       sync.Mutex
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		Queue:    NewList(),
		Items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) GetQueue() List {
	return c.Queue
}

func (c *lruCache) GetItems() map[Key]*ListItem {
	return c.Items
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if item, ok := c.Items[key]; ok {
		item.Value = value
		c.Queue.MoveToFront(item)
		return true
	}
	if c.Queue.Len() == c.capacity {
		last := c.Queue.Back()
		c.Queue.Remove(last)
	}
	item := c.Queue.PushFront(value)
	c.Items[key] = item
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if item, ok := c.Items[key]; ok {
		c.Queue.MoveToFront(item)
		return item.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Queue = NewList()
	c.Items = make(map[Key]*ListItem)
}

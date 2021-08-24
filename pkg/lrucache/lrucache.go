package lrucache

import (
	"sync"
)

type LRUCache struct {
	list    *DoublyLinkedList // For Caching sorted nodes by recent using
	items   map[string]*Node  // For selection nodes
	maxSize int
	mutex   sync.RWMutex
}

type KVPair struct {
	key   string
	value interface{}
}

func New(size int) *LRUCache {
	return &LRUCache{
		maxSize: size,
		items:   make(map[string]*Node),
		list:    NewList(),
	}
}

func (l *LRUCache) Get(key string) interface{} {
	node := l.get(key)
	if node == nil {
		return nil
	}

	defer func() {
		l.list.MoveFront(node)
	}()

	ele := node.Value.(*KVPair)

	return ele.value
}

func (l *LRUCache) Set(key string, value interface{}) interface{} {
	// Get node from map by key
	node := l.get(key)
	// If node exist
	if node != nil {
		l.mutex.Lock()
		defer l.mutex.Unlock()
		// Update the value into node
		node.Value.(*KVPair).value = value
		l.list.MoveFront(node)
		return nil
	}

	tail := new(Node)
	if l.list.Length() == l.maxSize {
		// Delete tail from cache
		tail = l.list.RemoveTail()

		// Delete element from map of selection by key
		delete(l.items, tail.Key)
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()
	node = l.list.Unshift(key, value)
	l.items[key] = node

	if tail.Value == nil {
		return nil
	}

	return tail.Value.(*KVPair).value
}

func (l *LRUCache) Invalidate(key string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	node, exists := l.items[key]
	if !exists {
		return
	}

	l.list.Remove(node)
	delete(l.items, key)

	return
}

func (l *LRUCache) get(key string) *Node {
	// Lock RW for reading
	l.mutex.RLock()
	// Unlock before end of function
	defer l.mutex.RUnlock()
	// Get from Map of selection nodes
	node, exists := l.items[key]
	if !exists {
		return nil
	}

	return node
}

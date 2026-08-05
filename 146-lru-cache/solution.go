package _146_lru_cache

import "container/list"

type LRUCache struct {
	capacity int
	list     *list.List
	items    map[int]*list.Element
}

type entry struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		items:    make(map[int]*list.Element),
	}
}

func (lru *LRUCache) Get(key int) int {
	elem, ok := lru.items[key]
	if !ok {
		return -1
	}

	lru.list.MoveToFront(elem)

	return elem.Value.(entry).val
}

func (lru *LRUCache) Put(key int, value int) {
	if elem, ok := lru.items[key]; ok {
		elem.Value = entry{key, value}
		lru.list.MoveToFront(elem)
	} else {
		lru.items[key] = lru.list.PushFront(entry{key, value})
	}

	if len(lru.items) > lru.capacity {
		back := lru.list.Back()
		if back != nil {
			delete(lru.items, back.Value.(entry).key)
			lru.list.Remove(back)
		}
	}
}

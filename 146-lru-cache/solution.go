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

func (this *LRUCache) Get(key int) int {
	elem, ok := this.items[key]
	if !ok {
		return -1
	}

	this.list.MoveToFront(elem)

	return elem.Value.(entry).val
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.items[key]; ok {
		elem.Value = entry{key, value}
		this.list.MoveToFront(elem)
	} else {
		this.items[key] = this.list.PushFront(entry{key, value})
	}

	if len(this.items) > this.capacity {
		back := this.list.Back()
		if back != nil {
			delete(this.items, back.Value.(entry).key)
			this.list.Remove(back)
		}
	}
}

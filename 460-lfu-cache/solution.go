package _460_lfu_cache

import "container/list"

type LFUCache struct {
	cap     int
	minFreq int
	keyMap  map[int]*list.Element
	freqMap map[int]*list.List
}

type entry struct {
	key, val, freq int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:     capacity,
		keyMap:  make(map[int]*list.Element),
		freqMap: make(map[int]*list.List),
	}
}

func (this *LFUCache) getList(freq int) *list.List {
	if this.freqMap[freq] == nil {
		this.freqMap[freq] = list.New()
	}

	return this.freqMap[freq]
}

func (this *LFUCache) increment(el *list.Element) {
	e := el.Value.(*entry)

	oldList := this.freqMap[e.freq]
	oldList.Remove(el)
	if oldList.Len() == 0 {
		delete(this.freqMap, e.freq)
		if this.minFreq == e.freq {
			this.minFreq++
		}
	}

	e.freq++
	newList := this.getList(e.freq)
	newEl := newList.PushFront(e)
	this.keyMap[e.key] = newEl
}

func (this *LFUCache) Get(key int) int {
	el, ok := this.keyMap[key]
	if !ok {
		return -1
	}
	this.increment(el)

	return el.Value.(*entry).val
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}

	if el, ok := this.keyMap[key]; ok {
		el.Value.(*entry).val = value
		this.increment(el)
		return
	}

	if len(this.keyMap) == this.cap {
		minList := this.freqMap[this.minFreq]

		tail := minList.Back()
		if tail != nil {
			minList.Remove(tail)
			delete(this.keyMap, tail.Value.(*entry).key)
			if minList.Len() == 0 {
				delete(this.freqMap, this.minFreq)
			}
		}
	}

	this.minFreq = 1
	e := &entry{key: key, val: value, freq: 1}
	lst := this.getList(1)
	el := lst.PushFront(e)
	this.keyMap[key] = el
}

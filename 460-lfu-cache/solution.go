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

func (lfu *LFUCache) getList(freq int) *list.List {
	if lfu.freqMap[freq] == nil {
		lfu.freqMap[freq] = list.New()
	}

	return lfu.freqMap[freq]
}

func (lfu *LFUCache) increment(el *list.Element) {
	e := el.Value.(*entry)

	oldList := lfu.freqMap[e.freq]
	oldList.Remove(el)
	if oldList.Len() == 0 {
		delete(lfu.freqMap, e.freq)
		if lfu.minFreq == e.freq {
			lfu.minFreq++
		}
	}

	e.freq++
	newList := lfu.getList(e.freq)
	newEl := newList.PushFront(e)
	lfu.keyMap[e.key] = newEl
}

func (lfu *LFUCache) Get(key int) int {
	el, ok := lfu.keyMap[key]
	if !ok {
		return -1
	}
	lfu.increment(el)

	return el.Value.(*entry).val
}

func (lfu *LFUCache) Put(key int, value int) {
	if lfu.cap == 0 {
		return
	}

	if el, ok := lfu.keyMap[key]; ok {
		el.Value.(*entry).val = value
		lfu.increment(el)
		return
	}

	if len(lfu.keyMap) == lfu.cap {
		minList := lfu.freqMap[lfu.minFreq]

		tail := minList.Back()
		if tail != nil {
			minList.Remove(tail)
			delete(lfu.keyMap, tail.Value.(*entry).key)
			if minList.Len() == 0 {
				delete(lfu.freqMap, lfu.minFreq)
			}
		}
	}

	lfu.minFreq = 1
	e := &entry{key: key, val: value, freq: 1}
	lst := lfu.getList(1)
	el := lst.PushFront(e)
	lfu.keyMap[key] = el
}

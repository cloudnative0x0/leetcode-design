# 460 · LFU Cache

**Difficulty:** Hard | **Time:** O(1) | **Space:** O(capacity)

---

## Solution uses frequency buckets and hash tables
Two maps and a set of doubly-linked lists.
One map gives O(1) access by key, the other groups nodes by frequency. Inside each frequency – LRU order.


---

## Architecture

Let's think in layers:
- keyMap[key]*list.Element – is direct access to the node;
- freqMap[freq]*list.List – list of all elements with the same frequency
> LFU = Least Frequently Used
> on a tie frequency, we evict the LRU within the group

```go
type entry struct {
    key, val, freq int
}

type LFUCache struct {
    cap     int
    minFreq int
    keyMap  map[int]*list.Element
    freqMap map[int]*list.List
}
```

## Core invariant
Each item has a frequency (freq). Items are grouped into frequency buckets. In case of a tie, the LRU item within the bucket is evicted.
```minFreq``` always tracks the minimum frequency currently in the cache. Within each frequency group, elements are ordered by their recency of use (LRU). 
Within each frequency group, elements are ordered by their recency of use (LRU).


## increment (heart of LFU)
Any operation: it can be ```Get``` or ```Put``` – increase a frequency
Steps:
- Get an entry point
- Remove it from old frequency list
- if the list become empty – delete it 
- Update ```minFreq```
- Increment ```freq```
- Put it to the new list (to the head)

```go
func (lru *LFUCache) increment(el *list.Element) {
    e := el.Value.(*entry)

    oldList := lru.freqMap[e.freq]
    oldList.Remove(el)

    if oldList.Len() == 0 {
        delete(lru.freqMap, e.freq)
        if lru.minFreq == e.freq {
            lru.minFreq++
        }
    }

    e.freq++
    newList := lru.getList(e.freq)
    newEl := newList.PushFront(e)

    lru.keyMap[e.key] = newEl
}
```

Важно: создаётся новый list.Element, поэтому keyMap нужно обновить

## Get

Стандартный паттерн:
нашли элемент → O(1)
увеличили частоту → increment
вернули значение

```go
func (lru *LFUCache) Get(key int) int {
    el, ok := lru.keyMap[key]
    if !ok {
        return -1
    }

    lru.increment(el)
    return el.Value.(*entry).val
}
```

## Put
We have three scenarios:
1. capacity = 0 – nothing to do
2. we have a key, so we can update the value and increment a frequency
3. new key: if our cache is full – we can get ```minFreq``` and go further in current list, delete a tail. 
Next: create a new entry with ```freq = 1```, put it in ```freq = 1``` and update ```minFreq = 1```

```go
func (lru *LFUCache) Put(key int, value int) {
    if lru.cap == 0 {
        return
    }

    if el, ok := lru.keyMap[key]; ok {
        el.Value.(*entry).val = value
        lru.increment(el)
        return
    }

    if len(lru.keyMap) == lru.cap {
        minList := lru.freqMap[lru.minFreq]
        tail := minList.Back()

        if tail != nil {
            minList.Remove(tail)
            delete(lru.keyMap, tail.Value.(*entry).key)

            if minList.Len() == 0 {
                delete(lru.freqMap, lru.minFreq)
            }
        }
    }

    lru.minFreq = 1
    e := &entry{key: key, val: value, freq: 1}

    lst := lru.getList(1)
    el := lst.PushFront(e)

    lru.keyMap[key] = el
}
```

## getList helper

Lazy initialization of the list by frequency.

```go
func (lru *LFUCache) getList(freq int) *list.List {
    if lru.freqMap[freq] == nil {
        lru.freqMap[freq] = list.New()
    }
    return lru.freqMap[freq]
}
```

## stdlib notes
- ```container/list``` — double linked list with sentinel nodes
- ```list.Element``` — keep the Value any + pointers
- ```PushFront``` — O(1), returns *Element
– ```Remove``` — O(1), pointer ops
– ```Back()``` — access to LRU within the frequency

--- 

> LFU without frequency buckets turns to linear hell.
> The whole thing hinges on the fact that every operation involves only pointer manipulation.

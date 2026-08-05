package _706_hashmap

const bucketCount = 10007

type Node struct {
	key   int
	value int
}

type MyHashMap struct {
	buckets [][]*Node
}

func Constructor() MyHashMap {
	return MyHashMap{buckets: make([][]*Node, bucketCount)}
}

func (hm *MyHashMap) hash(key int) int {
	return key % bucketCount
}

func (hm *MyHashMap) Put(key int, value int) {
	idx := hm.hash(key)
	for _, n := range hm.buckets[idx] {
		if n.key == key {
			n.value = value
			return
		}
	}

	hm.buckets[idx] = append(hm.buckets[idx], &Node{key: key, value: value})
}

func (hm *MyHashMap) Get(key int) int {
	idx := hm.hash(key)
	for _, n := range hm.buckets[idx] {
		if n.key == key {
			return n.value
		}
	}

	return -1
}

func (hm *MyHashMap) Remove(key int) {
	idx := hm.hash(key)
	for i, n := range hm.buckets[idx] {
		if n.key == key {
			hm.buckets[idx] = append(hm.buckets[idx][:i], hm.buckets[idx][i+1:]...)
			return
		}
	}
}

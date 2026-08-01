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

func (this *MyHashMap) hash(key int) int {
	return key % bucketCount
}

func (this *MyHashMap) Put(key int, value int) {
	idx := this.hash(key)
	for _, n := range this.buckets[idx] {
		if n.key == key {
			n.value = value
			return
		}
	}

	this.buckets[idx] = append(this.buckets[idx], &Node{key: key, value: value})
}

func (this *MyHashMap) Get(key int) int {
	idx := this.hash(key)
	for _, n := range this.buckets[idx] {
		if n.key == key {
			return n.value
		}
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	idx := this.hash(key)
	for i, n := range this.buckets[idx] {
		if n.key == key {
			this.buckets[idx] = append(this.buckets[idx][:i], this.buckets[idx][i+1:]...)
			return
		}
	}
}

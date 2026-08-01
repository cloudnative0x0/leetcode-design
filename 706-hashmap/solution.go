package _706_hashmap

const size = 769

type node struct {
	key  int
	val  int
	next *node
}

type MyHashMap struct {
	buckets [size]*node
}

func Constructor() MyHashMap {
	return MyHashMap{}
}

func (this *MyHashMap) Put(key int, value int) {
	idx := key % size

	curr := this.buckets[idx]
	for curr != nil {
		if curr.key == key {
			curr.val = value
			return
		}

		curr = curr.next
	}

	this.buckets[idx] = &node{
		key:  key,
		val:  value,
		next: this.buckets[idx],
	}
}

func (this *MyHashMap) Get(key int) int {
	idx := key % size
	curr := this.buckets[idx]

	for curr != nil {
		if curr.key == key {
			return curr.val
		}

		curr = curr.next
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	idx := key % size
	curr := this.buckets[idx]

	if curr == nil {
		return
	}

	if curr.key == key {
		this.buckets[idx] = curr.next
		return
	}

	for curr.next != nil {
		if curr.next.key == key {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}
}

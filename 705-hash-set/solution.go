package _705_hash_set

type Node struct {
	key  int
	next *Node
}

type MyHashSet struct {
	buckets     []*Node
	bucketCount int
}

func Constructor() MyHashSet {
	const initialCapacity = 10007

	return MyHashSet{
		buckets:     make([]*Node, initialCapacity),
		bucketCount: initialCapacity,
	}
}

func (mhs *MyHashSet) hash(key int) int {
	h := key % mhs.bucketCount
	if h < 0 {
		h += mhs.bucketCount
	}

	return h
}

func (mhs *MyHashSet) Add(key int) {
	index := mhs.hash(key)

	if mhs.buckets[index] == nil {
		mhs.buckets[index] = &Node{key: key}
		return
	}

	current := mhs.buckets[index]
	for {
		if current.key == key {
			return
		}
		if current.next == nil {
			break
		}

		current = current.next
	}

	current.next = &Node{key: key}
}

func (mhs *MyHashSet) Remove(key int) {
	index := mhs.hash(key)
	current := mhs.buckets[index]

	if current == nil {
		return
	}

	if current.key == key {
		mhs.buckets[index] = current.next
		return
	}

	for current.next != nil {
		if current.next.key == key {
			current.next = current.next.next
			return
		}

		current = current.next
	}
}

func (mhs *MyHashSet) Contains(key int) bool {
	index := mhs.hash(key)
	current := mhs.buckets[index]

	for current != nil {
		if current.key == key {
			return true
		}

		current = current.next
	}

	return false
}

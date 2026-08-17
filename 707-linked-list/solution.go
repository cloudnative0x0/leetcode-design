package _707_linked_list

type Node struct {
	Value int
	Next  *Node
}

type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (mll *MyLinkedList) Get(index int) int {
	if index < 0 || index >= mll.size {
		return -1
	}

	cur := mll.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	return cur.Value
}

func (mll *MyLinkedList) AddAtHead(val int) {
	node := &Node{Value: val, Next: mll.head}
	mll.head = node

	if mll.size == 0 {
		mll.tail = node
	}

	mll.size++
}

func (mll *MyLinkedList) AddAtTail(val int) {
	node := &Node{Value: val}

	if mll.size == 0 {
		mll.head = node
		mll.tail = node
	} else {
		mll.tail.Next = node
		mll.tail = node
	}

	mll.size++
}

func (mll *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > mll.size {
		return
	}

	if index == 0 {
		mll.AddAtHead(val)
		return
	}

	if index == mll.size {
		mll.AddAtTail(val)
		return
	}

	prev := mll.head
	for i := 0; i < index-1; i++ {
		prev = prev.Next
	}

	node := &Node{Value: val, Next: prev.Next}
	prev.Next = node

	mll.size++
}

func (mll *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= mll.size {
		return
	}

	if index == 0 {
		mll.head = mll.head.Next
		if mll.head == nil {
			mll.tail = nil
		}

		mll.size--
		return
	}

	prev := mll.head
	for i := 0; i < index-1; i++ {
		prev = prev.Next
	}

	if prev.Next == mll.tail {
		mll.tail = prev
	}

	prev.Next = prev.Next.Next
	mll.size--
}

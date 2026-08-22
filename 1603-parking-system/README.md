# 707 · Design Linked List

**Difficulty:** Medium | **Time:** O(1) for head/tail ops, O(n) for index ops | **Space:** O(n)

---

## Solution approach

Implement a linked list supporting six operations: get value by index, add at head, add at tail, add at index, delete at index. This is a classic problem of pointer manipulation.

The key choice is to use a singly linked list because all index‑based operations traverse the list forward only. A doubly linked list offers no advantage but increases memory usage and pointer maintenance complexity.

In addition to `head`, we store `tail` and `size`. They are not required for correctness but are critical for performance:

- `tail` makes `AddAtTail` O(1) instead of O(n).
- `size` gives O(1) bound checks in `Get`, `AddAtIndex`, `DeleteAtIndex` instead of O(n) traversal to count length.

All mutations update `size`. When the head or tail changes, `tail` is adjusted in special cases (deleting the only node or deleting the last node).

---

## Structure description

```go
type Node struct {
    Value int
    Next  *Node
}

type MyLinkedList struct {
    head *Node
    tail *Node
    size int
}
```

- `Node` – list element with a value and a pointer to the next node.
- `MyLinkedList` – stores head, tail, and current size. The zero value of the struct is valid (all fields are `nil` and `0`).

---

## Constructor

```go
func Constructor() MyLinkedList {
    return MyLinkedList{}
}
```

No initialisation is needed – an empty list already has `head = nil`, `tail = nil`, `size = 0`.

---

## Get(index int) int

```go
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
```

First checks index validity via `size`. Then walks from head to the target node. Returns the value or `-1` if the index is invalid.

---

## AddAtHead(val int)

```go
func (mll *MyLinkedList) AddAtHead(val int) {
    node := &Node{Value: val, Next: mll.head}
    mll.head = node
    if mll.size == 0 {
        mll.tail = node
    }
    mll.size++
}
```

The new node points to the old head. If the list was empty, this same node becomes the tail as well. Size increments.

---

## AddAtTail(val int)

```go
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
```

For an empty list, both head and tail point to the new node. Otherwise, attach the new node to the current tail and move the tail forward. Size increments.

---

## AddAtIndex(index int, val int)

```go
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
```

Index can be equal to `size` – that means appending at the end. Both boundary cases are delegated to `AddAtHead` and `AddAtTail` to avoid duplicating `tail`‑update logic. In the general case, find the node before the insertion point, insert the new node between `prev` and `prev.Next`. Size increments.

---

## DeleteAtIndex(index int)

```go
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
```

Special attention to `tail`. There are two cases where `tail` can become invalid:

1. Deleting the only node (index 0 and after shift `head == nil`) – then `tail` must be set to `nil`.
2. Deleting the last node (check that `prev.Next` points to `tail`) – then the new tail becomes `prev`.

In all other cases, simply skip the removed node by `prev.Next = prev.Next.Next`. Size decrements.

---

## Summary

All methods maintain the invariant: `tail` always points to the last reachable node, and `size` always equals the list length. Thanks to this, `AddAtTail` and bound checks run in O(1), which is crucial for frequent calls.
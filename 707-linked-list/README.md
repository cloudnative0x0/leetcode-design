# 707 · Design Linked List

**Difficulty:** Medium | **Time:** O(1) head/tail, O(n) index ops | **Space:** O(n)

---

## Solution uses a singly linked list with a tail pointer

One direction of `Next` pointers, plus a `tail` field and a `size` counter kept alongside `head`. No sentinel/dummy node, no `Prev` pointers.

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

## Why not doubly linked

A doubly linked list buys nothing here. The interface only ever walks forward — `Get`, `AddAtIndex`, `DeleteAtIndex` all count up from `head`. Nothing in the problem asks to go backward from a given node. Adding `Prev` would mean an extra pointer per node and an extra pointer to fix on every insert/delete, for zero benefit.

## Why keep a tail pointer at all

Without it, `AddAtTail` is O(n): walk from `head` to the last node every single call. With `tail` cached, appending is O(1) — set `mll.tail.Next`, move `mll.tail` forward. The list is doing `Add`/`Delete` calls on the same object repeatedly, so paying for the pointer up front is cheaper than re-walking the chain on every tail insert.

## Why keep a size counter

`Get`, `AddAtIndex`, `DeleteAtIndex` all need to know the current length to validate `index`. Without `size`, that means a full traversal just to count nodes before the real work starts — doubling the cost of every call. `size` is updated in place on every mutation, so the length check is O(1).

## Constructor

```go
func Constructor() MyLinkedList {
    return MyLinkedList{}
}
```

Zero value already works: `head`, `tail` are `nil`, `size` is `0`. No initialization logic needed.

## Get

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

Bounds check first, using `size` instead of walking to find out the list ran out. Then a plain forward walk — singly linked means there's no shortcut, O(n) is the floor for arbitrary index.

## AddAtHead

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

The one edge case: inserting into an empty list means the new node is both head and tail. That `if mll.size == 0` branch is the only place `tail` needs manual attention here — every other head insert leaves `tail` untouched, which is correct since the old tail is still the last node.

## AddAtTail

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

Mirror of `AddAtHead`: empty list is a separate branch because there's no `mll.tail.Next` to write into yet — `mll.tail` is `nil`, and dereferencing it would panic. Non-empty case is the O(1) append the whole `tail` field exists for.

## AddAtIndex

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

`index == size` is valid on purpose — the problem defines it as append. Rather than duplicate the empty-list and tail-update logic, both edge cases just delegate to `AddAtHead`/`AddAtTail`, which already handle the `tail`-pointer bookkeeping correctly. That leaves the general branch to only worry about the middle case: find the node *before* the target index, splice in.

## DeleteAtIndex

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

Two places `tail` can go stale, both handled explicitly:

- **Deleting the only node** (`index == 0`, `mll.head.Next == nil`): `head` becomes `nil`, and `tail` must follow — otherwise `tail` keeps pointing at a node that's no longer reachable from `head`, and the next `AddAtTail` would silently corrupt the list by writing into a detached node.
- **Deleting the last node** (`prev.Next == mll.tail`): `tail` has to move back to `prev` *before* the unlink, since after `prev.Next = prev.Next.Next` there's no way to reach the old predecessor from the removed node anymore.

Everything else is a standard "walk to `index-1`, skip one node" unlink.

---

> A tail pointer only pays for itself if every mutation that could invalidate it is caught — the two branches in `DeleteAtIndex` exist because deleting the head or the tail are the only operations that can leave `tail` pointing at a node no longer reachable from `head`.

`size` and `tail` are both derivable from the chain by walking it, so neither is strictly required for correctness — only for avoiding an O(n) traversal on every call that needs them.
# 706 · Design HashMap

**Difficulty:** Easy | **Time:** O(1) amortized | **Space:** O(n)

---

## Solution uses separate chaining

Fixed-size array of buckets, each bucket is the head of a linked list.
Collisions are resolved through chaining: if two keys land on the same index, they just get chained to each other via `next`.

---

## Architecture

- `buckets [size]*node` — array of list heads, size is fixed at compile time
- `node` — list element with `key`, `val` and a pointer to the next node

```go
const size = 769

type node struct {
    key  int
    val  int
    next *node
}

type MyHashMap struct {
    buckets [size]*node
}
```

> 769 is a prime number. Primes as a table size cut down on collisions for unlucky key sets (when keys share common divisors).

## Hash function

Nothing fancy — just remainder from division:

```go
idx := key % size
```

No extra bit mixing, no XOR with a shift. For LeetCode's integer keys this is enough, distribution comes out even.

## Put

First walk the chain — if the key is already there, update the value and return:

```go
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
```

If the loop reaches the end without a match — insert the new node at the head. Head insertion is O(1), no need to walk the list a second time to append.

## Get

Linear walk through the chain, buckets stay short in practice — close to O(1):

```go
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
```

Not found — return -1, as the problem requires.

## Remove

Two cases here, easy to mix up:

1. The node to remove is the head of the list. Just move `buckets[idx]` to `next`.
2. The node is somewhere in the middle or at the tail. Walk with a pointer to the previous element and relink `next`.

```go
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
```

The head can't be removed in the same loop as the rest of the nodes — it has no explicit "previous", so it's handled as a separate branch before the loop starts.

## Why an array instead of map[int]int

The problem statement asks to build a HashMap without using built-in hash tables. Hence a fixed-size array plus chaining — the minimal setup that implements the actual idea of a hash table by hand.

## Complexity

| Operation | Average | Worst case |
|-----------|---------|------------|
| Put       | O(1)    | O(n)       |
| Get       | O(1)    | O(n)       |
| Remove    | O(1)    | O(n)       |

Worst case — all keys collide into one bucket (key % size matches for every key). On random data with size = 769 this barely happens.

---

> Chaining is simpler than open addressing but pays for it with an extra allocation per node.
> For an interview or LeetCode it's enough — no resize, no load factor to worry about.
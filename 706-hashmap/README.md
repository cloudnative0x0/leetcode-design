# 706 · Design HashMap

**Difficulty:** Easy | **Time:** O(1) amortized | **Space:** O(n + buckets)

---

## Solution uses separate chaining
An array of buckets, each holding a slice of entries that share the same hash. No resizing, no open addressing. Fixed prime-sized table plus a chain per bucket for collisions.

---

## Why not open addressing

The first version used open addressing with linear probing on a fixed table of size 16. It timed out, for two reasons:

- **16 buckets is nothing.** LeetCode runs roughly 10^4 Put calls. Past the 16th insert every slot is occupied, and `for this.buckets[index] != nil` either loops forever or degrades into a full array scan on every single call.
- **Remove never freed the slot.** It did `this.buckets[index].key = -1` instead of an actual deletion. The slot stays non-nil, so Put still treats it as occupied and can never reuse it. The table clogs permanently — Remove calls pile up without ever giving Put any real room back.

A fixed-size table with no resize plus linear probing is a guaranteed dead end once enough put/remove calls accumulate. Combining open addressing with real deletion also needs tombstones and periodic rehashing to stay correct — extra machinery for a problem whose constraints are just `0 <= key, value <= 10^6`.

Chaining sidesteps all of it: buckets are independent, deletion is a plain slice removal, no half-dead state to track.

```go
const bucketCount = 10007

type Node struct {
    key   int
    value int
}

type MyHashMap struct {
    buckets [][]*Node
}
```

## Why 10007

A prime, comfortably above the 10^4 operations in the constraints. Prime bucket counts spread out keys better than round numbers when the input has patterns (powers of two, multiples of ten, etc.) — a standard hash table trick to avoid clustering. No resize logic needed: with 10007 buckets and at most 10^4 entries, average chain length stays near 0 or 1.

## hash

```go
func (this *MyHashMap) hash(key int) int {
    return key % bucketCount
}
```

Constraints guarantee `key >= 0`, so the negative-key branch from the old version is gone.

## Put

Walk the bucket's chain. Found the key — overwrite the value and return. Not found — append a new node.

```go
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
```

## Get

Linear scan through the chain, O(1) on average since chains stay short.

```go
func (this *MyHashMap) Get(key int) int {
    idx := this.hash(key)
    for _, n := range this.buckets[idx] {
        if n.key == key {
            return n.value
        }
    }
    return -1
}
```

## Remove

This is where the fix actually lives: the entry gets removed for real, not marked.

```go
func (this *MyHashMap) Remove(key int) {
    idx := this.hash(key)
    for i, n := range this.buckets[idx] {
        if n.key == key {
            this.buckets[idx] = append(this.buckets[idx][:i], this.buckets[idx][i+1:]...)
            return
        }
    }
}
```

`append(slice[:i], slice[i+1:]...)` is the usual Go idiom for cutting an element out of a slice — shifts the tail left, length drops by one. The backing array doesn't shrink, but that's irrelevant here: live entry count is tracked by slice length, not capacity.

---

> Open addressing without resizing is a time bomb: once load factor approaches 100%, every operation collapses into O(n) or an infinite loop.
> Chaining over a fixed prime-sized table clears the constraints without a single line of resize logic.

If the key range weren't bounded to `[0, 10^6]`, this would need dynamic resizing on top (double bucketCount and rehash once load factor passes ~0.75), but for this problem that's unnecessary weight.

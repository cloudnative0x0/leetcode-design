# 1656 · Ordered Stream

**Difficulty:** Easy | **Time:** O(1) amortized per insert | **Space:** O(n)

---

## Solution uses a fixed‑size slice with a pointer

Store values directly by `idKey` in a pre‑allocated slice. A `ptr` tracks the smallest `idKey` not yet returned. On each insert, fill the slot, then advance `ptr` through any contiguous filled entries, returning them as a chunk.

```go
type OrderedStream struct {
    stream []string
    ptr    int
}
```

## Why an array (slice) instead of a map or linked list

- The `idKey` domain is known in advance (`1..n`), so direct indexing gives O(1) write.
- A map would also be O(1) but adds overhead and doesn’t help with ordered scanning.
- A linked list would require traversal to find the next missing id, making each insert O(n) in the worst case.

The slice + pointer combination keeps the total work linear: each position is visited exactly once by the pointer across all inserts.

## Why keep a pointer (`ptr`)

Without it, each `Insert` would have to start from `1` and scan forward repeatedly, turning the whole process into O(n²). The pointer remembers where the last gap was, so the scan for the next chunk continues from where it left off.

## Why allocate `n+2` instead of `n+1`

- The loop condition is `odr.ptr < len(odr.stream)`.
- When `ptr` reaches `n+1`, the condition fails safely.
- Allocating `n+2` gives a valid index `n+1` (never used) so that after inserting at `idKey == n` and advancing `ptr` to `n+1`, the next loop iteration sees `ptr` out of bounds and stops – no extra bound check is needed.
- The extra slot also prevents index‑out‑of‑range when the internal loop tries to read `stream[n+1]` after `ptr` becomes `n+1` (though it wouldn’t happen because `ptr` is compared first).

## Constructor

```go
func Constructor(n int) OrderedStream {
    return OrderedStream{
        stream: make([]string, n+2),
        ptr:    1,
    }
}
```

- Pre‑allocates the slice with size `n+2`.
- Initialises `ptr` to `1` (the first valid key).

## Insert

```go
func (odr *OrderedStream) Insert(idKey int, value string) []string {
    odr.stream[idKey] = value

    var chunk []string
    for odr.ptr < len(odr.stream) && odr.stream[odr.ptr] != "" {
        chunk = append(chunk, odr.stream[odr.ptr])
        odr.ptr++
    }
    return chunk
}
```

- Writes the value at the given index.
- Starting from `ptr`, it walks forward while the slot is filled (non‑empty string), appending each value.
- The pointer advances past every returned element.
- Returns the collected chunk (may be empty if the next required id is still missing).

**Edge cases:**
- Inserting out‑of‑order (e.g., `idKey=3` before `1`) stores the value but returns empty because `ptr` is still `1`.
- Once all gaps before `ptr` are filled, the loop returns the maximal contiguous block.
- After the last element is inserted, `ptr` advances to `n+1` and subsequent inserts return empty (the slice still holds old values but `ptr` is beyond the valid range).
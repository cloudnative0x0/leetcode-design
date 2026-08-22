# 933 · Number of Recent Calls

**Difficulty:** Easy | **Time:** O(1) amortized per ping | **Space:** O(n)

---

## Solution keeps a sliding window in a slice

Incoming `t` values are strictly increasing, so it's enough to store them in a slice in arrival order and drop everything from the front that falls outside the window `[t-3000, t]`. The length of what's left is the answer.

```go
type RecentCounter struct {
    calls []int
}
```

## Why a slice instead of a linked-list queue or a ring buffer

- Data arrives in strictly non-decreasing order of `t`, so removal only ever happens at the front and insertion only ever happens at the back — exactly what a plain slice is good for.
- A linked list would give the same behavior but with extra allocations per node.
- A ring buffer only makes sense with a known upper bound on the window size, which the problem doesn't guarantee.

## Constructor

```go
func Constructor() RecentCounter {
    return RecentCounter{
        calls: make([]int, 0),
    }
}
```

Creates an empty slice to accumulate call timestamps.

## Ping

```go
func (rc *RecentCounter) Ping(t int) int {
    rc.calls = append(rc.calls, t)

    for rc.calls[0] < t-3000 {
        rc.calls = rc.calls[1:]
    }

    return len(rc.calls)
}
```

- The new `t` is appended to the end — order is preserved automatically since input values never decrease.
- The loop moves the start of the slice forward while the oldest element is older than the `t-3000` cutoff.
- `rc.calls[1:]` doesn't copy data, it just moves the pointer over the underlying array — the operation itself is cheap.
- Returns the current slice length — the count of calls within the last 3000 ms, inclusive.

**Edge cases:**
- The very first `Ping` call can never trigger `rc.calls[0] < t-3000`, since the only element is `t` itself, so the loop doesn't run.
- Trimming the slice from the left without reallocating the underlying array gradually leaves unused memory at the front — worth keeping in mind for long-running use, though not an issue for this problem's constraints.
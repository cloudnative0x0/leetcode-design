# 232 · Implement Queue using Stacks

**Difficulty:** Easy | **Time:** amortized O(1) | **Space:** O(n)

---

## Solution uses two stacks: one for input, one for output
Push always goes to the input stack.
Pop and Peek drain the input stack into the output stack when the output is empty, then operate on top of the output stack.

---

## Architecture

Two slices are used as stacks:

- `enStack` – collects pushed elements.
- `deStack` – holds elements in reversed order, ready for pop/peek.

```go
type MyQueue struct {
    enStack []int
    deStack []int
}

func Constructor() MyQueue {
    return MyQueue{
        enStack: make([]int, 0),
        deStack: make([]int, 0),
    }
}
```

## Core invariant
At any moment the logical front of the queue is the top of `deStack` (if `deStack` is not empty), otherwise it is the bottom of `enStack`.
The `move()` helper enforces that `deStack` is never empty when we need to pop or peek – it flips `enStack` onto `deStack` in one batch.

The order is preserved: pushing onto `enStack` adds to the rear; when moved to `deStack`, the first pushed becomes the top of `deStack`, ready to be popped.

## move – heart of the queue
Called before every Pop and Peek. If `deStack` is empty, every element from `enStack` is popped and pushed onto `deStack`. This reverses the order, placing the oldest element on top.

```go
func (lru *MyQueue) move() {
    if len(lru.deStack) == 0 {
        for len(lru.enStack) > 0 {
            top := len(lru.enStack) - 1
            lru.deStack = append(lru.deStack, lru.enStack[top])
            lru.enStack = lru.enStack[:top]
        }
    }
}
```

After `move()`, `deStack` contains all current elements in FIFO order (front on top).

## Push
Simply append to `enStack`.

```go
func (lru *MyQueue) Push(x int) {
    lru.enStack = append(lru.enStack, x)
}
```

## Pop
1. Call `move()` to ensure the front is accessible.
2. Take the top of `deStack`.
3. Shrink `deStack` and return the value.

```go
func (lru *MyQueue) Pop() int {
    lru.move()
    topIdx := len(lru.deStack) - 1
    val := lru.deStack[topIdx]
    lru.deStack = lru.deStack[:topIdx]
    return val
}
```

## Peek
Like Pop, but without removing the element.

```go
func (lru *MyQueue) Peek() int {
    lru.move()
    return lru.deStack[len(lru.deStack)-1]
}
```

## Empty
The queue is empty when both stacks are empty.

```go
func (lru *MyQueue) Empty() bool {
    return len(lru.enStack) == 0 && len(lru.deStack) == 0
}
```

## Complexity
- Every element is pushed to `enStack` once and popped from `enStack` once (during a `move`).
- Every element is pushed to `deStack` once and popped from `deStack` once (during a Pop/Peek).
- Thus each operation is amortized O(1). Worst‑case for a single Pop/Peek is O(n) when `move` is triggered, but that happens rarely.

---

> Using two slices with a lazy transfer keeps the code simple and avoids any linked‑list overhead.  
> The key is that `move` is called exactly when `deStack` is exhausted – never earlier.
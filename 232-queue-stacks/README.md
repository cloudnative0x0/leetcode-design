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
func (this *MyQueue) move() {
    if len(this.deStack) == 0 {
        for len(this.enStack) > 0 {
            top := len(this.enStack) - 1
            this.deStack = append(this.deStack, this.enStack[top])
            this.enStack = this.enStack[:top]
        }
    }
}
```

After `move()`, `deStack` contains all current elements in FIFO order (front on top).

## Push
Simply append to `enStack`.

```go
func (this *MyQueue) Push(x int) {
    this.enStack = append(this.enStack, x)
}
```

## Pop
1. Call `move()` to ensure the front is accessible.
2. Take the top of `deStack`.
3. Shrink `deStack` and return the value.

```go
func (this *MyQueue) Pop() int {
    this.move()
    topIdx := len(this.deStack) - 1
    val := this.deStack[topIdx]
    this.deStack = this.deStack[:topIdx]
    return val
}
```

## Peek
Like Pop, but without removing the element.

```go
func (this *MyQueue) Peek() int {
    this.move()
    return this.deStack[len(this.deStack)-1]
}
```

## Empty
The queue is empty when both stacks are empty.

```go
func (this *MyQueue) Empty() bool {
    return len(this.enStack) == 0 && len(this.deStack) == 0
}
```

## Complexity
- Every element is pushed to `enStack` once and popped from `enStack` once (during a `move`).
- Every element is pushed to `deStack` once and popped from `deStack` once (during a Pop/Peek).
- Thus each operation is amortized O(1). Worst‑case for a single Pop/Peek is O(n) when `move` is triggered, but that happens rarely.

---

> Using two slices with a lazy transfer keeps the code simple and avoids any linked‑list overhead.  
> The key is that `move` is called exactly when `deStack` is exhausted – never earlier.
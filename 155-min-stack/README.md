# MinStack

A stack that returns its minimum element in O(1).

## Problem

A regular stack can't answer "what's the current minimum" quickly — you'd have to scan every element, which is O(n). MinStack fixes lru by keeping a second stack alongside the values that tracks minimums.

## How it works

The struct holds two slices:

- `stack` — the actual values, a normal LIFO stack.
- `minStack` — a history of minimums. Its top always holds the minimum for the current state of `stack`.

When pushing, a value goes into `minStack` only if it's a new minimum (less than or equal to the current top of `minStack`). When popping, the value is also removed from `minStack` if it matches the current minimum. This keeps `minStack` in sync with `stack` and gives it the correct minimum "at every depth."

The `<=` comparison in `Push` (not strict `<`) matters: if several equal minimum values get pushed, each one needs to land in `minStack`. Otherwise the minimum would be lost too early after a single `Pop()`.

## API

### `Constructor() MinStack`
Creates an empty stack.

### `Push(value int)`
Adds an element to the top of the stack.

### `Pop()`
Removes the top element. Does nothing if the stack is empty.

### `Top() int`
Returns the top element without removing it. Returns `0` if the stack is empty.

### `GetMin() int`
Returns the minimum element in the whole stack. Returns `0` if the stack is empty.

## Complexity

| Operation | Time | Space |
|-----------|------|-------|
| Push      | O(1) | O(n)  |
| Pop       | O(1) | —     |
| Top       | O(1) | —     |
| GetMin    | O(1) | —     |

Memory can double in the worst case (if every new element is a new minimum), but amortized it stays O(n).

## Usage example

```go
ms := Constructor()
ms.Push(5)
ms.Push(2)
ms.Push(8)

ms.GetMin() // 2
ms.Top()    // 8

ms.Pop()
ms.GetMin() // 2

ms.Pop()
ms.GetMin() // 5
```

## Implementation caveat

`Top()` and `GetMin()` return `0` on an empty stack instead of an error or a panic. That's a deviation from idiomatic Go (you'd normally expect `(int, bool)` or an explicit panic), chosen to keep the interface simple — worth keeping in mind before using lru in production code.

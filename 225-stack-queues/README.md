# 225 · Implement Stack using Queues

**Difficulty:** Easy | **Time:** Push O(n), Pop O(1) | **Space:** O(n)

---

## Solution uses a single queue and rotates it after every push
We maintain the invariant that the front of the queue is always the top of the stack.
After pushing an element to the back, we move all previous elements behind it one by one – this places the newest element at the front.

---

## Architecture

One dynamic slice acts as the queue:

```go
type MyStack struct {
    queue []int
}

func Constructor() MyStack {
    return MyStack{
        queue: make([]int, 0),
    }
}
```

## Core invariant
The element at index `0` of `queue` is the top of the stack. All other elements follow in LIFO order – the oldest pushed element is at the very end.

This invariant is restored inside `Push`.

## Push (the heart of the stack)

Add the new element to the back, then rotate the queue exactly `len(queue)-1` times.  
Each rotation takes the front element and moves it to the back. This brings the newly added element to the front while preserving the relative LIFO order of the rest.

```go
func (ms *MyStack) Push(x int) {
    ms.queue = append(ms.queue, x)

    for i := 0; i < len(ms.queue)-1; i++ {
        head := ms.queue[0]
        ms.queue = ms.queue[1:]
        ms.queue = append(ms.queue, head)
    }
}
```

*Пример:* очередь `[2,1]`, Push(3) → `[2,1,3]`, два поворота → `[3,2,1]`. Вершина – 3.

## Pop
The top is always the first element. Just remove and return it.

```go
func (ms *MyStack) Pop() int {
    head := ms.queue[0]
    ms.queue = ms.queue[1:]
    return head
}
```

## Top
Returns the first element without removing it.

```go
func (ms *MyStack) Top() int {
    return ms.queue[0]
}
```

## Empty
True when the queue has no elements.

```go
func (ms *MyStack) Empty() bool {
    return len(ms.queue) == 0
}
```

## Complexity
- `Push` – O(n) because of the rotation loop; every existing element is moved once per push.
- `Pop` and `Top` – O(1), simple slice index access.
- `Empty` – O(1).
- Space – O(n) for the single slice.

> Using one queue with post‑push rotation keeps the code minimal.  
> The invariant that index `0` is the top makes `Pop` and `Top` trivial – no reverse scanning, no second container.

---

**[GitHub](https://github.com/cloudnative0x0/leetcode-design)**
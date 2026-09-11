package _703_kth_largest_element

import "container/heap"

type ScoreHeap []int

func (h ScoreHeap) Len() int {
	return len(h)
}

func (h ScoreHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h ScoreHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ScoreHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *ScoreHeap) Pop() any {
	current := *h
	n := len(current)
	x := current[n-1]
	*h = current[0 : n-1]

	return x
}

func (h *ScoreHeap) Peek() int {
	return (*h)[0]
}

type KthLargest struct {
	h *ScoreHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	startingScores := ScoreHeap(nums)

	o := KthLargest{h: &startingScores, k: k}

	heap.Init(o.h)

	for range len(nums) - k {
		heap.Pop(o.h)
	}

	return o
}

func (kl *KthLargest) Add(val int) int {
	if kl.h.Len() < kl.k {
		heap.Push(kl.h, val)

		return kl.h.Peek()
	}

	current := kl.h.Peek()

	if val > current {
		(*(kl.h))[0] = val
		heap.Fix(kl.h, 0)

		current = kl.h.Peek()
	}

	return current
}

package _155_min_stack

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (ms *MinStack) Push(value int) {
	ms.stack = append(ms.stack, value)

	if len(ms.minStack) == 0 || value <= ms.minStack[len(ms.minStack)-1] {
		ms.minStack = append(ms.minStack, value)
	}
}

func (ms *MinStack) Pop() {
	if len(ms.stack) == 0 {
		return
	}

	topVal := ms.stack[len(ms.stack)-1]

	if topVal == ms.minStack[len(ms.minStack)-1] {
		ms.minStack = ms.minStack[:len(ms.minStack)-1]
	}

	ms.stack = ms.stack[:len(ms.stack)-1]
}

func (ms *MinStack) Top() int {
	if len(ms.stack) == 0 {
		return 0
	}

	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	if len(ms.minStack) == 0 {
		return 0
	}

	return ms.minStack[len(ms.minStack)-1]
}

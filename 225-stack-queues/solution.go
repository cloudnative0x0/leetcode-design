package _225_stack_queues

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

func (ms *MyStack) Push(x int) {
	ms.queue = append(ms.queue, x)

	for i := 0; i < len(ms.queue)-1; i++ {
		head := ms.queue[0]
		ms.queue = ms.queue[1:]
		ms.queue = append(ms.queue, head)
	}
}

func (ms *MyStack) Pop() int {
	head := ms.queue[0]
	ms.queue = ms.queue[1:]

	return head
}

func (ms *MyStack) Top() int {
	return ms.queue[0]
}

func (ms *MyStack) Empty() bool {
	return len(ms.queue) == 0
}

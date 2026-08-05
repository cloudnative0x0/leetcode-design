package _232_queue_stacks

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

func (mq *MyQueue) Push(x int) {
	mq.enStack = append(mq.enStack, x)
}

func (mq *MyQueue) Pop() int {
	mq.move()

	topIdx := len(mq.deStack) - 1
	topVal := mq.deStack[topIdx]

	mq.deStack = mq.deStack[:topIdx]

	return topVal
}

func (mq *MyQueue) move() {
	if len(mq.deStack) == 0 {
		for len(mq.enStack) > 0 {
			top := len(mq.enStack) - 1
			topVal := mq.enStack[top]

			mq.enStack = mq.enStack[:top]

			mq.deStack = append(mq.deStack, topVal)
		}
	}
}

func (mq *MyQueue) Peek() int {
	mq.move()

	return mq.deStack[len(mq.deStack)-1]
}

func (mq *MyQueue) Empty() bool {
	return len(mq.enStack) == 0 && len(mq.deStack) == 0
}

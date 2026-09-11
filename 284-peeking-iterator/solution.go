package _284_peeking_iterator

type Iterator struct {
	values []int
	index  int
}

func NewIterator(values []int) *Iterator {
	return &Iterator{values: values}
}

func (it *Iterator) hasNext() bool {
	return it.index < len(it.values)
}

func (it *Iterator) next() int {
	value := it.values[it.index]
	it.index++
	return value
}

type PeekingIterator struct {
	iter *Iterator
	arr  []int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter: iter,
		arr:  make([]int, 0),
	}
}

func (pi *PeekingIterator) hasNext() bool {
	if len(pi.arr) > 0 {
		return true
	}

	return pi.iter.hasNext()
}

func (pi *PeekingIterator) next() int {
	if len(pi.arr) > 0 {
		value := pi.arr[0]
		pi.arr = pi.arr[1:]
		return value
	}

	return pi.iter.next()
}

func (pi *PeekingIterator) peek() int {
	if len(pi.arr) > 0 {
		return pi.arr[0]
	}

	pi.arr = append(pi.arr, pi.iter.next())
	return pi.arr[0]
}

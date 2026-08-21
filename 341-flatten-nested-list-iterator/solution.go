package _341_flatten_nested_list_iterator

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger struct {
	value int
	list  []*NestedInteger
	isInt bool
}

func (ni NestedInteger) IsInteger() bool {
	return ni.isInt
}

func (ni NestedInteger) GetInteger() int {
	return ni.value
}

func (n *NestedInteger) SetInteger(value int) {
	n.value = value
	n.isInt = true
	n.list = nil
}

func (ni *NestedInteger) Add(elem NestedInteger) {
	ni.isInt = false
	ni.list = append(ni.list, &elem)
}

func (ni NestedInteger) GetList() []*NestedInteger {
	return ni.list
}

func NewInt(v int) *NestedInteger {
	return &NestedInteger{value: v, isInt: true}
}

func NewList(items ...*NestedInteger) *NestedInteger {
	return &NestedInteger{list: items, isInt: false}
}

type NestedIterator struct {
	stack [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		stack: [][]*NestedInteger{nestedList},
	}
}

func (ni *NestedIterator) Next() int {
	top := ni.stack[len(ni.stack)-1]

	elem := top[0]

	ni.stack[len(ni.stack)-1] = top[1:]

	return elem.GetInteger()
}

func (ni *NestedIterator) HasNext() bool {
	for len(ni.stack) > 0 {
		top := len(ni.stack) - 1

		if len(ni.stack[top]) == 0 {
			ni.stack = ni.stack[:top]
			continue
		}

		firstElem := ni.stack[top][0]

		if firstElem.IsInteger() {
			return true
		}

		ni.stack[top] = ni.stack[top][1:]

		ni.stack = append(ni.stack, firstElem.GetList())
	}

	return false
}

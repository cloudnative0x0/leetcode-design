package _380_insert_delete_get_random_o1

import "math/rand/v2"

type RandomizedSet struct {
	indexMap map[int]int
	elements []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		indexMap: make(map[int]int),
		elements: make([]int, 0),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.indexMap[val]; ok {
		return false
	}

	rs.indexMap[val] = len(rs.elements)

	rs.elements = append(rs.elements, val)

	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	idx, ok := rs.indexMap[val]
	if !ok {
		return false
	}

	lastIdx := len(rs.elements) - 1
	lastVal := rs.elements[lastIdx]

	rs.elements[idx] = lastVal
	rs.indexMap[lastVal] = idx

	rs.elements = rs.elements[:lastIdx]

	delete(rs.indexMap, val)

	return true
}

func (rs *RandomizedSet) GetRandom() int {
	randomIndex := rand.N(len(rs.elements))

	return rs.elements[randomIndex]
}

package _341_flatten_nested_list_iterator

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func flattenReference(nestedList []*NestedInteger) []int {
	result := make([]int, 0)
	for _, ni := range nestedList {
		flattenOne(ni, &result)
	}

	return result
}

func flattenOne(ni *NestedInteger, out *[]int) {
	if ni.IsInteger() {
		*out = append(*out, ni.GetInteger())
		return
	}
	for _, child := range ni.GetList() {
		flattenOne(child, out)
	}
}

type genOptions struct {
	maxDepth      int
	maxWidth      int
	emptyListProb float64
	valueRange    int
}

func defaultGenOptions() genOptions {
	return genOptions{
		maxDepth:      6,
		maxWidth:      6,
		emptyListProb: 0.15,
		valueRange:    1000,
	}
}

func genNestedList(r *rand.Rand, depth int, opt genOptions) []*NestedInteger {
	width := r.Intn(opt.maxWidth + 1)
	list := make([]*NestedInteger, 0, width)

	for i := 0; i < width; i++ {
		list = append(list, genNestedInteger(r, depth, opt))
	}

	return list
}

func genNestedInteger(r *rand.Rand, depth int, opt genOptions) *NestedInteger {
	canNest := depth < opt.maxDepth

	if !canNest || r.Float64() < 0.5 {
		v := r.Intn(2*opt.valueRange+1) - opt.valueRange
		return NewInt(v)
	}

	if r.Float64() < opt.emptyListProb {
		return NewList()
	}

	children := genNestedList(r, depth+1, opt)

	return &NestedInteger{list: children, isInt: false}
}

func drainStrict(it *NestedIterator) []int {
	out := make([]int, 0)
	for it.HasNext() {
		out = append(out, it.Next())
	}

	return out
}

func drainWithRedundantHasNext(t *testing.T, it *NestedIterator, r *rand.Rand) []int {
	out := make([]int, 0)
	for {
		extra := r.Intn(3)

		var last bool
		for i := 0; i <= extra; i++ {
			last = it.HasNext()
		}

		if !last {
			break
		}

		out = append(out, it.Next())
	}

	for i := 0; i < 5; i++ {
		if it.HasNext() {
			t.Fatalf("HasNext() returned true after iterator (repeated #%d)", i)
		}
	}

	return out
}

func TestStress_NestedIterator_RandomStructures(t *testing.T) {
	seed := time.Now().UnixNano()
	t.Logf("seed=%d", seed)
	r := rand.New(rand.NewSource(seed))

	const iterations = 2000
	opt := defaultGenOptions()

	for i := 0; i < iterations; i++ {
		nestedList := genNestedList(r, 0, opt)
		expected := flattenReference(nestedList)

		got1 := drainStrict(Constructor(nestedList))
		if !equalIntSlices(expected, got1) {
			t.Fatalf("iteration %d (strict): mismatch\nexpected=%v\ngot=%v\nstructure=%s",
				i, expected, got1, describeStructure(nestedList))
		}

		got2 := drainWithRedundantHasNext(t, Constructor(nestedList), r)
		if !equalIntSlices(expected, got2) {
			t.Fatalf("iteration %d (redundant HasNext): mismatch\nexpected=%v\ngot=%v\nstructure=%s",
				i, expected, got2, describeStructure(nestedList))
		}
	}
}

func TestStress_NestedIterator_EdgeCases(t *testing.T) {
	cases := []struct {
		name     string
		build    func() []*NestedInteger
		expected []int
	}{
		{
			build:    func() []*NestedInteger { return []*NestedInteger{} },
			expected: []int{},
		},
		{
			build: func() []*NestedInteger {
				return []*NestedInteger{NewList()}
			},
			expected: []int{},
		},
		{
			build: func() []*NestedInteger {
				return []*NestedInteger{NewList(), NewList(), NewList(NewList()), NewInt(42)}
			},
			expected: []int{42},
		},
		{
			build: func() []*NestedInteger {
				return []*NestedInteger{
					NewList(NewInt(1), NewInt(1)),
					NewInt(2),
					NewList(NewInt(1), NewInt(1)),
				}
			},
			expected: []int{1, 1, 2, 1, 1},
		},
		{
			build: func() []*NestedInteger {
				return []*NestedInteger{
					NewInt(1),
					NewList(NewInt(4), NewList(NewInt(6))),
				}
			},
			expected: []int{1, 4, 6},
		},
		{
			build: func() []*NestedInteger {
				inner := NewList(NewInt(7))
				for i := 0; i < 50; i++ {
					inner = NewList(NewList(), inner, NewList())
				}
				return []*NestedInteger{inner}
			},
			expected: []int{7},
		},
		{
			build: func() []*NestedInteger {
				return []*NestedInteger{NewInt(-5), NewList(NewInt(0), NewInt(-1)), NewInt(0)}
			},
			expected: []int{-5, 0, -1, 0},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			it := Constructor(c.build())
			got := drainStrict(it)
			if !equalIntSlices(c.expected, got) {
				t.Fatalf("mismatch\nexpected=%v\ngot=%v", c.expected, got)
			}
		})
	}
}

func TestStress_NestedIterator_LargeScale(t *testing.T) {
	if testing.Short() {
		t.Skip("missed in -short")
	}

	seed := time.Now().UnixNano()
	t.Logf("seed=%d", seed)
	r := rand.New(rand.NewSource(seed))

	t.Run("", func(t *testing.T) {
		opt := genOptions{maxDepth: 3, maxWidth: 60, emptyListProb: 0.1, valueRange: 1_000_000}
		nestedList := genNestedList(r, 0, opt)
		expected := flattenReference(nestedList)
		got := drainStrict(Constructor(nestedList))
		if !equalIntSlices(expected, got) {
			t.Fatalf("mismatch: len(expected)=%d len(got)=%d", len(expected), len(got))
		}
		t.Logf("flatten elements: %d", len(expected))
	})

	t.Run("deep structure", func(t *testing.T) {
		const depth = 5000
		inner := NewInt(123)
		for i := 0; i < depth; i++ {
			inner = NewList(inner)
		}
		nestedList := []*NestedInteger{inner}

		expected := []int{123}
		got := drainStrict(Constructor(nestedList))
		if !equalIntSlices(expected, got) {
			t.Fatalf("mismatch on deep structure: expected=%v got=%v", expected, got)
		}
	})
}

func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func describeStructure(list []*NestedInteger) string {
	s := "["
	for i, ni := range list {
		if i > 0 {
			s += ","
		}
		s += describeOne(ni)
	}
	s += "]"

	return s
}

func describeOne(ni *NestedInteger) string {
	if ni.IsInteger() {
		return fmt.Sprintf("%d", ni.GetInteger())
	}

	return describeStructure(ni.GetList())
}

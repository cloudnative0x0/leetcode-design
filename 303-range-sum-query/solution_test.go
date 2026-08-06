package _303_range_sum_query

import (
	"math/rand"
	"testing"
)

func TestSumRange_Example(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	na := Constructor(nums)

	tests := []struct {
		left     int
		right    int
		expected int
	}{
		{0, 2, 1},  // -2 + 0 + 3 = 1
		{2, 5, -1}, // 3 + -5 + 2 + -1 = -1
		{0, 5, -3}, // sum of all elems
	}

	for _, tt := range tests {
		got := na.SumRange(tt.left, tt.right)
		if got != tt.expected {
			t.Errorf("SumRange(%d, %d) = %d, expected %d", tt.left, tt.right, got, tt.expected)
		}
	}
}

func TestSumRange_SingleElement(t *testing.T) {
	nums := []int{7}
	na := Constructor(nums)

	got := na.SumRange(0, 0)
	if got != 7 {
		t.Errorf("SumRange(0, 0) = %d, expected 7", got)
	}
}

func TestSumRange_AllZeros(t *testing.T) {
	nums := []int{0, 0, 0, 0}
	na := Constructor(nums)

	got := na.SumRange(0, 3)
	if got != 0 {
		t.Errorf("SumRange(0, 3) = %d, expected 0", got)
	}
}

func TestSumRange_NegativeNumbers(t *testing.T) {
	nums := []int{-1, -2, -3, -4}
	na := Constructor(nums)

	got := na.SumRange(1, 3)
	expected := -2 + -3 + -4
	if got != expected {
		t.Errorf("SumRange(1, 3) = %d, expected %d", got, expected)
	}
}

func TestSumRange_LeftEqualsRight(t *testing.T) {
	nums := []int{5, 10, 15, 20}
	na := Constructor(nums)

	for i, want := range nums {
		got := na.SumRange(i, i)
		if got != want {
			t.Errorf("SumRange(%d, %d) = %d, expected %d", i, i, got, want)
		}
	}
}

func TestSumRange_MultipleQueriesOnSameArray(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	na := Constructor(nums)

	if got := na.SumRange(0, 4); got != 15 {
		t.Errorf("SumRange(0, 4) = %d, expected 15", got)
	}
	if got := na.SumRange(1, 3); got != 9 {
		t.Errorf("SumRange(1, 3) = %d, expected 9", got)
	}
	if got := na.SumRange(0, 4); got != 15 {
		t.Errorf("повторный SumRange(0, 4) = %d, expected 15", got)
	}
}

func bruteForceSumRange(nums []int, left, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += nums[i]
	}
	return sum
}

func TestSumRange_Stress(t *testing.T) {
	const iterations = 2000
	const maxSize = 200
	const maxQueries = 50

	rng := rand.New(rand.NewSource(42))

	for iter := 0; iter < iterations; iter++ {
		n := rng.Intn(maxSize) + 1
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rng.Intn(2001) - 1000
		}

		na := Constructor(nums)

		queries := rng.Intn(maxQueries) + 1
		for q := 0; q < queries; q++ {
			a := rng.Intn(n)
			b := rng.Intn(n)
			left, right := a, b
			if left > right {
				left, right = right, left
			}

			got := na.SumRange(left, right)
			expected := bruteForceSumRange(nums, left, right)

			if got != expected {
				t.Fatalf(
					"iter=%d nums=%v SumRange(%d, %d) = %d, expected %d",
					iter, nums, left, right, got, expected,
				)
			}
		}
	}
}

func BenchmarkSumRange(b *testing.B) {
	nums := make([]int, 100000)
	for i := range nums {
		nums[i] = i
	}
	na := Constructor(nums)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		na.SumRange(0, len(nums)-1)
	}
}

func BenchmarkConstructor(b *testing.B) {
	nums := make([]int, 100000)
	for i := range nums {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Constructor(nums)
	}
}

package _703_kth_largest_element

import "testing"

func TestKthLargestExample(t *testing.T) {
	kthLargest := Constructor(3, []int{4, 5, 8, 2})

	tests := []struct {
		value int
		want  int
	}{
		{value: 3, want: 4},
		{value: 5, want: 5},
		{value: 10, want: 5},
		{value: 9, want: 8},
		{value: 4, want: 8},
	}

	for _, tt := range tests {
		if got := kthLargest.Add(tt.value); got != tt.want {
			t.Errorf("Add(%d) = %d; want %d", tt.value, got, tt.want)
		}
	}
}

func TestConstructorWithMoreThanKValues(t *testing.T) {
	kthLargest := Constructor(2, []int{9, 1, 7, 3, 8})

	if got := kthLargest.h.Len(); got != 2 {
		t.Fatalf("heap length = %d; want 2", got)
	}

	if got := kthLargest.Add(6); got != 8 {
		t.Errorf("Add(6) = %d; want 8", got)
	}

	if got := kthLargest.Add(10); got != 9 {
		t.Errorf("Add(10) = %d; want 9", got)
	}
}

func TestConstructorWithFewerThanKValues(t *testing.T) {
	kthLargest := Constructor(4, []int{5, 2, 8})

	tests := []struct {
		value int
		want  int
	}{
		{value: 1, want: 1},
		{value: 10, want: 2},
		{value: 9, want: 5},
	}

	for _, tt := range tests {
		if got := kthLargest.Add(tt.value); got != tt.want {
			t.Errorf("Add(%d) = %d; want %d", tt.value, got, tt.want)
		}
	}
}

func TestKEqualsOne(t *testing.T) {
	kthLargest := Constructor(1, []int{-10, -7, -20})

	tests := []struct {
		value int
		want  int
	}{
		{value: -15, want: -7},
		{value: -3, want: -3},
		{value: -3, want: -3},
		{value: 0, want: 0},
	}

	for _, tt := range tests {
		if got := kthLargest.Add(tt.value); got != tt.want {
			t.Errorf("Add(%d) = %d; want %d", tt.value, got, tt.want)
		}
	}
}

func TestDuplicateValues(t *testing.T) {
	kthLargest := Constructor(3, []int{5, 5, 5})

	if got := kthLargest.Add(5); got != 5 {
		t.Errorf("Add(5) = %d; want 5", got)
	}

	if got := kthLargest.Add(6); got != 5 {
		t.Errorf("Add(6) = %d; want 5", got)
	}

	if got := kthLargest.Add(7); got != 5 {
		t.Errorf("Add(7) = %d; want 5", got)
	}

	if got := kthLargest.Add(8); got != 6 {
		t.Errorf("Add(8) = %d; want 6", got)
	}
}

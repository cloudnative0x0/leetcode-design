package _225_stack_queues

import (
	"math/rand"
	"testing"
)

func TestStackBasic(t *testing.T) {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Top() != 3 {
		t.Fatalf("Top() = %d, want 3", s.Top())
	}
	if s.Pop() != 3 {
		t.Fatalf("Pop() = %d, want 3", s.Pop())
	}
	if s.Pop() != 2 {
		t.Fatalf("Pop() = %d, want 2", s.Pop())
	}
	if s.Pop() != 1 {
		t.Fatalf("Pop() = %d, want 1", s.Pop())
	}
	if !s.Empty() {
		t.Fatal("stack should be empty")
	}
}

// TestStackEmpty проверяет Empty на пустом и непустом стеке.
func TestStackEmpty(t *testing.T) {
	s := Constructor()
	if !s.Empty() {
		t.Fatal("new stack must be empty")
	}
	s.Push(42)
	if s.Empty() {
		t.Fatal("stack must not be empty after Push")
	}
	_ = s.Pop()
	if !s.Empty() {
		t.Fatal("stack must be empty after Pop")
	}
}

func TestStressStack(t *testing.T) {
	const operations = 200_000
	s := Constructor()
	ref := make([]int, 0)

	for i := 0; i < operations; i++ {
		if len(ref) == 0 {
			val := rand.Intn(10000)
			s.Push(val)
			ref = append(ref, val)
			continue
		}

		switch rand.Intn(3) {
		case 0: // Push
			val := rand.Intn(10000)
			s.Push(val)
			ref = append(ref, val)

		case 1: // Pop
			expected := ref[len(ref)-1]
			ref = ref[:len(ref)-1]
			if got := s.Pop(); got != expected {
				t.Fatalf("Pop() = %d, want %d", got, expected)
			}

		case 2: // Top
			expected := ref[len(ref)-1]
			if got := s.Top(); got != expected {
				t.Fatalf("Top() = %d, want %d", got, expected)
			}
		}
	}
}

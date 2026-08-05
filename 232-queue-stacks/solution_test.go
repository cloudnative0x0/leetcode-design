package _232_queue_stacks

import (
	"math/rand"
	"testing"
)

func TestStressMyQueue(t *testing.T) {
	const operations = 500_000
	q := Constructor()
	ref := make([]int, 0)

	for i := 0; i < operations; i++ {
		if len(ref) == 0 {
			val := rand.Intn(10000)
			q.Push(val)
			ref = append(ref, val)
			continue
		}

		switch rand.Intn(3) {
		case 0: // Push
			val := rand.Intn(10000)
			q.Push(val)
			ref = append(ref, val)

		case 1: // Pop
			if q.Empty() != (len(ref) == 0) {
				t.Fatalf("Empty() mismatch: got %v, reference size %d", q.Empty(), len(ref))
			}
			expected := ref[0]
			ref = ref[1:]
			got := q.Pop()
			if got != expected {
				t.Fatalf("Pop() = %d, want %d", got, expected)
			}

		case 2: // Peek
			expected := ref[0]
			got := q.Peek()
			if got != expected {
				t.Fatalf("Peek() = %d, want %d", got, expected)
			}
		}
	}
}

func TestAlternatingPushPop(t *testing.T) {
	q := Constructor()
	const rounds = 1000
	const batchSize = 2000

	for r := 0; r < rounds; r++ {
		// push batch
		for i := 0; i < batchSize; i++ {
			q.Push(i)
		}
		// pop batch
		for i := 0; i < batchSize; i++ {
			if got := q.Pop(); got != i {
				t.Fatalf("round %d: Pop() = %d, want %d", r, got, i)
			}
		}
		if !q.Empty() {
			t.Fatalf("round %d: queue not empty after full pop", r)
		}
	}
}

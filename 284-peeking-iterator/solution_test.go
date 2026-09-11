package _284_peeking_iterator

import "testing"

func TestIterator(t *testing.T) {
	it := NewIterator([]int{4, 8})

	if !it.hasNext() {
		t.Fatal("hasNext() = false before the first element")
	}
	if got := it.next(); got != 4 {
		t.Fatalf("first next() = %d; want 4", got)
	}
	if got := it.next(); got != 8 {
		t.Fatalf("second next() = %d; want 8", got)
	}
	if it.hasNext() {
		t.Fatal("hasNext() = true after the last element")
	}
}

func TestPeekingIteratorExample(t *testing.T) {
	pi := Constructor(NewIterator([]int{1, 2, 3}))

	assertNext(t, pi, 1)
	assertPeek(t, pi, 2)
	assertNext(t, pi, 2)
	assertNext(t, pi, 3)

	if pi.hasNext() {
		t.Fatal("hasNext() = true after all elements were read")
	}
}

func TestPeekDoesNotAdvance(t *testing.T) {
	pi := Constructor(NewIterator([]int{7, 8}))

	assertPeek(t, pi, 7)
	assertPeek(t, pi, 7)
	assertNext(t, pi, 7)
	assertPeek(t, pi, 8)
	assertNext(t, pi, 8)
}

func TestHasNextUsesBufferedElement(t *testing.T) {
	pi := Constructor(NewIterator([]int{42}))

	assertPeek(t, pi, 42)

	if pi.iter.hasNext() {
		t.Fatal("underlying iterator must be exhausted after peek()")
	}
	if !pi.hasNext() {
		t.Fatal("hasNext() = false while an element is buffered")
	}

	assertNext(t, pi, 42)

	if pi.hasNext() {
		t.Fatal("hasNext() = true after the buffered element was read")
	}
}

func TestOrderIsPreserved(t *testing.T) {
	values := []int{3, 1, 4, 1, 5}
	pi := Constructor(NewIterator(values))

	for _, want := range values {
		if !pi.hasNext() {
			t.Fatalf("hasNext() = false before value %d", want)
		}

		assertPeek(t, pi, want)
		assertNext(t, pi, want)
	}

	if pi.hasNext() {
		t.Fatal("hasNext() = true after all elements were read")
	}
}

func assertPeek(t *testing.T, pi *PeekingIterator, want int) {
	t.Helper()

	if got := pi.peek(); got != want {
		t.Fatalf("peek() = %d; want %d", got, want)
	}
}

func assertNext(t *testing.T, pi *PeekingIterator, want int) {
	t.Helper()

	if got := pi.next(); got != want {
		t.Fatalf("next() = %d; want %d", got, want)
	}
}

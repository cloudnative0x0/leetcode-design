package _707_linked_list

import (
	"fmt"
	"math/rand"
	"testing"
)

type refModel struct {
	data []int
}

func (r *refModel) Get(index int) int {
	if index < 0 || index >= len(r.data) {
		return -1
	}
	return r.data[index]
}

func (r *refModel) AddAtHead(val int) {
	r.data = append([]int{val}, r.data...)
}

func (r *refModel) AddAtTail(val int) {
	r.data = append(r.data, val)
}

func (r *refModel) AddAtIndex(index, val int) {
	if index < 0 || index > len(r.data) {
		return
	}
	r.data = append(r.data, 0)
	copy(r.data[index+1:], r.data[index:])
	r.data[index] = val
}

func (r *refModel) DeleteAtIndex(index int) {
	if index < 0 || index >= len(r.data) {
		return
	}
	r.data = append(r.data[:index], r.data[index+1:]...)
}

func snapshot(mll *MyLinkedList, size int) []int {
	out := make([]int, size)
	for i := 0; i < size; i++ {
		out[i] = mll.Get(i)
	}
	return out
}

func equal(a, b []int) bool {
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

func TestStress(t *testing.T) {
	const (
		iterations = 200
		opsPerRun  = 500
		valueRange = 1000
		indexSlack = 3
		seed       = 42
	)

	rng := rand.New(rand.NewSource(seed))

	for run := 0; run < iterations; run++ {
		mll := Constructor()
		ref := &refModel{}

		for op := 0; op < opsPerRun; op++ {
			action := rng.Intn(5)
			val := rng.Intn(valueRange) - valueRange/2

			switch action {
			case 0: // Get
				idx := rng.Intn(len(ref.data)+indexSlack+1) - indexSlack
				got := mll.Get(idx)
				want := ref.Get(idx)
				if got != want {
					t.Fatalf("run=%d op=%d Get(%d) = %d, want %d (state=%v)",
						run, op, idx, got, want, ref.data)
				}

			case 1: // AddAtHead
				mll.AddAtHead(val)
				ref.AddAtHead(val)

			case 2: // AddAtTail
				mll.AddAtTail(val)
				ref.AddAtTail(val)

			case 3: // AddAtIndex
				idx := rng.Intn(len(ref.data)+indexSlack+1) - indexSlack
				mll.AddAtIndex(idx, val)
				ref.AddAtIndex(idx, val)

			case 4: // DeleteAtIndex
				idx := rng.Intn(len(ref.data)+indexSlack+1) - indexSlack
				mll.DeleteAtIndex(idx)
				ref.DeleteAtIndex(idx)
			}

			got := snapshot(&mll, len(ref.data))
			if !equal(got, ref.data) {
				t.Fatalf("run=%d op=%d state mismatch:\n  got:  %v\n  want: %v",
					run, op, got, ref.data)
			}

			if mll.size != len(ref.data) {
				t.Fatalf("run=%d op=%d size mismatch: got %d, want %d",
					run, op, mll.size, len(ref.data))
			}

			checkTailConsistency(t, &mll, run, op)
		}
	}
}

func checkTailConsistency(t *testing.T, mll *MyLinkedList, run, op int) {
	t.Helper()

	if mll.size == 0 {
		if mll.head != nil || mll.tail != nil {
			t.Fatalf("run=%d op=%d: expected head/tail nil on empty list, got head=%v tail=%v",
				run, op, mll.head, mll.tail)
		}
		return
	}

	cur := mll.head
	for i := 0; i < mll.size-1; i++ {
		if cur == nil {
			t.Fatalf("run=%d op=%d: list shorter than size=%d", run, op, mll.size)
		}
		cur = cur.Next
	}

	if cur != mll.tail {
		t.Fatalf("run=%d op=%d: tail pointer inconsistent, expected node value %v, tail value %v",
			run, op, cur.Value, mll.tail.Value)
	}
	if mll.tail.Next != nil {
		t.Fatalf("run=%d op=%d: tail.Next is not nil", run, op)
	}
}

func TestStressFuzzLike(t *testing.T) {
	rng := rand.New(rand.NewSource(1337))

	for run := 0; run < 50; run++ {
		mll := Constructor()
		ref := &refModel{}
		var history []string

		for op := 0; op < 300; op++ {
			action := rng.Intn(5)
			val := rng.Intn(200) - 100

			switch action {
			case 0:
				idx := rng.Intn(len(ref.data)+4) - 2
				got := mll.Get(idx)
				want := ref.Get(idx)
				history = append(history, fmt.Sprintf("Get(%d)", idx))
				if got != want {
					t.Fatalf("run=%d op=%d mismatch Get(%d)=%d want %d\nhistory: %v",
						run, op, idx, got, want, history)
				}
				continue
			case 1:
				mll.AddAtHead(val)
				ref.AddAtHead(val)
				history = append(history, fmt.Sprintf("AddAtHead(%d)", val))
			case 2:
				mll.AddAtTail(val)
				ref.AddAtTail(val)
				history = append(history, fmt.Sprintf("AddAtTail(%d)", val))
			case 3:
				idx := rng.Intn(len(ref.data)+4) - 2
				mll.AddAtIndex(idx, val)
				ref.AddAtIndex(idx, val)
				history = append(history, fmt.Sprintf("AddAtIndex(%d, %d)", idx, val))
			case 4:
				idx := rng.Intn(len(ref.data)+4) - 2
				mll.DeleteAtIndex(idx)
				ref.DeleteAtIndex(idx)
				history = append(history, fmt.Sprintf("DeleteAtIndex(%d)", idx))
			}

			got := snapshot(&mll, len(ref.data))
			if !equal(got, ref.data) {
				t.Fatalf("run=%d op=%d state mismatch:\n  got:  %v\n  want: %v\n  history: %v",
					run, op, got, ref.data, history)
			}
		}
	}
}

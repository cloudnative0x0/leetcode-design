package _1656_ordered_stream

import (
	"reflect"
	"testing"
)

func TestOrderedStream_InsertSequential(t *testing.T) {
	os := Constructor(5)

	result := os.Insert(1, "aaaaa")
	if !reflect.DeepEqual(result, []string{"aaaaa"}) {
		t.Errorf("expected [aaaaa], got %v", result)
	}

	result = os.Insert(2, "bbbbb")
	if !reflect.DeepEqual(result, []string{"bbbbb"}) {
		t.Errorf("expected [bbbbb], got %v", result)
	}

	result = os.Insert(3, "ccccc")
	if !reflect.DeepEqual(result, []string{"ccccc"}) {
		t.Errorf("expected [ccccc], got %v", result)
	}
}

func TestOrderedStream_InsertGap(t *testing.T) {
	os := Constructor(5)

	result := os.Insert(3, "ccccc")
	if len(result) != 0 {
		t.Errorf("expected empty slice, got %v", result)
	}

	result = os.Insert(1, "aaaaa")
	if !reflect.DeepEqual(result, []string{"aaaaa"}) {
		t.Errorf("expected [aaaaa], got %v", result)
	}

	result = os.Insert(2, "bbbbb")
	if !reflect.DeepEqual(result, []string{"bbbbb", "ccccc"}) {
		t.Errorf("expected [bbbbb ccccc], got %v", result)
	}

	result = os.Insert(4, "ddddd")
	if !reflect.DeepEqual(result, []string{"ddddd"}) {
		t.Errorf("expected [ddddd], got %v", result)
	}

	result = os.Insert(5, "eeeee")
	if !reflect.DeepEqual(result, []string{"eeeee"}) {
		t.Errorf("expected [eeeee], got %v", result)
	}
}

func TestOrderedStream_InsertMultipleChunks(t *testing.T) {
	os := Constructor(6)

	os.Insert(2, "b")
	os.Insert(4, "d")
	os.Insert(6, "f")

	result := os.Insert(1, "a")
	if !reflect.DeepEqual(result, []string{"a", "b"}) {
		t.Errorf("expected [a b], got %v", result)
	}

	result = os.Insert(3, "c")
	if !reflect.DeepEqual(result, []string{"c", "d"}) {
		t.Errorf("expected [c d], got %v", result)
	}

	result = os.Insert(5, "e")
	if !reflect.DeepEqual(result, []string{"e", "f"}) {
		t.Errorf("expected [e f], got %v", result)
	}
}

func TestOrderedStream_InsertAtEnd(t *testing.T) {
	os := Constructor(3)

	os.Insert(2, "b")
	os.Insert(3, "c")

	result := os.Insert(1, "a")
	if !reflect.DeepEqual(result, []string{"a", "b", "c"}) {
		t.Errorf("expected [a b c], got %v", result)
	}
}

func TestOrderedStream_InsertAfterPtrReachesEnd(t *testing.T) {
	os := Constructor(2)

	os.Insert(1, "a")
	os.Insert(2, "b")

	result := os.Insert(1, "x") // id already filled but we ignore, ptr not moved
	if len(result) != 0 {
		t.Errorf("expected empty slice, got %v", result)
	}
}

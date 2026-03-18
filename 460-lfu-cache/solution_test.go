package _460_lfu_cache

import "testing"

func TestLFUCache_BasicExample(t *testing.T) {
	lfu := Constructor(2)

	lfu.Put(1, 1)
	lfu.Put(2, 2)

	if got := lfu.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, want 1", got)
	}

	lfu.Put(3, 3)

	if got := lfu.Get(2); got != -1 {
		t.Errorf("Get(2) = %d, want -1", got)
	}
	if got := lfu.Get(3); got != 3 {
		t.Errorf("Get(3) = %d, want 3", got)
	}

	lfu.Put(4, 4)

	if got := lfu.Get(1); got != -1 {
		t.Errorf("Get(1) = %d, want -1", got)
	}
	if got := lfu.Get(3); got != 3 {
		t.Errorf("Get(3) = %d, want 3", got)
	}
	if got := lfu.Get(4); got != 4 {
		t.Errorf("Get(4) = %d, want 4", got)
	}
}

func TestLFUCache_GetMiss(t *testing.T) {
	lfu := Constructor(2)

	if got := lfu.Get(99); got != -1 {
		t.Errorf("Get on empty cache = %d, want -1", got)
	}

	lfu.Put(1, 10)
	if got := lfu.Get(2); got != -1 {
		t.Errorf("Get missing key = %d, want -1", got)
	}
}

func TestLFUCache_UpdateExistingKey(t *testing.T) {
	lfu := Constructor(2)

	lfu.Put(1, 1)
	lfu.Put(1, 100)

	if got := lfu.Get(1); got != 100 {
		t.Errorf("Get(1) after update = %d, want 100", got)
	}
}

func TestLFUCache_EvictLFUNotLRU(t *testing.T) {
	lfu := Constructor(3)

	lfu.Put(1, 1)
	lfu.Put(2, 2)
	lfu.Put(3, 3)

	lfu.Get(1)
	lfu.Get(1)
	lfu.Get(2)

	lfu.Put(4, 4)

	if got := lfu.Get(3); got != -1 {
		t.Errorf("Get(3) = %d, want -1 (should be evicted)", got)
	}
	if got := lfu.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, want 1", got)
	}
	if got := lfu.Get(2); got != 2 {
		t.Errorf("Get(2) = %d, want 2", got)
	}
	if got := lfu.Get(4); got != 4 {
		t.Errorf("Get(4) = %d, want 4", got)
	}
}

func TestLFUCache_TieBreakByLRU(t *testing.T) {
	lfu := Constructor(2)

	lfu.Put(1, 1)
	lfu.Put(2, 2)

	lfu.Put(3, 3)

	if got := lfu.Get(1); got != -1 {
		t.Errorf("Get(1) = %d, want -1 (LRU evicted)", got)
	}
	if got := lfu.Get(2); got != 2 {
		t.Errorf("Get(2) = %d, want 2", got)
	}
	if got := lfu.Get(3); got != 3 {
		t.Errorf("Get(3) = %d, want 3", got)
	}
}

func TestLFUCache_CapacityOne(t *testing.T) {
	lfu := Constructor(1)

	lfu.Put(1, 1)
	if got := lfu.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, want 1", got)
	}

	lfu.Put(2, 2)
	if got := lfu.Get(1); got != -1 {
		t.Errorf("Get(1) = %d, want -1", got)
	}
	if got := lfu.Get(2); got != 2 {
		t.Errorf("Get(2) = %d, want 2", got)
	}

	lfu.Put(3, 3)
	if got := lfu.Get(2); got != -1 {
		t.Errorf("Get(2) = %d, want -1", got)
	}
	if got := lfu.Get(3); got != 3 {
		t.Errorf("Get(3) = %d, want 3", got)
	}
}

func TestLFUCache_ZeroCapacity(t *testing.T) {
	lfu := Constructor(0)

	lfu.Put(1, 1)
	if got := lfu.Get(1); got != -1 {
		t.Errorf("Get(1) on zero-cap cache = %d, want -1", got)
	}
}

func TestLFUCache_PutUpdatesFrequency(t *testing.T) {
	lfu := Constructor(2)

	lfu.Put(1, 1)
	lfu.Put(2, 2)
	lfu.Put(1, 10)

	lfu.Put(3, 3)

	if got := lfu.Get(2); got != -1 {
		t.Errorf("Get(2) = %d, want -1 (should be evicted)", got)
	}
	if got := lfu.Get(1); got != 10 {
		t.Errorf("Get(1) = %d, want 10", got)
	}
	if got := lfu.Get(3); got != 3 {
		t.Errorf("Get(3) = %d, want 3", got)
	}
}

func TestLFUCache_MinFreqResetOnInsert(t *testing.T) {
	lfu := Constructor(2)

	lfu.Put(1, 1)
	lfu.Put(2, 2)
	lfu.Get(1)
	lfu.Get(1)

	lfu.Put(3, 3)
	lfu.Put(4, 4)

	if got := lfu.Get(3); got != -1 {
		t.Errorf("Get(3) = %d, want -1", got)
	}
	if got := lfu.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, want 1", got)
	}
	if got := lfu.Get(4); got != 4 {
		t.Errorf("Get(4) = %d, want 4", got)
	}
}

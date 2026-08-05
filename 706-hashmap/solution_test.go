package _706_hashmap

import (
	"math/rand"
	"testing"
)

func TestMyHashMap_BaseScenario(t *testing.T) {
	hm := Constructor()

	hm.Put(1, 1)
	hm.Put(2, 2)

	if val := hm.Get(1); val != 1 {
		t.Errorf("Get(1) = %d, want 1", val)
	}
	if val := hm.Get(3); val != -1 {
		t.Errorf("Get(3) = %d, want -1", val)
	}

	hm.Put(2, 1) // update
	if val := hm.Get(2); val != 1 {
		t.Errorf("Get(2) after update = %d, want 1", val)
	}

	hm.Remove(2)
	if val := hm.Get(2); val != -1 {
		t.Errorf("Get(2) after remove = %d, want -1", val)
	}
}

func TestMyHashMap_Collisions(t *testing.T) {
	hm := Constructor()

	k1 := 100
	k2 := 100 + bucketCount
	k3 := 100 + 2*bucketCount

	hm.Put(k1, 10)
	hm.Put(k2, 20)
	hm.Put(k3, 30)

	if val := hm.Get(k1); val != 10 {
		t.Errorf("Get(k1) = %d, want 10", val)
	}
	if val := hm.Get(k2); val != 20 {
		t.Errorf("Get(k2) = %d, want 20", val)
	}
	if val := hm.Get(k3); val != 30 {
		t.Errorf("Get(k3) = %d, want 30", val)
	}

	hm.Remove(k2)
	if val := hm.Get(k2); val != -1 {
		t.Errorf("k2 should be removed, got %d", val)
	}
	if val := hm.Get(k1); val != 10 {
		t.Errorf("k1 should still be 10, got %d", val)
	}
	if val := hm.Get(k3); val != 30 {
		t.Errorf("k3 should still be 30, got %d", val)
	}
}

func TestMyHashMap_RemoveNonExistent(t *testing.T) {
	hm := Constructor()

	assertNoPanic(t, func() { hm.Remove(9999) })

	hm.Put(5, 50)
	assertNoPanic(t, func() { hm.Remove(5 + bucketCount) })
}

func assertNoPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic: %v", r)
		}
	}()
	f()
}

func TestStressHashMap(t *testing.T) {
	const ops = 300_000
	hm := Constructor()
	ref := make(map[int]int)

	for i := 0; i < ops; i++ {
		key := rand.Intn(20000)
		switch rand.Intn(3) {
		case 0: // Put
			val := rand.Intn(10000)
			hm.Put(key, val)
			ref[key] = val
		case 1: // Get
			got := hm.Get(key)
			exp, ok := ref[key]
			if !ok {
				exp = -1
			}
			if got != exp {
				t.Fatalf("Get(%d) = %d, want %d", key, got, exp)
			}
		case 2: // Remove
			hm.Remove(key)
			delete(ref, key)
		}
	}
}

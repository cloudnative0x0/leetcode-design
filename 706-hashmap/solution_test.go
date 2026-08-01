package _706_hashmap

import (
	"testing"
)

func TestMyHashMap_BaseScenario(t *testing.T) {
	hashMap := Constructor()

	hashMap.Put(1, 1)
	hashMap.Put(2, 2)

	if val := hashMap.Get(1); val != 1 {
		t.Errorf("Expected Get(1) to return 1, got %d", val)
	}
	if val := hashMap.Get(3); val != -1 {
		t.Errorf("Expected Get(3) to return -1 (not found), got %d", val)
	}

	hashMap.Put(2, 1)
	if val := hashMap.Get(2); val != 1 {
		t.Errorf("Expected Get(2) to return updated value 1, got %d", val)
	}

	hashMap.Remove(2)
	if val := hashMap.Get(2); val != -1 {
		t.Errorf("Expected Get(2) to return -1 after removal, got %d", val)
	}
}

func TestMyHashMap_Collisions(t *testing.T) {
	hashMap := Constructor()

	key1 := 10
	key2 := 10 + size

	hashMap.Put(key1, 100)
	hashMap.Put(key2, 200)

	if val := hashMap.Get(key1); val != 100 {
		t.Errorf("Expected Get(key1) = 100, got %d", val)
	}
	if val := hashMap.Get(key2); val != 200 {
		t.Errorf("Expected Get(key2) = 200, got %d", val)
	}

	hashMap.Remove(key1)
	if val := hashMap.Get(key1); val != -1 {
		t.Errorf("Expected key1 to be removed, got %d", val)
	}

	if val := hashMap.Get(key2); val != 200 {
		t.Errorf("Expected key2 to still exist with value 200, got %d", val)
	}
}

func TestMyHashMap_RemoveFromChain(t *testing.T) {
	hashMap := Constructor()

	k1 := 5
	k2 := 5 + size
	k3 := 5 + (size * 2)

	hashMap.Put(k1, 50)
	hashMap.Put(k2, 60)
	hashMap.Put(k3, 70)

	hashMap.Remove(k2)

	if val := hashMap.Get(k2); val != -1 {
		t.Errorf("Expected k2 to be removed, got %d", val)
	}
	if val := hashMap.Get(k1); val != 50 {
		t.Errorf("Expected k1 to remain, got %d", val)
	}
	if val := hashMap.Get(k3); val != 70 {
		t.Errorf("Expected k3 to remain, got %d", val)
	}
}

func TestMyHashMap_RemoveNonExistent(t *testing.T) {
	hashMap := Constructor()

	assertNoPanic(t, func() { hashMap.Remove(999) })

	hashMap.Put(1, 10)
	assertNoPanic(t, func() { hashMap.Remove(1 + size) })
}

func assertNoPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Function panicked: %v", r)
		}
	}()
	f()
}

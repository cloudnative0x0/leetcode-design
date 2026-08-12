package _705_hash_set

import "testing"

func TestAddAndContains(t *testing.T) {
	mhs := Constructor()

	if mhs.Contains(1) {
		t.Fatalf("Contains(1) = true, expected false on an empty set")
	}

	mhs.Add(1)
	mhs.Add(2)

	if !mhs.Contains(1) {
		t.Errorf("Contains(1) = false, expected true")
	}
	if !mhs.Contains(2) {
		t.Errorf("Contains(2) = false, expected true")
	}
	if mhs.Contains(3) {
		t.Errorf("Contains(3) = true, expected false")
	}
}

func TestAddDuplicateDoesNotBreakChain(t *testing.T) {
	mhs := Constructor()

	mhs.Add(1)
	mhs.Add(1)

	if !mhs.Contains(1) {
		t.Fatalf("Contains(1) = false after double Add(1), expected true")
	}

	index := mhs.hash(1)
	node := mhs.buckets[index]

	count := 0
	for node != nil {
		if node.key == 1 {
			count++
		}
		node = node.next
	}

	if count != 1 {
		t.Errorf("key 1 appears in the chain %d times, expected 1", count)
	}
}

func TestRemoveExisting(t *testing.T) {
	mhs := Constructor()

	mhs.Add(1)
	mhs.Remove(1)

	if mhs.Contains(1) {
		t.Errorf("Contains(1) = true after Remove(1), expected false")
	}
}

func TestRemoveNonExistingOnEmptySet(t *testing.T) {
	mhs := Constructor()

	// must not panic when removing from an empty bucket
	mhs.Remove(42)

	if mhs.Contains(42) {
		t.Errorf("Contains(42) = true, expected false")
	}
}

func TestRemoveNonExistingKeyInNonEmptyBucket(t *testing.T) {
	mhs := Constructor()

	mhs.Add(1)
	mhs.Remove(2) // key 2 does not exist, bucket for 1 is not empty

	if !mhs.Contains(1) {
		t.Errorf("Contains(1) = false, Remove(2) should not have affected key 1")
	}
}

func TestRemoveHeadOfChain(t *testing.T) {
	mhs := Constructor()
	bucketCount := mhs.bucketCount

	// keys land in the same bucket: k and k+bucketCount hash to the same value
	first := 5
	second := 5 + bucketCount

	mhs.Add(first)
	mhs.Add(second)

	mhs.Remove(first) // removes the head of the chain

	if mhs.Contains(first) {
		t.Errorf("Contains(%d) = true after Remove, expected false", first)
	}
	if !mhs.Contains(second) {
		t.Errorf("Contains(%d) = false, the second chain element should not have been affected", second)
	}
}

func TestRemoveTailOfChain(t *testing.T) {
	mhs := Constructor()
	bucketCount := mhs.bucketCount

	first := 5
	second := 5 + bucketCount
	third := 5 + 2*bucketCount

	mhs.Add(first)
	mhs.Add(second)
	mhs.Add(third)

	mhs.Remove(third) // removes the tail of the chain

	if !mhs.Contains(first) {
		t.Errorf("Contains(%d) = false, expected true", first)
	}
	if !mhs.Contains(second) {
		t.Errorf("Contains(%d) = false, expected true", second)
	}
	if mhs.Contains(third) {
		t.Errorf("Contains(%d) = true after Remove, expected false", third)
	}
}

func TestRemoveMiddleOfChain(t *testing.T) {
	mhs := Constructor()
	bucketCount := mhs.bucketCount

	first := 5
	second := 5 + bucketCount
	third := 5 + 2*bucketCount

	mhs.Add(first)
	mhs.Add(second)
	mhs.Add(third)

	mhs.Remove(second) // removes the middle node

	if !mhs.Contains(first) {
		t.Errorf("Contains(%d) = false, expected true", first)
	}
	if mhs.Contains(second) {
		t.Errorf("Contains(%d) = true after Remove, expected false", second)
	}
	if !mhs.Contains(third) {
		t.Errorf("Contains(%d) = false, expected true", third)
	}

	// verify the chain is not broken: first points directly to third
	index := mhs.hash(first)
	node := mhs.buckets[index]
	if node == nil || node.key != first {
		t.Fatalf("chain head is corrupted")
	}
	if node.next == nil || node.next.key != third {
		t.Errorf("after removing the middle node, head.next.key = %v, expected %d", node.next, third)
	}
}

func TestNegativeKeys(t *testing.T) {
	mhs := Constructor()

	mhs.Add(-1)
	mhs.Add(-100)

	if !mhs.Contains(-1) {
		t.Errorf("Contains(-1) = false, expected true")
	}
	if !mhs.Contains(-100) {
		t.Errorf("Contains(-100) = false, expected true")
	}

	mhs.Remove(-1)
	if mhs.Contains(-1) {
		t.Errorf("Contains(-1) = true after Remove(-1), expected false")
	}
	if !mhs.Contains(-100) {
		t.Errorf("Contains(-100) = false, Remove(-1) should not have affected it")
	}
}

func TestHashNeverNegative(t *testing.T) {
	mhs := Constructor()

	keys := []int{-1, -10007, -20014, -1000000000, 0, 10007, 20014}
	for _, k := range keys {
		h := mhs.hash(k)
		if h < 0 || h >= mhs.bucketCount {
			t.Errorf("hash(%d) = %d, expected range [0, %d)", k, h, mhs.bucketCount)
		}
	}
}

func TestCollisionBothKeysCoexist(t *testing.T) {
	mhs := Constructor()
	bucketCount := mhs.bucketCount

	a := 3
	b := 3 + bucketCount // same hash as a

	if mhs.hash(a) != mhs.hash(b) {
		t.Fatalf("test keys do not collide: hash(%d)=%d, hash(%d)=%d", a, mhs.hash(a), b, mhs.hash(b))
	}

	mhs.Add(a)
	mhs.Add(b)

	if !mhs.Contains(a) {
		t.Errorf("Contains(%d) = false, expected true", a)
	}
	if !mhs.Contains(b) {
		t.Errorf("Contains(%d) = false, expected true", b)
	}
}

func TestZeroKey(t *testing.T) {
	mhs := Constructor()

	mhs.Add(0)
	if !mhs.Contains(0) {
		t.Errorf("Contains(0) = false, expected true")
	}

	mhs.Remove(0)
	if mhs.Contains(0) {
		t.Errorf("Contains(0) = true after Remove(0), expected false")
	}
}

func TestAddRemoveAddSameKey(t *testing.T) {
	mhs := Constructor()

	mhs.Add(7)
	mhs.Remove(7)
	mhs.Add(7)

	if !mhs.Contains(7) {
		t.Errorf("Contains(7) = false after re-adding, expected true")
	}
}

func TestMultipleIndependentKeysInDifferentBuckets(t *testing.T) {
	mhs := Constructor()

	keys := []int{1, 2, 3, 4, 5, 100, 1000, 9999}
	for _, k := range keys {
		mhs.Add(k)
	}

	for _, k := range keys {
		if !mhs.Contains(k) {
			t.Errorf("Contains(%d) = false, expected true", k)
		}
	}

	mhs.Remove(3)

	for _, k := range keys {
		want := k != 3
		if got := mhs.Contains(k); got != want {
			t.Errorf("Contains(%d) = %v, expected %v", k, got, want)
		}
	}
}

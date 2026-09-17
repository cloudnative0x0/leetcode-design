package _380_insert_delete_get_random_o1

import "testing"

func TestConstructorCreatesEmptySet(t *testing.T) {
	rs := Constructor()

	if rs.indexMap == nil {
		t.Fatal("Constructor() left indexMap nil")
	}
	if rs.elements == nil {
		t.Fatal("Constructor() left elements nil")
	}
	if len(rs.indexMap) != 0 || len(rs.elements) != 0 {
		t.Fatalf("Constructor() created a non-empty set: map=%v, elements=%v", rs.indexMap, rs.elements)
	}
}

func TestInsert(t *testing.T) {
	rs := Constructor()

	if !rs.Insert(10) {
		t.Fatal("Insert(10) = false; want true for a new value")
	}
	if !rs.Insert(-3) {
		t.Fatal("Insert(-3) = false; want true for a new value")
	}

	assertStateIsConsistent(t, &rs)
	assertContainsExactly(t, &rs, 10, -3)
}

func TestInsertDuplicate(t *testing.T) {
	rs := Constructor()

	if !rs.Insert(7) {
		t.Fatal("first Insert(7) = false; want true")
	}
	if rs.Insert(7) {
		t.Fatal("second Insert(7) = true; want false")
	}

	assertStateIsConsistent(t, &rs)
	assertContainsExactly(t, &rs, 7)
}

func TestRemoveMissingValue(t *testing.T) {
	rs := Constructor()
	rs.Insert(1)

	if rs.Remove(2) {
		t.Fatal("Remove(2) = true; want false for a missing value")
	}

	assertStateIsConsistent(t, &rs)
	assertContainsExactly(t, &rs, 1)
}

func TestRemoveOnlyValue(t *testing.T) {
	rs := Constructor()
	rs.Insert(42)

	if !rs.Remove(42) {
		t.Fatal("Remove(42) = false; want true")
	}
	if len(rs.indexMap) != 0 || len(rs.elements) != 0 {
		t.Fatalf("set is not empty after removing its only value: map=%v, elements=%v", rs.indexMap, rs.elements)
	}
}

func TestRemoveMiddleMovesLastValue(t *testing.T) {
	rs := Constructor()
	rs.Insert(10)
	rs.Insert(20)
	rs.Insert(30)

	if !rs.Remove(20) {
		t.Fatal("Remove(20) = false; want true")
	}

	assertStateIsConsistent(t, &rs)
	assertContainsExactly(t, &rs, 10, 30)

	if got := rs.indexMap[30]; got != 1 {
		t.Errorf("indexMap[30] = %d; want 1 after moving the last value", got)
	}
}

func TestRemoveLastValue(t *testing.T) {
	rs := Constructor()
	rs.Insert(10)
	rs.Insert(20)
	rs.Insert(30)

	if !rs.Remove(30) {
		t.Fatal("Remove(30) = false; want true")
	}

	assertStateIsConsistent(t, &rs)
	assertContainsExactly(t, &rs, 10, 20)
}

func TestRemoveThenInsertSameValue(t *testing.T) {
	rs := Constructor()
	rs.Insert(5)
	rs.Remove(5)

	if !rs.Insert(5) {
		t.Fatal("Insert(5) = false after the value was removed; want true")
	}

	assertStateIsConsistent(t, &rs)
	assertContainsExactly(t, &rs, 5)
}

func TestGetRandomFromSingleton(t *testing.T) {
	rs := Constructor()
	rs.Insert(-11)

	for range 20 {
		if got := rs.GetRandom(); got != -11 {
			t.Fatalf("GetRandom() = %d; want -11", got)
		}
	}
}

func TestGetRandomReturnsStoredValue(t *testing.T) {
	rs := Constructor()
	values := []int{-10, 0, 8, 99}
	for _, value := range values {
		rs.Insert(value)
	}

	for range 200 {
		got := rs.GetRandom()
		if _, ok := rs.indexMap[got]; !ok {
			t.Fatalf("GetRandom() returned %d, which is not stored in the set", got)
		}
	}
}

func TestLeetCodeExample(t *testing.T) {
	rs := Constructor()

	if !rs.Insert(1) {
		t.Error("Insert(1) = false; want true")
	}
	if rs.Remove(2) {
		t.Error("Remove(2) = true; want false")
	}
	if !rs.Insert(2) {
		t.Error("Insert(2) = false; want true")
	}
	if got := rs.GetRandom(); got != 1 && got != 2 {
		t.Errorf("GetRandom() = %d; want 1 or 2", got)
	}
	if !rs.Remove(1) {
		t.Error("Remove(1) = false; want true")
	}
	if rs.Insert(2) {
		t.Error("Insert(2) = true for a duplicate; want false")
	}
	if got := rs.GetRandom(); got != 2 {
		t.Errorf("GetRandom() = %d; want 2", got)
	}

	assertStateIsConsistent(t, &rs)
}

func assertStateIsConsistent(t *testing.T, rs *RandomizedSet) {
	t.Helper()

	if len(rs.indexMap) != len(rs.elements) {
		t.Fatalf("len(indexMap) = %d, len(elements) = %d", len(rs.indexMap), len(rs.elements))
	}

	for index, value := range rs.elements {
		mappedIndex, ok := rs.indexMap[value]
		if !ok {
			t.Fatalf("elements[%d] = %d, but the value is missing from indexMap", index, value)
		}
		if mappedIndex != index {
			t.Fatalf("indexMap[%d] = %d; want %d", value, mappedIndex, index)
		}
	}
}

func assertContainsExactly(t *testing.T, rs *RandomizedSet, want ...int) {
	t.Helper()

	if len(rs.elements) != len(want) {
		t.Fatalf("elements = %v; want %v in any order", rs.elements, want)
	}

	for _, value := range want {
		if _, ok := rs.indexMap[value]; !ok {
			t.Errorf("value %d is missing; state=%v", value, rs.elements)
		}
	}
}

# 1603 · Design Parking System

**Difficulty:** Easy | **Time:** O(1) per call | **Space:** O(1)

---

## Solution keeps three counters of free slots

There's no need to model the slots themselves — only how many of each type remain free. Three fields, `slotBig`, `slotMedium`, `slotSmall`, get decremented when a car parks and are never allowed to go negative.

```go
type ParkingSystem struct {
    slotBig    int
    slotMedium int
    slotSmall  int
}
```

## Why counters instead of an array/slice of slots

- The problem never asks which exact slot is taken — only whether a free one exists.
- An array of slots would give the same correctness guarantee but cost O(n) memory and an extra scan on every attempt to park.
- Three ints solve it with constant memory and constant time per operation.

## Constructor

```go
func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{
        slotBig:    big,
        slotMedium: medium,
        slotSmall:  small,
    }
}
```

Just copies the starting capacities into the struct's fields.

## AddCar

```go
func (ps *ParkingSystem) AddCar(carType int) bool {
    switch carType {
    case 1:
        if ps.slotBig > 0 {
            ps.slotBig--
            return true
        }
    case 2:
        if ps.slotMedium > 0 {
            ps.slotMedium--
            return true
        }
    case 3:
        if ps.slotSmall > 0 {
            ps.slotSmall--
            return true
        }
    }

    return false
}
```

- `carType` maps directly onto one of the three counters (1 — big, 2 — medium, 3 — small), so a plain `switch` is enough, no extra lookup needed.
- If a slot is free, the counter is decremented and the function returns `true`.
- If none is free, the counter is left untouched and the function returns `false`.

**Edge cases:**
- A `carType` outside the 1–3 range doesn't match any `case` and falls through to `return false` — no panic.
- A capacity of 0 for a given type just always returns `false` for it from the start, no special handling required.
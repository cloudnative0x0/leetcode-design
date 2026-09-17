# 380 · Insert Delete GetRandom O(1)

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

**Difficulty:** Medium | **Time:** O(1) average | **Space:** O(n)

---

## Русский

### Решение через карту индексов и плотный срез
Значения лежат в обычном срезе без пустых ячеек, а карта хранит позицию каждого значения в этом срезе. Карта даёт быстрый поиск, срез — быстрый доступ по случайному индексу.

Удаление из середины не сдвигает хвост. На освободившееся место переносится последний элемент, после чего длина среза уменьшается на один.

---

### Структура

```go
type RandomizedSet struct {
    indexMap map[int]int
    elements []int
}
```

- `elements[i]` — значение, которое занимает позицию `i`
- `indexMap[val]` — текущая позиция `val` в `elements`
- длины карты и среза всегда совпадают, потому что дубликаты не хранятся

> Срез отвечает за выбор элемента, карта — за поиск его позиции. По отдельности ни одна из этих структур не даёт все три операции за O(1).

---

### Главный инвариант

Для каждого индекса `i` должно выполняться:

```go
indexMap[elements[i]] == i
```

Если в множестве хранится $n$ значений, то допустимые индексы образуют плотный диапазон:

$$
0 \le i < n
$$

Пустых мест внутри `elements` нет. Благодаря этому любой индекс, который возвращает `rand.N(n)`, указывает на существующий элемент.

---

### Почему удаление не сдвигает хвост

Обычное удаление `elements[idx]` из среза потребовало бы сдвинуть все элементы справа от `idx`. После такого сдвига пришлось бы обновить их позиции и в `indexMap`, поэтому стоимость операции была бы O(n).

Здесь порядок элементов не имеет значения. Значит, удаляемую ячейку можно занять последним значением:

```go
lastIdx := len(rs.elements) - 1
lastVal := rs.elements[lastIdx]

rs.elements[idx] = lastVal
rs.indexMap[lastVal] = idx
rs.elements = rs.elements[:lastIdx]
delete(rs.indexMap, val)
```

Меняются только одна ячейка среза и одна запись карты. Если удаляется последний элемент, `lastVal == val`: запись сначала обновляется тем же индексом, а затем удаляется. Отдельная ветка для этого случая не нужна.

---

### Равномерный случайный выбор

```go
randomIndex := rand.N(len(rs.elements))
return rs.elements[randomIndex]
```

`rand.N(n)` возвращает индекс из диапазона $[0, n)$. Поскольку каждое значение занимает ровно одну позицию, вероятность выбрать конкретное значение равна:

$$
P(x) = \frac{1}{n}
$$

Метод вызывается только для непустого множества — это гарантируется условием задачи. Самому типу не требуется дополнительная проверка или специальное значение для пустого состояния.

---

### Insert

```go
func (rs *RandomizedSet) Insert(val int) bool {
    if _, ok := rs.indexMap[val]; ok {
        return false
    }

    rs.indexMap[val] = len(rs.elements)
    rs.elements = append(rs.elements, val)
    return true
}
```

Сначала карта отсекает дубликат. Для нового значения запоминается индекс, равный прежней длине среза, затем значение добавляется в конец. Поиск в карте занимает O(1) в среднем, `append` — O(1) амортизированно.

---

### Remove

```go
func (rs *RandomizedSet) Remove(val int) bool {
    idx, ok := rs.indexMap[val]
    if !ok {
        return false
    }
    ...
}
```

Карта сразу даёт индекс удаляемого значения. Последний элемент переносится на этот индекс, его запись в карте исправляется, срез укорачивается, а запись `val` удаляется. Ни обхода, ни сдвига среза нет.

---

### GetRandom

```go
func (rs *RandomizedSet) GetRandom() int {
    randomIndex := rand.N(len(rs.elements))
    return rs.elements[randomIndex]
}
```

Случайный индекс вычисляется один раз, обращение к элементу среза по индексу также занимает O(1).

---

### Заметки по stdlib

- `map[int]int` хранит именно индексы, а не признаки наличия: одна структура одновременно проверяет существование значения и находит его для удаления
- `math/rand/v2` предоставляет обобщённую функцию `rand.N`, поэтому преобразовывать длину среза к другому числовому типу не требуется
- порядок в `elements` намеренно нестабилен: `Remove` может переставить последний элемент, но порядок не входит в контракт множества

---

> Вся конструкция держится на одном правиле: карта всегда знает настоящий индекс каждого элемента среза.

---

## English

### Solution uses an index map and a dense slice
Values live in an ordinary slice with no empty slots, while the map stores the position of every value in that slice. The map provides fast lookup; the slice provides fast access by a random index.

Removing from the middle does not shift the tail. The last element moves into the freed slot, then the slice is shortened by one.

---

### Architecture

```go
type RandomizedSet struct {
    indexMap map[int]int
    elements []int
}
```

- `elements[i]` — the value occupying position `i`
- `indexMap[val]` — the current position of `val` in `elements`
- the map and slice always have the same length because duplicates are never stored

> The slice handles selection; the map handles position lookup. Neither structure alone supports all three operations in O(1).

---

### Core invariant

For every index `i`, the following must hold:

```go
indexMap[elements[i]] == i
```

If the set contains $n$ values, its valid indices form one dense range:

$$
0 \le i < n
$$

There are no holes in `elements`. Therefore, every index returned by `rand.N(n)` points to an existing value.

---

### Why removal does not shift the tail

Deleting `elements[idx]` from a slice in the usual way would shift every element to the right of `idx`. Their entries in `indexMap` would then need updating as well, making the operation O(n).

Element order does not matter here, so the last value can occupy the slot being removed:

```go
lastIdx := len(rs.elements) - 1
lastVal := rs.elements[lastIdx]

rs.elements[idx] = lastVal
rs.indexMap[lastVal] = idx
rs.elements = rs.elements[:lastIdx]
delete(rs.indexMap, val)
```

Only one slice slot and one map entry change. When the last element itself is removed, `lastVal == val`: its entry is first assigned the same index and then deleted. No special branch is needed.

---

### Uniform random selection

```go
randomIndex := rand.N(len(rs.elements))
return rs.elements[randomIndex]
```

`rand.N(n)` returns an index in $[0, n)$. Since every value occupies exactly one position, the probability of selecting any particular value is:

$$
P(x) = \frac{1}{n}
$$

The method is called only when the set is non-empty, as guaranteed by the problem statement. The type does not need an extra check or a sentinel value for the empty state.

---

### Insert

```go
func (rs *RandomizedSet) Insert(val int) bool {
    if _, ok := rs.indexMap[val]; ok {
        return false
    }

    rs.indexMap[val] = len(rs.elements)
    rs.elements = append(rs.elements, val)
    return true
}
```

The map rejects a duplicate first. For a new value, the previous slice length is stored as its index, then the value is appended. Map lookup is O(1) on average; `append` is amortized O(1).

---

### Remove

```go
func (rs *RandomizedSet) Remove(val int) bool {
    idx, ok := rs.indexMap[val]
    if !ok {
        return false
    }
    ...
}
```

The map gives the removed value's index immediately. The last element moves to that index, its map entry is corrected, the slice is shortened, and the `val` entry is deleted. There is no scan and no slice shift.

---

### GetRandom

```go
func (rs *RandomizedSet) GetRandom() int {
    randomIndex := rand.N(len(rs.elements))
    return rs.elements[randomIndex]
}
```

The random index is generated once, and indexing a slice is also O(1).

---

### stdlib notes

- `map[int]int` stores indices rather than presence flags: one structure both checks membership and locates a value for removal
- `math/rand/v2` provides the generic `rand.N` function, so the slice length needs no conversion to another numeric type
- order in `elements` is deliberately unstable: `Remove` may relocate the last element, but ordering is not part of a set's contract

---

> The entire design rests on one rule: the map always knows the actual index of every slice element.

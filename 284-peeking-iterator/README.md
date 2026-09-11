# 284 · PeekingIterator

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

## Русский

### Идея

`Iterator` позволяет получить следующий элемент, но обычный вызов `next()` сразу сдвигает позицию. Задача `PeekingIterator` — добавить метод `peek()`, который показывает следующий элемент без его удаления из последовательности.

В решении для этого используется буфер `arr`. Метод `peek()` забирает значение у исходного итератора и сохраняет его. Последующий `next()` сначала проверяет буфер и возвращает сохранённое значение.

---

### Базовый Iterator

На LeetCode тип `Iterator` уже определён. Для локального запуска он хранит исходные значения и индекс следующего элемента:

```go
type Iterator struct {
    values []int
    index  int
}

func NewIterator(values []int) *Iterator {
    return &Iterator{values: values}
}

func (it *Iterator) hasNext() bool {
    return it.index < len(it.values)
}

func (it *Iterator) next() int {
    value := it.values[it.index]
    it.index++
    return value
}
```

`hasNext()` проверяет, остались ли непрочитанные элементы. `next()` возвращает элемент с текущим индексом и переводит индекс на следующую позицию.

---

### Структура PeekingIterator

```go
type PeekingIterator struct {
    iter *Iterator
    arr  []int
}
```

- `iter` — исходный итератор;
- `arr` — буфер для элемента, который уже прочитан через `peek()`, но ещё не выдан через `next()`.

При таком порядке вызовов в `arr` находится не больше одного элемента.

---

### Constructor

```go
func Constructor(iter *Iterator) *PeekingIterator {
    return &PeekingIterator{
        iter: iter,
        arr:  make([]int, 0),
    }
}
```

Конструктор сохраняет ссылку на исходный итератор и создаёт пустой буфер. Чтение элементов начинается только при вызове `next()` или `peek()`.

---

### hasNext

```go
func (pi *PeekingIterator) hasNext() bool {
    if len(pi.arr) > 0 {
        return true
    }

    return pi.iter.hasNext()
}
```

Сначала проверяется буфер. Если в нём есть значение, следующий элемент доступен независимо от состояния `iter`. При пустом буфере проверка передаётся исходному итератору.

Этот порядок важен для последнего элемента: после `peek()` исходный итератор уже пуст, но сохранённое значение ещё можно получить через `next()`.

---

### next

```go
func (pi *PeekingIterator) next() int {
    if len(pi.arr) > 0 {
        v := pi.arr[0]
        pi.arr = pi.arr[1:]
        return v
    }

    return pi.iter.next()
}
```

Если `peek()` уже сохранил элемент, `next()` возвращает его из `arr` и очищает буфер. В остальных случаях метод напрямую вызывает `iter.next()`.

---

### peek

```go
func (pi *PeekingIterator) peek() int {
    if len(pi.arr) > 0 {
        return pi.arr[0]
    }

    pi.arr = append(pi.arr, pi.iter.next())
    return pi.arr[0]
}
```

При первом вызове `peek()` следующий элемент извлекается из `iter` и помещается в `arr`. Повторный вызов возвращает `arr[0]`, поэтому позиция больше не меняется. Элемент удалит только `next()`.

---

### Пример работы буфера

Для последовательности `[1, 2, 3]`:

| Вызов | Результат | Буфер после вызова | Следующая позиция `iter` |
| --- | ---: | --- | ---: |
| `next()` | `1` | `[]` | `2` |
| `peek()` | `2` | `[2]` | `3` |
| `peek()` | `2` | `[2]` | `3` |
| `next()` | `2` | `[]` | `3` |
| `next()` | `3` | `[]` | конец |

---

### Почему это работает

Буфер разделяет чтение элемента из `iter` и его выдачу через `next()`. `peek()` выполняет чтение заранее, а `next()` завершает операцию, возвращая сохранённое значение. Благодаря этому каждый элемент извлекается из исходного итератора один раз, а порядок последовательности не меняется.

---

### Сложность

| Операция | Время | Дополнительная память |
| --- | --- | --- |
| `Constructor` | O(1) | O(1) |
| `hasNext()` | O(1) | O(1) |
| `next()` | O(1) | O(1) |
| `peek()` | O(1) | O(1) |

Буфер содержит максимум один ожидающий элемент, поэтому общий расход дополнительной памяти остаётся O(1).

---

### Тесты

В тестах проверяются базовый `Iterator`, пример с последовательностью `[1, 2, 3]`, повторный `peek()`, последний элемент в буфере и сохранение исходного порядка.

```bash
GO111MODULE=off go test ./...
```

---

## English

### Idea

`Iterator` returns the next element, but a regular `next()` call advances its position immediately. `PeekingIterator` adds `peek()`, which shows the next element without removing it from the sequence.

The solution uses `arr` as a buffer. `peek()` reads a value from the underlying iterator and stores it. The following `next()` checks the buffer first and returns the stored value.

---

### Base Iterator

LeetCode provides the `Iterator` type. The local implementation stores the input values and the index of the next element:

```go
type Iterator struct {
    values []int
    index  int
}

func NewIterator(values []int) *Iterator {
    return &Iterator{values: values}
}

func (it *Iterator) hasNext() bool {
    return it.index < len(it.values)
}

func (it *Iterator) next() int {
    value := it.values[it.index]
    it.index++
    return value
}
```

`hasNext()` checks whether unread elements remain. `next()` returns the value at the current index and advances the index.

---

### PeekingIterator structure

```go
type PeekingIterator struct {
    iter *Iterator
    arr  []int
}
```

- `iter` — the underlying iterator;
- `arr` — a buffer for an element already read by `peek()` but not yet returned by `next()`.

With this call flow, `arr` contains at most one element.

---

### Constructor

```go
func Constructor(iter *Iterator) *PeekingIterator {
    return &PeekingIterator{
        iter: iter,
        arr:  make([]int, 0),
    }
}
```

The constructor stores the underlying iterator and creates an empty buffer. It does not read an element until `next()` or `peek()` is called.

---

### hasNext

```go
func (pi *PeekingIterator) hasNext() bool {
    if len(pi.arr) > 0 {
        return true
    }

    return pi.iter.hasNext()
}
```

The buffer is checked first. If it contains a value, another element is available regardless of the state of `iter`. With an empty buffer, the method delegates to the underlying iterator.

This order matters for the final element: after `peek()`, the underlying iterator is exhausted, but the buffered value is still available through `next()`.

---

### next

```go
func (pi *PeekingIterator) next() int {
    if len(pi.arr) > 0 {
        v := pi.arr[0]
        pi.arr = pi.arr[1:]
        return v
    }

    return pi.iter.next()
}
```

If `peek()` has stored an element, `next()` returns it from `arr` and clears the buffer. Otherwise, it calls `iter.next()` directly.

---

### peek

```go
func (pi *PeekingIterator) peek() int {
    if len(pi.arr) > 0 {
        return pi.arr[0]
    }

    pi.arr = append(pi.arr, pi.iter.next())
    return pi.arr[0]
}
```

On the first call, `peek()` takes the next element from `iter` and puts it into `arr`. Repeated calls return `arr[0]`, so the position does not move again. Only `next()` removes the element.

---

### Buffer example

For the sequence `[1, 2, 3]`:

| Call | Result | Buffer after the call | Next `iter` position |
| --- | ---: | --- | ---: |
| `next()` | `1` | `[]` | `2` |
| `peek()` | `2` | `[2]` | `3` |
| `peek()` | `2` | `[2]` | `3` |
| `next()` | `2` | `[]` | `3` |
| `next()` | `3` | `[]` | end |

---

### Why it works

The buffer separates reading an element from `iter` and returning it through `next()`. `peek()` performs the read in advance, while `next()` completes the operation by returning the stored value. Each element is read from the underlying iterator once, and the original order is preserved.

---

### Complexity

| Operation | Time | Extra space |
| --- | --- | --- |
| `Constructor` | O(1) | O(1) |
| `hasNext()` | O(1) | O(1) |
| `next()` | O(1) | O(1) |
| `peek()` | O(1) | O(1) |

The buffer holds at most one pending element, so total extra space remains O(1).

---

### Tests

The tests cover the base `Iterator`, the `[1, 2, 3]` example, repeated `peek()` calls, the last buffered element, and preservation of the original order.

```bash
GO111MODULE=off go test ./...
```
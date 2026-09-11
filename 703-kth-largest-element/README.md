# KthLargest

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

## Русский

### Идея

После каждого добавления числа нужно вернуть `k`-й по величине элемент всего потока. Хранить и сортировать все полученные значения необязательно: для ответа достаточно знать `k` самых больших.

Решение хранит их в min-heap размером не больше `k`. Корень такой кучи — наименьший среди выбранных значений, то есть текущий `k`-й наибольший элемент.

---

### Реализация min-heap

```go
type ScoreHeap []int
```

`ScoreHeap` реализует интерфейс `heap.Interface` из стандартного пакета `container/heap`:

```go
func (h ScoreHeap) Len() int {
    return len(h)
}

func (h ScoreHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h ScoreHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
```

Сравнение `h[i] < h[j]` помещает минимальный элемент в корень кучи.

Методы `Push` и `Pop` изменяют срез через указатель:

```go
func (h *ScoreHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *ScoreHeap) Pop() any {
    current := *h
    n := len(current)
    x := current[n-1]
    *h = current[0 : n-1]

    return x
}
```

Перед вызовом `Pop()` пакет `heap` переставляет удаляемый корень в конец среза. Поэтому метод забирает последний элемент.

Для чтения корня добавлен отдельный метод:

```go
func (h *ScoreHeap) Peek() int {
    return (*h)[0]
}
```

---

### Структура KthLargest

```go
type KthLargest struct {
    h *ScoreHeap
    k int
}
```

- `h` хранит `k` наибольших элементов потока;
- `k` определяет требуемую позицию.

Если в куче уже находится `k` элементов, её корень является ответом.

---

### Constructor

```go
func Constructor(k int, nums []int) KthLargest {
    startingScores := ScoreHeap(nums)

    o := KthLargest{
        h: &startingScores,
        k: k,
    }

    heap.Init(o.h)

    for range len(nums) - k {
        heap.Pop(o.h)
    }

    return o
}
```

Сначала `nums` преобразуется в `ScoreHeap`, после чего `heap.Init` перестраивает элементы в min-heap.

Если элементов больше `k`, минимальные значения удаляются. После `len(nums) - k` удалений в куче остаются только `k` наибольших элементов исходного массива.

Запись `for range len(nums) - k` требует Go 1.22 или новее.

---

### Add

#### В куче меньше k элементов

```go
if kl.h.Len() < kl.k {
    heap.Push(kl.h, val)
    return kl.h.Peek()
}
```

Новое значение добавляется без сравнения: куча ещё не набрала требуемый размер.

По условию задачи начальный массив содержит не меньше `k - 1` элементов. Поэтому после первого вызова `Add` в потоке уже будет достаточно значений для определения `k`-го наибольшего элемента.

#### Новое значение больше корня

```go
current := kl.h.Peek()

if val > current {
    (*(kl.h))[0] = val
    heap.Fix(kl.h, 0)

    current = kl.h.Peek()
}
```

Корень — наименьший элемент среди текущих `k` наибольших. Если `val` больше корня, старый корень больше не входит в нужную группу.

Он заменяется новым значением, а `heap.Fix` восстанавливает порядок кучи после изменения корневого элемента.

#### Новое значение не больше корня

Если `val <= current`, значение не может попасть в `k` наибольших. Куча остаётся без изменений, а метод возвращает прежний корень.

---

### Пример

Для `k = 3` и `nums = [4, 5, 8, 2]` после конструктора куча логически содержит три наибольших значения: `4`, `5`, `8`. Корнем является `4`.

| Вызов | Значения в группе top k | Результат |
| --- | --- | ---: |
| `Add(3)` | `4, 5, 8` | `4` |
| `Add(5)` | `5, 5, 8` | `5` |
| `Add(10)` | `5, 8, 10` | `5` |
| `Add(9)` | `8, 9, 10` | `8` |
| `Add(4)` | `8, 9, 10` | `8` |

Порядок элементов внутри кучи может отличаться от порядка в таблице. Для решения важны корень и соблюдение свойства min-heap, а не полная сортировка.

---

### Почему это работает

Куча сохраняет инвариант: после обработки очередного значения в ней находятся `k` наибольших элементов потока.

Минимальный элемент этой группы расположен в корне. В общем порядке он занимает `k`-ю позицию по убыванию.

Значения меньше или равные корню можно пропустить. Значения больше корня заменяют его и входят в группу `k` наибольших.

---

### Сложность

| Операция | Время | Хранимые элементы |
| --- | --- | --- |
| `Constructor` | O(n + (n - k) log n) | O(k) |
| `Add` | O(log k) | O(k) |
| `Peek` | O(1) | O(1) |

`heap.Init` работает за O(n). Затем конструктор удаляет `n - k` минимальных элементов.

Вызов `Add` либо ничего не меняет, либо восстанавливает кучу высотой O(log k).

---

### Детали реализации

- `ScoreHeap(nums)` использует тот же базовый массив, что и входной срез, поэтому конструктор может изменить порядок элементов в `nums`;
- полная сортировка не выполняется;
- одинаковые значения обрабатываются как отдельные элементы потока;
- по условию задачи к моменту возврата результата в потоке есть не меньше `k` элементов.

---

## English

### Idea

After each inserted number, the service must return the `k`th largest value in the complete stream. There is no need to store and sort every received number: the answer depends only on the `k` largest values.

The solution keeps them in a min-heap containing at most `k` elements. Its root is the smallest value among the selected elements, which makes it the current `k`th largest value.

---

### Min-heap implementation

```go
type ScoreHeap []int
```

`ScoreHeap` implements `heap.Interface` from the standard `container/heap` package:

```go
func (h ScoreHeap) Len() int {
    return len(h)
}

func (h ScoreHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h ScoreHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
```

The `h[i] < h[j]` comparison places the minimum value at the heap root.

`Push` and `Pop` modify the slice through a pointer:

```go
func (h *ScoreHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *ScoreHeap) Pop() any {
    current := *h
    n := len(current)
    x := current[n-1]
    *h = current[0 : n-1]

    return x
}
```

Before calling `Pop()`, the `heap` package moves the root being removed to the end of the slice. The method therefore removes the last element.

The root can be read through a separate method:

```go
func (h *ScoreHeap) Peek() int {
    return (*h)[0]
}
```

---

### KthLargest structure

```go
type KthLargest struct {
    h *ScoreHeap
    k int
}
```

- `h` stores the `k` largest stream elements;
- `k` identifies the required position.

When the heap contains `k` elements, its root is the answer.

---

### Constructor

```go
func Constructor(k int, nums []int) KthLargest {
    startingScores := ScoreHeap(nums)

    o := KthLargest{
        h: &startingScores,
        k: k,
    }

    heap.Init(o.h)

    for range len(nums) - k {
        heap.Pop(o.h)
    }

    return o
}
```

First, `nums` is converted to `ScoreHeap`. `heap.Init` then rearranges the values into a min-heap.

If there are more than `k` elements, the minimum values are removed. After `len(nums) - k` removals, the heap contains only the `k` largest values from the initial array.

The `for range len(nums) - k` syntax requires Go 1.22 or newer.

---

### Add

#### The heap contains fewer than k elements

```go
if kl.h.Len() < kl.k {
    heap.Push(kl.h, val)
    return kl.h.Peek()
}
```

The new value is inserted without comparison because the heap has not reached its required size yet.

The problem guarantees that the initial array contains at least `k - 1` elements. Therefore, after the first `Add` call, the stream contains enough values to determine the `k`th largest element.

#### The new value is greater than the root

```go
current := kl.h.Peek()

if val > current {
    (*(kl.h))[0] = val
    heap.Fix(kl.h, 0)

    current = kl.h.Peek()
}
```

The root is the smallest of the current `k` largest values. When `val` is greater than the root, the old root no longer belongs to the required group.

It is replaced with the new value, and `heap.Fix` restores the heap order after the root has been changed.

#### The new value is not greater than the root

When `val <= current`, it cannot enter the `k` largest group. The heap remains unchanged, and the method returns the existing root.

---

### Example

For `k = 3` and `nums = [4, 5, 8, 2]`, the heap logically contains the three largest values after construction: `4`, `5`, and `8`. Its root is `4`.

| Call | Values in the top-k group | Result |
| --- | --- | ---: |
| `Add(3)` | `4, 5, 8` | `4` |
| `Add(5)` | `5, 5, 8` | `5` |
| `Add(10)` | `5, 8, 10` | `5` |
| `Add(9)` | `8, 9, 10` | `8` |
| `Add(4)` | `8, 9, 10` | `8` |

The internal heap order can differ from the order shown in the table. The solution needs the root and the min-heap property, not a fully sorted sequence.

---

### Why it works

The heap maintains the following invariant: after processing a value, it contains the `k` largest elements seen so far.

The smallest value in that group is at the root. In descending order, that value occupies position `k`.

Values smaller than or equal to the root can be ignored. Values greater than the root replace it and enter the `k` largest group.

---

### Complexity

| Operation | Time | Stored elements |
| --- | --- | --- |
| `Constructor` | O(n + (n - k) log n) | O(k) |
| `Add` | O(log k) | O(k) |
| `Peek` | O(1) | O(1) |

`heap.Init` takes O(n). The constructor then removes `n - k` minimum elements.

An `Add` call either leaves the heap unchanged or restores a heap with height O(log k).

---

### Implementation details

- `ScoreHeap(nums)` uses the same backing array as the input slice, so the constructor can change the order of values in `nums`;
- the values are never fully sorted;
- duplicate values are handled as separate stream elements;
- the problem guarantees that the stream contains at least `k` elements when a result is returned.
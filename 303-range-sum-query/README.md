# NumArray

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

## Русский

### Идея

Задача — по неизменяемому массиву `nums` отвечать на запросы суммы отрезка `SumRange(left, right)`. Массив не меняется между запросами, значит сумму каждого отрезка можно подготовить один раз, а не пересчитывать на каждый вызов.

---

### Подход

Строится массив префиксных сумм `prefix` длиной `len(nums) + 1`:

```
prefix[i] = сумма первых i элементов nums
```

`prefix[0] = 0` — база, задаётся автоматически: `make([]int, n)` в Go заполняет срез нулями.

---

### Структура

```go
type NumArray struct {
    prefix []int
}
```

* `prefix` — единственное поле, вся логика завязана на нём

---

### Constructor

```go
func Constructor(nums []int) NumArray {
    prefix := make([]int, len(nums)+1)

    for i := 0; i < len(nums); i++ {
        prefix[i+1] = prefix[i] + nums[i]
    }

    return NumArray{
        prefix: prefix,
    }
}
```

Размер `prefix` берётся из `len(nums)`, магических констант нет.

Каждый шаг цикла добавляет к уже накопленной сумме ровно один элемент — `prefix[i+1]` строится через `prefix[i]`, а не пересчитывается заново.

---

### SumRange

```go
func (this *NumArray) SumRange(left int, right int) int {
    return this.prefix[right+1] - this.prefix[left]
}
```

Сумма отрезка получается вычитанием двух готовых префиксов, без обхода `nums`.

---

### Почему это работает

`prefix[right+1]` — сумма всех элементов вплоть до `right` включительно. `prefix[left]` — сумма всех элементов до `left`, не включая его. Разница между ними — ровно сумма `nums[left..right]`.

Построение выполняется один раз в `Constructor`, каждый `SumRange` только читает готовые значения.

---

### Операции

| Операция      | Сложность | Описание                          |
| ------------- | --------- | ----------------------------------- |
| `Constructor` | O(n)      | построение префиксных сумм          |
| `SumRange`    | O(1)      | сумма отрезка через вычитание       |

n — длина `nums`.

---

### Детали реализации

* размер `prefix` берётся из `len(nums)`, а не задаётся заранее
* индексация сдвинута на 1, чтобы `prefix[0] = 0` покрывал случай `left == 0` без отдельной ветки
* получатель методов `*NumArray` — не копирует срез при каждом вызове

---

### Почему не считать сумму на каждом запросе

```go
func (this *NumArray) SumRangeBrute(left int, right int) int {
    sum := 0
    for i := left; i <= right; i++ {
        sum += this.nums[i]
    }
    return sum
}
```

Работает корректно, но O(n) на каждый вызов. При q запросах суммарная сложность становится O(n·q). С префиксными суммами — O(n) один раз плюс O(1) на каждый из q запросов, то есть O(n + q).

---

### Тесты

Файл `numarray_test.go` содержит:

* табличные тесты на известных примерах (в том числе пример из условия задачи), отрицательные числа, нули, один элемент
* стресс-тест: сравнение с brute force суммированием на 2000 случайных массивах, фиксированный seed для воспроизводимости
* бенчмарки отдельно для `Constructor` и `SumRange`

```
go test -v ./...
go test -run Stress -v
go test -bench . -v
```

---

### Ограничения

* `nums` должен оставаться неизменным после вызова `Constructor`; если элементы `nums` меняются снаружи, `prefix` не обновляется автоматически
* нет проверки границ `left`/`right`

---

## English

### Idea

The task is to answer range sum queries `SumRange(left, right)` on an immutable array `nums`. Since the array does not change between queries, the sum of any range can be prepared once instead of recomputed on every call.

---

### Approach

A prefix sum array `prefix` of length `len(nums) + 1` is built:

```
prefix[i] = sum of the first i elements of nums
```

`prefix[0] = 0` — the base case, set automatically: `make([]int, n)` in Go zero-fills the slice.

---

### Structure

```go
type NumArray struct {
    prefix []int
}
```

* `prefix` — the only field, all logic is built around it

---

### Constructor

```go
func Constructor(nums []int) NumArray {
    prefix := make([]int, len(nums)+1)

    for i := 0; i < len(nums); i++ {
        prefix[i+1] = prefix[i] + nums[i]
    }

    return NumArray{
        prefix: prefix,
    }
}
```

The size of `prefix` comes from `len(nums)`, no magic constants.

Each loop step adds exactly one element to the already accumulated sum — `prefix[i+1]` is built from `prefix[i]`, not recomputed from scratch.

---

### SumRange

```go
func (this *NumArray) SumRange(left int, right int) int {
    return this.prefix[right+1] - this.prefix[left]
}
```

The range sum comes from subtracting two already-built prefixes, without scanning `nums`.

---

### Why it works

`prefix[right+1]` is the sum of everything up to and including `right`. `prefix[left]` is the sum of everything before `left`. The difference between them is exactly the sum of `nums[left..right]`.

The build happens once in `Constructor`, every `SumRange` call only reads already computed values.

---

### Operations

| Operation     | Complexity | Description                     |
| ------------- | ---------- | --------------------------------- |
| `Constructor` | O(n)       | build the prefix sum array        |
| `SumRange`    | O(1)       | range sum via subtraction         |

n — length of `nums`.

---

### Implementation details

* `prefix` size comes from `len(nums)`, not set upfront
* indexing is shifted by 1 so `prefix[0] = 0` covers `left == 0` without a separate branch
* method receiver is `*NumArray` — the slice is not copied on every call

---

### Why not sum on every query

```go
func (this *NumArray) SumRangeBrute(left int, right int) int {
    sum := 0
    for i := left; i <= right; i++ {
        sum += this.nums[i]
    }
    return sum
}
```

Correct, but O(n) per call. With q queries, total cost becomes O(n·q). With prefix sums it is O(n) once plus O(1) per query, so O(n + q).

---

### Tests

`numarray_test.go` contains:

* table-driven tests on known examples (including the example from the problem statement), negative numbers, zeros, a single element
* stress test: comparison against brute force summation on 2000 random arrays, fixed seed for reproducibility
* separate benchmarks for `Constructor` and `SumRange`

```
go test -v ./...
go test -run Stress -v
go test -bench . -v
```

---

### Limitations

* `nums` must stay unchanged after `Constructor` is called; if `nums` is mutated externally, `prefix` does not update automatically
* no bounds check on `left`/`right`

---

<br>

> Сумма не собирается заново — она уже разложена по шагам.
>
> The sum is not rebuilt — it is already broken down step by step.
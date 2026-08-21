# 341 ·NestedIterator

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

## Русский

### Идея

Задача — превратить вложенный список `NestedList` (список, где элементом может быть либо число, либо снова список) в плоский поток чисел, но без предварительной сборки всего результата в память. Итератор должен отдавать числа по одному через `Next()`, а `HasNext()` — говорить, остался ли ещё хотя бы один элемент.

Наивный вариант — рекурсивно развернуть всё в `[]int` заранее в `Constructor`. Он работает, но если список огромный, а прочитать нужно только первые несколько элементов, вся остальная работа была сделана зря. Ниже — вариант, где раскрытие списков происходит лениво, по мере необходимости, через `stack`.

---

### Подход

`stack` — это срез срезов: `[][]*NestedInteger`. Каждый элемент `stack` — это "то, что осталось обработать" на одном уровне вложенности. Верхний элемент стека (последний в срезе) — самый глубокий уровень, с которым сейчас идёт работа.

Изначально в стеке лежит один уровень — сам входной список:

```
stack = [ nestedList ]
```

Дальше есть два действия:

* если на вершине стека лежит число — можно сразу его вернуть;
* если на вершине стека лежит вложенный список — этот список нужно "развернуть", то есть положить его содержимое новым уровнем поверх стека, а из текущего уровня этот элемент убрать.

Раскрытие идёт до тех пор, пока на вершине стека не окажется число или стек не опустеет полностью.

---

### Структура

```go
type NestedIterator struct {
    stack [][]*NestedInteger
}
```

* `stack` — единственное поле, это и есть вся память итератора; никакого заранее посчитанного плоского списка в структуре нет

---

### Constructor

```go
func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{
       stack: [][]*NestedInteger{nestedList},
    }
}
```

Ничего не раскрывается заранее. Входной список просто кладётся в стек как первый (и пока единственный) уровень. Вся дальнейшая работа откладывается до вызовов `HasNext()` — это и есть ленивость подхода.

---

### HasNext

```go
func (ni *NestedIterator) HasNext() bool {
    for len(ni.stack) > 0 {
       top := len(ni.stack) - 1

       if len(ni.stack[top]) == 0 {
          ni.stack = ni.stack[:top]
          continue
       }

       firstElem := ni.stack[top][0]

       if firstElem.IsInteger() {
          return true
       }

       ni.stack[top] = ni.stack[top][1:]

       ni.stack = append(ni.stack, firstElem.GetList())
    }

    return false
}
```

Это основная логика всего решения. Цикл на каждом шаге смотрит на вершину стека и разбирает три случая:

1. **Верхний уровень пуст** (`len(ni.stack[top]) == 0`) — на этом уровне обрабатывать больше нечего, он снимается со стека (`ni.stack = ni.stack[:top]`), и цикл переходит к уровню, который был под ним.

2. **Первый элемент верхнего уровня — число** (`firstElem.IsInteger()`) — раскрывать нечего, стек уже в состоянии "готов к чтению", функция возвращает `true`.

3. **Первый элемент верхнего уровня — вложенный список** — этот элемент убирается из текущего уровня (`ni.stack[top] = ni.stack[top][1:]`), а его содержимое кладётся новым уровнем поверх стека (`append(ni.stack, firstElem.GetList())`). Цикл повторяется уже с новой, более глубокой вершиной.

Цикл заканчивается либо возвратом `true` (на вершине число), либо когда стек полностью опустел — тогда `false`.

---

### Next

```go
func (ni *NestedIterator) Next() int {
    top := ni.stack[len(ni.stack)-1]

    elem := top[0]

    ni.stack[len(ni.stack)-1] = top[1:]

    return elem.GetInteger()
}
```

`Next()` ничего не разворачивает — вся эта работа уже сделана в `HasNext()`. Он предполагает, что на вершине стека прямо сейчас лежит число (контракт итератора: `HasNext()` вызывается перед `Next()`). Остаётся забрать первый элемент верхнего уровня, сдвинуть верхний уровень на один элемент вперёд (`top[1:]`) и вернуть значение через `GetInteger()`.

---

### Почему это работает

Стек всегда хранит путь от текущей позиции чтения до корня списка: каждый его уровень — это "хвост", который ещё не обработан на соответствующей глубине вложенности. Когда `HasNext()` доходит до числа, это гарантированно следующее число в порядке обхода — потому что все элементы перед ним на всех уровнях стека уже были либо возвращены через `Next()`, либо развёрнуты в более глубокие уровни.

Пустой уровень на вершине стека означает, что этот уровень вложенности полностью пройден, и обработка возвращается на уровень выше — ровно так работает обычный обход дерева в глубину, только без рекурсии: стек тут используется явно, вместо стека вызовов.

---

### Операции

| Операция    | Сложность              | Описание                                        |
| ----------- | ----------------------- | ------------------------------------------------ |
| `Constructor` | O(1)                   | список просто кладётся на стек, ничего не считается |
| `HasNext`   | амортизированно O(1)    | каждый элемент разворачивается стеком не более одного раза за всё время жизни итератора |
| `Next`      | O(1)                    | чтение и сдвиг верхнего уровня стека              |

Каждый элемент — число это или список — попадает в стек и снимается с него ровно один раз за весь обход, поэтому суммарная стоимость всех вызовов `HasNext()` за полный обход линейна относительно общего числа элементов (чисел и вложенных списков вместе), а не квадратична.

---

### Детали реализации

* уровень стека — это `[]*NestedInteger`, "хвост" необработанных элементов; продвижение по уровню — это отбрасывание головы среза (`top[1:]`), без явного индекса-счётчика
* пустые вложенные списки (`[]`) не требуют отдельной обработки: такой уровень просто добавляется в стек и на следующей итерации `HasNext()` тут же снимается как пустой
* `HasNext()` можно вызывать сколько угодно раз подряд без побочных эффектов — если стек уже приведён к состоянию "на вершине число", повторные вызовы просто сразу возвращают `true`
* получатель методов — `*NestedIterator`, стек изменяется на месте, без копирования на каждом вызове

---

### Почему не разворачивать всё заранее

```go
func ConstructorEager(nestedList []*NestedInteger) *NestedIterator {
    flat := make([]int, 0)
    var flatten func([]*NestedInteger)
    flatten = func(list []*NestedInteger) {
        for _, ni := range list {
            if ni.IsInteger() {
                flat = append(flat, ni.GetInteger())
            } else {
                flatten(ni.GetList())
            }
        }
    }
    flatten(nestedList)
    // ...дальше Next() просто читает flat по индексу
}
```

Тоже корректно, но вся работа по раскрытию списков выполняется в `Constructor`, даже если снаружи ни разу не вызовут `Next()`. Плюс требуется отдельная память под весь плоский список. Вариант со `stack` разворачивает ровно столько, сколько реально было прочитано — если обход прервать на середине, лишняя часть списка вообще не будет тронута.

---


### Ограничения

* `Next()` не проверяет, есть ли элемент — вызов без предварительного `HasNext() == true` приведёт к панике на пустом срезе
* `NestedList`, переданный в `Constructor`, не должен изменяться извне во время обхода — стек хранит срезы, указывающие на те же данные

---

## English

### Idea

The task is to turn a nested list `NestedList` (a list where each element is either an integer or another such list) into a flat stream of integers, without building the whole flattened result in memory upfront. The iterator has to yield numbers one at a time through `Next()`, while `HasNext()` reports whether at least one element is still left.

The naive approach is to recursively flatten everything into `[]int` inside `Constructor`. That works, but if the list is huge and only the first few elements are ever read, the rest of that work is wasted. Below is a version where lists are unpacked lazily, only as needed, using a `stack`.

---

### Approach

`stack` is a slice of slices: `[][]*NestedInteger`. Each entry in `stack` is "what's left to process" at one level of nesting. The top of the stack (the last entry in the slice) is the deepest level currently being worked on.

Initially the stack holds a single level — the input list itself:

```
stack = [ nestedList ]
```

From there, two things can happen:

* if an integer sits at the top of the stack, it can be returned right away;
* if a nested list sits at the top of the stack, it gets "unpacked" — its contents become a new level pushed on top of the stack, and that element is removed from the current level.

Unpacking keeps happening until either an integer is at the top of the stack, or the stack is empty.

---

### Structure

```go
type NestedIterator struct {
    stack [][]*NestedInteger
}
```

* `stack` — the only field, and it is the entire memory of the iterator; there is no pre-computed flat list stored anywhere in the struct

---

### Constructor

```go
func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{
       stack: [][]*NestedInteger{nestedList},
    }
}
```

Nothing gets unpacked upfront. The input list is simply placed on the stack as the first (and for now only) level. All the actual work is deferred until `HasNext()` is called — that's the laziness of the approach.

---

### HasNext

```go
func (ni *NestedIterator) HasNext() bool {
    for len(ni.stack) > 0 {
       top := len(ni.stack) - 1

       if len(ni.stack[top]) == 0 {
          ni.stack = ni.stack[:top]
          continue
       }

       firstElem := ni.stack[top][0]

       if firstElem.IsInteger() {
          return true
       }

       ni.stack[top] = ni.stack[top][1:]

       ni.stack = append(ni.stack, firstElem.GetList())
    }

    return false
}
```

This is the core logic of the whole solution. On each step, the loop looks at the top of the stack and handles three cases:

1. **The top level is empty** (`len(ni.stack[top]) == 0`) — nothing left to process at this level, so it gets popped off the stack (`ni.stack = ni.stack[:top]`), and the loop moves down to whatever level was underneath it.

2. **The first element of the top level is an integer** (`firstElem.IsInteger()`) — there is nothing to unpack, the stack is already in a "ready to read" state, so the function returns `true`.

3. **The first element of the top level is a nested list** — that element is removed from the current level (`ni.stack[top] = ni.stack[top][1:]`), and its contents are pushed as a new level on top of the stack (`append(ni.stack, firstElem.GetList())`). The loop repeats with a new, deeper top.

The loop ends either by returning `true` (an integer sits on top) or once the stack is fully empty — then it returns `false`.

---

### Next

```go
func (ni *NestedIterator) Next() int {
    top := ni.stack[len(ni.stack)-1]

    elem := top[0]

    ni.stack[len(ni.stack)-1] = top[1:]

    return elem.GetInteger()
}
```

`Next()` doesn't unpack anything — that work has already been done by `HasNext()`. It assumes an integer is sitting on top of the stack right now (the iterator contract: `HasNext()` is called before `Next()`). What's left is to grab the first element of the top level, advance that level by one (`top[1:]`), and return its value via `GetInteger()`.

---

### Why it works

The stack always holds the path from the current read position back to the root of the list: each level is the "tail" not yet processed at that depth of nesting. Once `HasNext()` reaches an integer, it's guaranteed to be the next number in traversal order — every element before it, at every level of the stack, has already been either returned through `Next()` or unpacked into deeper levels.

An empty level on top of the stack means that level of nesting has been fully consumed, and processing falls back to the level above it — this is exactly how a depth-first traversal works, just without recursion: the stack here is explicit, replacing the call stack.

---

### Operations

| Operation     | Complexity          | Description                                        |
| ------------- | -------------------- | ---------------------------------------------------- |
| `Constructor` | O(1)                 | the list is just pushed onto the stack, nothing is computed |
| `HasNext`     | amortized O(1)       | each element gets unpacked by the stack at most once over the iterator's lifetime |
| `Next`        | O(1)                 | reading and advancing the top level of the stack     |

Every element — integer or list — enters the stack and leaves it exactly once over a full traversal, so the total cost of all `HasNext()` calls across the whole traversal is linear in the total number of elements (integers and nested lists combined), not quadratic.

---

### Implementation details

* a stack level is a `[]*NestedInteger`, the "tail" of unprocessed elements; advancing through a level means dropping the head of the slice (`top[1:]`), no separate index counter needed
* empty nested lists (`[]`) don't need special handling: such a level just gets pushed onto the stack and gets popped as empty on the very next `HasNext()` iteration
* `HasNext()` can be called any number of times in a row without side effects — if the stack is already in the "integer on top" state, repeated calls just return `true` immediately
* the method receiver is `*NestedIterator`, the stack is mutated in place, no copying on every call

---

### Why not unpack everything upfront

```go
func ConstructorEager(nestedList []*NestedInteger) *NestedIterator {
    flat := make([]int, 0)
    var flatten func([]*NestedInteger)
    flatten = func(list []*NestedInteger) {
        for _, ni := range list {
            if ni.IsInteger() {
                flat = append(flat, ni.GetInteger())
            } else {
                flatten(ni.GetList())
            }
        }
    }
    flatten(nestedList)
    // ...Next() would then just read flat by index
}
```

Also correct, but all the unpacking work runs inside `Constructor`, even if `Next()` never gets called. It also needs separate memory for the entire flat list. The `stack`-based version only ever unpacks as much as was actually read — if the traversal stops halfway through, the rest of the list is never touched.

---

### Limitations

* `Next()` doesn't check whether an element is available — calling it without a preceding `HasNext() == true` will panic on an empty slice
* the `NestedList` passed into `Constructor` must not be mutated from outside during traversal — the stack holds slices pointing at that same underlying data

---

<br>

> Список не разворачивается заранее — он разворачивается ровно настолько, насколько его успели прочитать.
>
> The list isn't unpacked in advance — it's unpacked exactly as far as it's actually been read.
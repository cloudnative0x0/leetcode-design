# 705 · Design HashSet

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

**Difficulty:** Easy | **Time:** O(1) amortized | **Space:** O(n + m)

---

## Русский

### Решение через метод цепочек над фиксированным массивом бакетов
Срез из голов связных списков плюс одна хеш-функция. Без ресайза, без порога заполнения — `bucketCount` задаётся один раз в конструкторе и не меняется на протяжении жизни объекта.

Число бакетов — фиксированное простое число `10007`, выбранное один раз и больше не пересматриваемое.

---

### Структура

```go
type Node struct {
    key  int
    next *Node
}

type MyHashSet struct {
    buckets     []*Node
    bucketCount int
}
```

- `buckets[i]` — голова односвязного списка, либо `nil`, если бакет пуст
- `bucketCount` — число ячеек в `buckets`, задаётся один раз в `Constructor`
- `Node.next` — обычный указатель Go, без ручного владения, сборщик мусора делает всю работу

> LFU держит две карты под два разных индекса. HashSet — ровно одну: ключ → бакет.

---

### Хеш-функция

Всё решение сворачивается в одну строку:

```go
h := key % mhs.bucketCount
if h < 0 {
    h += mhs.bucketCount
}
```

В виде математики, где $k$ — ключ, $m$ — `bucketCount`:

$$
h(k) = k \bmod m
$$

Оператор `%` в Go сохраняет знак делимого, поэтому при $k < 0$ результат сам по себе может быть отрицательным. Индекс бакета обязан лежать в $[0, m)$, поэтому отрицательная ветка возвращает его обратно в диапазон:

$$
h(k) =
\begin{cases}
(k \bmod m) + m, & k \bmod m < 0 \\
k \bmod m, & k \bmod m \ge 0
\end{cases}
$$

Битовых сдвигов в этой функции нет нигде, и на то есть конкретная причина.

---

### Почему нет битовой маски и почему нет log

Распространённая альтернатива для индексации бакетов — та, что использует `java.util.HashMap` — держит число бакетов степенью двойки, $m = 2^p$, и заменяет остаток от деления на побитовое И:

$$
h(k) = k \,\&\, (m - 1)
$$

Этот трюк работает только потому, что $m - 1$ в двоичном виде — это подряд идущие $p$ единиц, когда $m$ — степень двойки; операция И с такой маской математически идентична `k % m` для неотрицательных `k`, но стоит одну инструкцию вместо деления. $p = \log_2 m$ появляется там только как величина сдвига / ширина маски при росте таблицы: удвоение вместимости означает пересчёт $p$ и перестройку всех бакетов заново.

Данная реализация идёт другим путём. `bucketCount = 10007` — простое число, а не степень двойки, поэтому $m - 1$ не даёт чистую битовую маску — И с ней рассеивало бы ключи неравномерно, скучивая их в подмножестве бакетов вместо равномерного распределения. Ресайза в коде тоже нет, а значит нет и события роста, которое вообще потребовало бы пересчёта сдвига через $\log_2 m$. Реализация фиксируется на делении по модулю относительно простого числа и на этом останавливается.

Простой модуль — как раз причина, по которой `%` здесь безопасен: при плохо распределённой последовательности ключей (кратные некоторому небольшому числу, последовательные ID и т.п.) простое $m$ разбивает общие делители между ключами и модулем, распределяя ключи по бакетам равномернее, чем это сделало бы круглое число вроде $10000$.

---

### Заполнение и ожидаемая длина цепочки

При $n$ хранимых ключах и $m$ бакетах коэффициент заполнения:

$$
\alpha = \frac{n}{m}
$$

При примерно равномерном хешировании ожидаемая занятость бакета:

$$
E[\text{длина цепочки}] = \alpha
$$

что даёт `Add` / `Remove` / `Contains` ожидаемую стоимость

$$
O(1 + \alpha)
$$

Поскольку $m$ никогда не растёт, $\alpha$ растёт без ограничений при $n \to \infty$ — это осознанный компромисс отказа от логики ресайза. При $n \ll m$ поведение неотличимо от $O(1)$; за этим пределом оно смещается к $O(n)$ на вызов в вырожденном случае одного перегруженного бакета.

---

### Add

```go
func (mhs *MyHashSet) Add(key int) {
    index := mhs.hash(key)
    ...
}
```

Пустой бакет → создаётся новый головной узел, готово. Непустой бакет → обход цепочки; совпадение ключа означает, что значение уже есть и ничего не меняется; достижение хвоста без совпадения означает добавление нового узла туда.

---

### Remove

```go
func (mhs *MyHashSet) Remove(key int) {
    index := mhs.hash(key)
    ...
}
```

Совпадение с головой — отдельный случай: `buckets[index]` переприсваивается напрямую на `current.next`. В любом другом месте цепочки `next` предыдущего узла перепривязывается в обход удаляемого узла.

---

### Contains

```go
func (mhs *MyHashSet) Contains(key int) bool {
    index := mhs.hash(key)
    ...
}
```

Обычный линейный проход по цепочке одного бакета, без мутаций.

---

### Заметки по stdlib

- `container/list` здесь не используется — цепочка это самописный односвязный список, только `Node.next`, обратный указатель не нужен, так как обход всегда однонаправленный
- тип `map` для бакетов не используется — срез `[]*Node`, индексируемый хешем, даёт O(1) доступ к бакету без накладных расходов внутреннего устройства карт Go
- простота `bucketCount` делает ту работу по распределению, которую иначе пришлось бы делать таблице побольше или более сложной хеш-функции

---

> Вся хеш-функция — это один остаток от деления и одна ветка; всё остальное в классе существует, чтобы корректно обойти связный список.

---

## English

### Solution uses separate chaining over a fixed bucket array
A slice of linked-list heads plus a single hash function. No resizing, no load-factor threshold — `bucketCount` is fixed at construction and stays that way for the whole lifetime of the set.

The bucket count is a fixed prime, `10007`, chosen once and never revisited.

---

### Architecture

```go
type Node struct {
    key  int
    next *Node
}

type MyHashSet struct {
    buckets     []*Node
    bucketCount int
}
```

- `buckets[i]` — head of a singly linked list, or `nil` if the bucket is empty
- `bucketCount` — number of slots in `buckets`, set once in `Constructor`
- `Node.next` — no ownership tricks here, plain Go pointers, GC does the cleanup

> LFU keeps two maps for two different indexes. HashSet needs exactly one: key → bucket.

---

### Hash function

The whole set collapses to one line:

```go
h := key % mhs.bucketCount
if h < 0 {
    h += mhs.bucketCount
}
```

As math, with $k$ the key and $m$ = `bucketCount`:

$$
h(k) = k \bmod m
$$

Go's `%` keeps the sign of the dividend, so for $k < 0$ the raw result can itself be negative. The bucket index has to live in $[0, m)$, so the negative branch shifts it back into range:

$$
h(k) =
\begin{cases}
(k \bmod m) + m, & k \bmod m < 0 \\
k \bmod m, & k \bmod m \ge 0
\end{cases}
$$

No bit shifting happens anywhere in this function, and there's a specific reason for that.

---

### Why no bit masking, why no log

A common alternative for bucket indexing — the one `java.util.HashMap` uses — keeps the bucket count as a power of two, $m = 2^p$, and replaces the modulo with a bitwise AND:

$$
h(k) = k \,\&\, (m - 1)
$$

That trick only works because $m - 1$ is a run of $p$ ones in binary when $m$ is a power of two — AND-ing against it is mathematically identical to `k % m` for non-negative `k`, but costs one instruction instead of a division. $p = \log_2 m$ shows up there only as the shift amount / mask width when the table grows: doubling capacity means recomputing $p$ and rebuilding every bucket.

This implementation doesn't take that path. `bucketCount = 10007` is prime, not a power of two, so $m - 1$ is not a clean bitmask — AND-ing against it would scatter keys unevenly, clustering them into a subset of buckets instead of spreading them out. There's also no resize step anywhere in the code, so there's no growth event that would ever need a $\log_2 m$ shift recalculation in the first place. The design commits to modulo division against a fixed prime and stays there.

A prime modulus is the reason `%` is safe here in the first place: for a poorly-distributed key sequence (multiples of some small number, sequential IDs, etc.), a prime $m$ breaks up shared factors between keys and the modulus, spreading keys across buckets more evenly than a round number like $10000$ would.

---

### Load and expected chain length

With $n$ keys stored in $m$ buckets, load factor:

$$
\alpha = \frac{n}{m}
$$

Under a roughly uniform hash, expected bucket occupancy:

$$
E[\text{chain length}] = \alpha
$$

which gives `Add` / `Remove` / `Contains` an expected cost of

$$
O(1 + \alpha)
$$

Since $m$ never grows, $\alpha$ climbs without bound as $n \to \infty$ — this is the deliberate trade-off of skipping resize logic. At $n \ll m$ it's indistinguishable from $O(1)$; past that it slides toward $O(n)$ per call in the degenerate case of one overloaded bucket.

---

### Add

```go
func (mhs *MyHashSet) Add(key int) {
    index := mhs.hash(key)
    ...
}
```

Empty bucket → new head node, done. Non-empty bucket → walk the chain; a matching key means the value already exists and nothing changes; hitting the tail means append a new node there.

---

### Remove

```go
func (mhs *MyHashSet) Remove(key int) {
    index := mhs.hash(key)
    ...
}
```

Head match is a special case — `buckets[index]` gets reassigned straight to `current.next`. Anywhere else in the chain, the previous node's `next` gets rewired around the removed node.

---

### Contains

```go
func (mhs *MyHashSet) Contains(key int) bool {
    index := mhs.hash(key)
    ...
}
```

Plain linear scan of one bucket's chain, no mutation.

---

### stdlib notes

- no `container/list` here — the chain is a hand-rolled singly linked list, `Node.next` only, no back-pointer needed since traversal is one-directional
- no `map` type used for the buckets — a `[]*Node` slice indexed by hash gives O(1) bucket access without the overhead of Go's map internals
- primality of `bucketCount` is doing the distribution work that a bigger table or a fancier hash would otherwise have to do

---

> The whole hash function is one modulo and one branch — everything else in the class exists to walk a linked list correctly.
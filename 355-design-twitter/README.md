# Twitter

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

## Русский

### Идея

Сервис хранит твиты каждого пользователя, связи подписок и общий счётчик времени. Счётчик определяет порядок публикаций: чем больше `timeStamp`, тем позже был создан твит.

При запросе ленты собираются твиты самого пользователя и всех пользователей, на которых он подписан. Кандидаты сортируются от новых к старым, после чего возвращаются первые десять идентификаторов.

---

### Tweet

```go
type Tweet struct {
    id        int
    timeStamp int
}
```

- `id` — идентификатор твита;
- `timeStamp` — порядковый номер публикации.

Отдельная структура нужна, потому что для результата используется `id`, а для построения ленты — время публикации.

---

### Twitter

```go
type Twitter struct {
    tweets    map[int][]Tweet
    follows   map[int]map[int]bool
    timeStamp int
}
```

- `tweets[userId]` содержит твиты пользователя в порядке добавления;
- `follows[followerId]` хранит множество пользователей, на которых подписан `followerId`;
- `timeStamp` — единый счётчик для всех публикаций.

Вложенная карта используется как множество: ключом является `followeeId`, а значение `bool` отмечает наличие подписки.

---

### Constructor

```go
func Constructor() Twitter {
    return Twitter{
        tweets:    make(map[int][]Tweet),
        follows:   make(map[int]map[int]bool),
        timeStamp: 0,
    }
}
```

Конструктор создаёт пустые карты и устанавливает счётчик времени в ноль. Пользователи появляются в структурах только после публикации твита или оформления подписки.

---

### PostTweet

```go
func (tw *Twitter) PostTweet(
    userId int,
    tweetId int,
) {
    tw.timeStamp++

    newTweet := Tweet{
        id:        tweetId,
        timeStamp: tw.timeStamp,
    }

    tw.tweets[userId] = append(
        tw.tweets[userId],
        newTweet,
    )
}
```

Перед созданием твита глобальный счётчик увеличивается. Благодаря этому две публикации не получают одинаковое время, даже если они принадлежат разным пользователям.

Новый твит добавляется в конец среза соответствующего пользователя.

---

### Follow

```go
func (tw *Twitter) Follow(
    followerId int,
    followeeId int,
) {
    if tw.follows[followerId] == nil {
        tw.follows[followerId] = make(map[int]bool)
    }

    tw.follows[followerId][followeeId] = true
}
```

Если у пользователя ещё нет множества подписок, оно создаётся при первом вызове.

Повторная подписка записывает `true` по тому же ключу и не создаёт дубликатов.

---

### Unfollow

```go
func (tw *Twitter) Unfollow(
    followerId int,
    followeeId int,
) {
    if _, ok := tw.follows[followerId]; ok {
        delete(tw.follows[followerId], followeeId)
    }
}
```

Метод удаляет `followeeId` из множества подписок. Если пользователя или связи не существует, состояние не меняется.

---

### GetNewsFeed

#### Сбор кандидатов

```go
var candidates []Tweet

candidates = append(
    candidates,
    tw.tweets[userId]...,
)

for followeeId := range tw.follows[userId] {
    if followeeId != userId {
        candidates = append(
            candidates,
            tw.tweets[followeeId]...,
        )
    }
}
```

Сначала добавляются собственные публикации пользователя, затем публикации его подписок.

Проверка `followeeId != userId` не позволяет добавить собственные твиты второй раз, если пользователь подписался на себя.

Обход отсутствующей карты безопасен: для пользователя без подписок цикл не выполнится.

#### Сортировка

```go
for i := 0; i < len(candidates); i++ {
    for j := i + 1; j < len(candidates); j++ {
        if candidates[j].timeStamp >
            candidates[i].timeStamp {
            candidates[i], candidates[j] =
                candidates[j], candidates[i]
        }
    }
}
```

Для каждой позиции ищется более новый твит в оставшейся части среза.

После завершения циклов кандидаты расположены по убыванию `timeStamp`: от последней публикации к самой ранней.

#### Ограничение ленты

```go
feedSize := 10

if len(candidates) < 10 {
    feedSize = len(candidates)
}

res := make([]int, feedSize)

for i := 0; i < feedSize; i++ {
    res[i] = candidates[i].id
}
```

Размер результата равен десяти или количеству кандидатов, если их меньше. В итоговый срез переносятся только идентификаторы твитов.

---

### Пример

```go
twitter := Constructor()

twitter.PostTweet(1, 5)
twitter.GetNewsFeed(1) // [5]

twitter.Follow(1, 2)
twitter.PostTweet(2, 6)
twitter.GetNewsFeed(1) // [6, 5]

twitter.Unfollow(1, 2)
twitter.GetNewsFeed(1) // [5]
```

Твит `6` отображается первым, поскольку он опубликован позже твита `5`.

После удаления подписки публикации второго пользователя больше не входят в ленту первого.

---

### Почему это работает

Глобальный `timeStamp` задаёт однозначный порядок всех твитов независимо от автора. Поэтому после объединения нужных срезов достаточно отсортировать записи по этому полю.

Карта подписок определяет, публикации каких пользователей входят в выборку. Обрезание уже отсортированного списка до десяти элементов оставляет десять самых свежих твитов.

---

### Сложность

Пусть `c` — общее количество твитов пользователя и всех его подписок, попавших в список кандидатов.

| Операция | Время | Дополнительная память |
| --- | --- | --- |
| `Constructor` | O(1) | O(1) |
| `PostTweet` | O(1) амортизированно | O(1) на твит |
| `Follow` | O(1) в среднем | O(1) на связь |
| `Unfollow` | O(1) в среднем | O(1) |
| `GetNewsFeed` | O(c²) | O(c) |

Квадратичное время `GetNewsFeed` связано с двумя вложенными циклами сортировки.

Итоговый срез содержит не больше десяти элементов, но перед сортировкой в `candidates` собираются все доступные твиты.

---

### Детали реализации

- лента пользователя без твитов и подписок является пустым срезом;
- повторный `Follow` не дублирует связь;
- подписка на самого себя не дублирует собственные твиты в ленте;
- `Unfollow` отсутствующей связи ничего не меняет;
- данные существуют только в памяти объекта `Twitter`;
- структура не защищена от одновременного изменения из нескольких goroutine.

---

## English

### Idea

The service stores each user's tweets, follow relationships, and one global time counter. The counter determines publication order: a larger `timeStamp` means that the tweet was posted later.

When a news feed is requested, the service collects the user's own tweets and the tweets of every followed user. The candidates are sorted from newest to oldest, and the first ten identifiers are returned.

---

### Tweet

```go
type Tweet struct {
    id        int
    timeStamp int
}
```

- `id` — the tweet identifier;
- `timeStamp` — the publication sequence number.

A separate structure is needed because the result uses `id`, while feed ordering uses the publication time.

---

### Twitter

```go
type Twitter struct {
    tweets    map[int][]Tweet
    follows   map[int]map[int]bool
    timeStamp int
}
```

- `tweets[userId]` contains the user's tweets in insertion order;
- `follows[followerId]` stores the set of users followed by `followerId`;
- `timeStamp` is a single counter shared by every publication.

The nested map works as a set: `followeeId` is the key, and its Boolean value marks the presence of the relationship.

---

### Constructor

```go
func Constructor() Twitter {
    return Twitter{
        tweets:    make(map[int][]Tweet),
        follows:   make(map[int]map[int]bool),
        timeStamp: 0,
    }
}
```

The constructor creates empty maps and sets the time counter to zero. A user appears in the data structures only after posting a tweet or following another user.

---

### PostTweet

```go
func (tw *Twitter) PostTweet(
    userId int,
    tweetId int,
) {
    tw.timeStamp++

    newTweet := Tweet{
        id:        tweetId,
        timeStamp: tw.timeStamp,
    }

    tw.tweets[userId] = append(
        tw.tweets[userId],
        newTweet,
    )
}
```

The global counter is incremented before a tweet is created. Two publications therefore cannot receive the same time, even when they belong to different users.

The new tweet is appended to the corresponding user's slice.

---

### Follow

```go
func (tw *Twitter) Follow(
    followerId int,
    followeeId int,
) {
    if tw.follows[followerId] == nil {
        tw.follows[followerId] = make(map[int]bool)
    }

    tw.follows[followerId][followeeId] = true
}
```

If the user does not have a follow set yet, one is created on the first call.

Following the same user repeatedly writes `true` to the same key and does not create duplicates.

---

### Unfollow

```go
func (tw *Twitter) Unfollow(
    followerId int,
    followeeId int,
) {
    if _, ok := tw.follows[followerId]; ok {
        delete(tw.follows[followerId], followeeId)
    }
}
```

The method removes `followeeId` from the follow set. If the user or relationship does not exist, the state remains unchanged.

---

### GetNewsFeed

#### Collecting candidates

```go
var candidates []Tweet

candidates = append(
    candidates,
    tw.tweets[userId]...,
)

for followeeId := range tw.follows[userId] {
    if followeeId != userId {
        candidates = append(
            candidates,
            tw.tweets[followeeId]...,
        )
    }
}
```

The user's own publications are added first, followed by the publications of every followed user.

The `followeeId != userId` check prevents a self-follow from adding the user's tweets twice.

Ranging over a missing map is safe, so the loop does nothing for a user without follows.

#### Sorting

```go
for i := 0; i < len(candidates); i++ {
    for j := i + 1; j < len(candidates); j++ {
        if candidates[j].timeStamp >
            candidates[i].timeStamp {
            candidates[i], candidates[j] =
                candidates[j], candidates[i]
        }
    }
}
```

For each position, the remaining slice is searched for a newer tweet.

When the loops finish, candidates are ordered by decreasing `timeStamp`, from the latest publication to the earliest one.

#### Limiting the feed

```go
feedSize := 10

if len(candidates) < 10 {
    feedSize = len(candidates)
}

res := make([]int, feedSize)

for i := 0; i < feedSize; i++ {
    res[i] = candidates[i].id
}
```

The result size is ten or the number of candidates when fewer are available. Only tweet identifiers are copied into the resulting slice.

---

### Example

```go
twitter := Constructor()

twitter.PostTweet(1, 5)
twitter.GetNewsFeed(1) // [5]

twitter.Follow(1, 2)
twitter.PostTweet(2, 6)
twitter.GetNewsFeed(1) // [6, 5]

twitter.Unfollow(1, 2)
twitter.GetNewsFeed(1) // [5]
```

Tweet `6` appears first because it was posted after tweet `5`.

Once the follow relationship is removed, the second user's publications no longer appear in the first user's feed.

---

### Why it works

The global `timeStamp` establishes an unambiguous order for every tweet regardless of its author. After the required slices are combined, sorting by this field is enough to build the feed.

The follow map determines which users contribute publications. Taking the first ten entries from the sorted candidates leaves the ten most recent tweets.

---

### Complexity

Let `c` be the total number of candidate tweets belonging to the user and all followed users.

| Operation | Time | Extra space |
| --- | --- | --- |
| `Constructor` | O(1) | O(1) |
| `PostTweet` | O(1) amortized | O(1) per tweet |
| `Follow` | O(1) average | O(1) per relationship |
| `Unfollow` | O(1) average | O(1) |
| `GetNewsFeed` | O(c²) | O(c) |

The two nested sorting loops make `GetNewsFeed` quadratic.

The result contains at most ten elements, but all available tweets are collected in `candidates` before sorting.

---

### Implementation details

- a user without tweets or follows receives an empty slice;
- repeated `Follow` calls do not duplicate a relationship;
- following oneself does not duplicate one's own tweets in the feed;
- `Unfollow` on a missing relationship has no effect;
- all data exists only in the memory of the `Twitter` instance;
- the structure is not protected against concurrent modification from multiple goroutines.
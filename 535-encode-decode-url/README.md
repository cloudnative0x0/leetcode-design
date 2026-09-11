# Codec

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

## Русский

### Идея

`Codec` преобразует длинный URL в короткий адрес и умеет восстановить исходную строку. Сокращать содержимое URL не требуется: достаточно выдать ему короткий уникальный идентификатор и сохранить связь между двумя адресами.

В решении идентификаторы создаются последовательно: `1`, `2`, `3` и далее. Полный короткий адрес состоит из префикса `http://tinyurl.com/` и текущего идентификатора.

---

### Структура

```go
type Codec struct {
    hashMap map[string]string
    id      int
}
```

- `hashMap` хранит соответствие между коротким и исходным URL;
- `id` содержит номер, использованный при последнем кодировании.

Ключом карты является готовый короткий адрес, а значением — исходная строка:

```text
http://tinyurl.com/1 -> https://leetcode.com/problems/design-tinyurl
http://tinyurl.com/2 -> https://example.com/articles/42
```

---

### Constructor

```go
func Constructor() Codec {
    return Codec{
        hashMap: make(map[string]string),
        id:      0,
    }
}
```

Конструктор создаёт пустую карту и устанавливает счётчик в ноль. Первый вызов `encode()` увеличит его до единицы.

Каждый объект `Codec` имеет собственную карту и собственный счётчик.

---

### encode

```go
func (cc *Codec) encode(longUrl string) string {
    cc.id++

    shortUrl := "http://tinyurl.com/" +
        strconv.Itoa(cc.id)

    cc.hashMap[shortUrl] = longUrl

    return shortUrl
}
```

Метод выполняет три действия:

1. увеличивает `id`;
2. переводит число в строку через `strconv.Itoa` и добавляет его к префиксу;
3. сохраняет пару `shortUrl -> longUrl` в карте.

Например, первый переданный адрес получит ключ `http://tinyurl.com/1`, а второй — `http://tinyurl.com/2`.

Повторное кодирование одного и того же длинного URL создаёт новый короткий адрес. Обратной карты `longUrl -> shortUrl` в этой реализации нет.

---

### decode

```go
func (cc *Codec) decode(shortUrl string) string {
    return cc.hashMap[shortUrl]
}
```

Короткий URL уже является ключом карты, поэтому исходный адрес возвращается прямым обращением к `hashMap`.

Если ключ отсутствует, Go возвращает нулевое значение для `string` — пустую строку.

---

### Пример

```go
codec := Constructor()

shortUrl := codec.encode(
    "https://leetcode.com/problems/design-tinyurl",
)

longUrl := codec.decode(shortUrl)
```

Состояние после вызова `encode()`:

```text
id = 1

hashMap = {
    "http://tinyurl.com/1":
        "https://leetcode.com/problems/design-tinyurl"
}
```

Вызов `decode("http://tinyurl.com/1")` возвращает сохранённый длинный URL.

---

### Почему это работает

Каждый вызов `encode()` использует новое значение счётчика, поэтому внутри одного объекта `Codec` короткие адреса не повторяются.

Сохранённая в карте пара позволяет выполнить обратное преобразование без разбора идентификатора.

Длинный URL хранится без изменений, поэтому параметры запроса, фрагменты и другие части адреса восстанавливаются в исходном виде.

---

### Сложность

| Операция | Среднее время | Дополнительная память |
| --- | --- | --- |
| `Constructor` | O(1) | O(1) |
| `encode` | O(1) | O(1) на одну запись |
| `decode` | O(1) | O(1) |

После кодирования `n` адресов карта занимает O(n) памяти. Формирование строк зависит от их длины, но операции с картой в среднем выполняются за O(1).

---

### Детали реализации

- данные существуют только в памяти текущего объекта `Codec`;
- новый объект начинает нумерацию заново;
- одинаковый длинный URL при каждом вызове получает новый идентификатор;
- неизвестный короткий URL декодируется в пустую строку;
- счётчик и карта не защищены от одновременного доступа из нескольких goroutine;
- префикс `http://tinyurl.com/` является частью формата короткого адреса.

---

## English

### Idea

`Codec` converts a long URL into a short address and restores the original string later. The URL contents do not need to be compressed. It is enough to assign a short unique identifier and store the relationship between the two addresses.

The solution creates identifiers sequentially: `1`, `2`, `3`, and so on. A complete short address consists of the `http://tinyurl.com/` prefix followed by the current identifier.

---

### Structure

```go
type Codec struct {
    hashMap map[string]string
    id      int
}
```

- `hashMap` stores the mapping from a short URL to its original URL;
- `id` contains the number used by the most recent encoding operation.

The completed short address is the map key, and the original string is its value:

```text
http://tinyurl.com/1 -> https://leetcode.com/problems/design-tinyurl
http://tinyurl.com/2 -> https://example.com/articles/42
```

---

### Constructor

```go
func Constructor() Codec {
    return Codec{
        hashMap: make(map[string]string),
        id:      0,
    }
}
```

The constructor creates an empty map and sets the counter to zero. The first `encode()` call increments it to one.

Each `Codec` instance has its own map and counter.

---

### encode

```go
func (cc *Codec) encode(longUrl string) string {
    cc.id++

    shortUrl := "http://tinyurl.com/" +
        strconv.Itoa(cc.id)

    cc.hashMap[shortUrl] = longUrl

    return shortUrl
}
```

The method performs three steps:

1. increments `id`;
2. converts the number to a string with `strconv.Itoa` and appends it to the prefix;
3. stores the `shortUrl -> longUrl` pair in the map.

For example, the first supplied address receives `http://tinyurl.com/1`, and the second receives `http://tinyurl.com/2`.

Encoding the same long URL repeatedly creates a new short URL each time. This implementation does not have a reverse `longUrl -> shortUrl` map.

---

### decode

```go
func (cc *Codec) decode(shortUrl string) string {
    return cc.hashMap[shortUrl]
}
```

The short URL is already a map key, so the original address is returned through a direct `hashMap` lookup.

If the key does not exist, Go returns the zero value for `string`, which is an empty string.

---

### Example

```go
codec := Constructor()

shortUrl := codec.encode(
    "https://leetcode.com/problems/design-tinyurl",
)

longUrl := codec.decode(shortUrl)
```

State after the `encode()` call:

```text
id = 1

hashMap = {
    "http://tinyurl.com/1":
        "https://leetcode.com/problems/design-tinyurl"
}
```

Calling `decode("http://tinyurl.com/1")` returns the stored long URL.

---

### Why it works

Every `encode()` call uses a new counter value, so short addresses do not repeat within one `Codec` instance.

The pair stored in the map makes the reverse conversion possible without parsing the identifier.

The long URL is stored without modification, so query parameters, fragments, and other address components are restored exactly as supplied.

---

### Complexity

| Operation | Average time | Extra space |
| --- | --- | --- |
| `Constructor` | O(1) | O(1) |
| `encode` | O(1) | O(1) per entry |
| `decode` | O(1) | O(1) |

After encoding `n` addresses, the map occupies O(n) memory. String construction depends on string length, while map operations take O(1) average time.

---

### Implementation details

- data exists only in the memory of the current `Codec` instance;
- a new instance starts its numbering from the beginning;
- the same long URL receives a new identifier on every call;
- an unknown short URL decodes to an empty string;
- the counter and map are not protected from concurrent goroutine access;
- the `http://tinyurl.com/` prefix is part of the short URL format.
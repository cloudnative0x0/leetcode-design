# NeighborSum

<p style="text-align: left">
  <a href="#русский">Русский</a> ・ <a href="#english">English</a>
</p>

---

## Русский

### Идея

Для заданного значения нужно найти сумму соседей двух типов:

- смежные — сверху, снизу, слева и справа;
- диагональные — по четырём диагоналям.

Все значения в `grid` уникальны. Поэтому в конструкторе для каждого значения сохраняются его строка и столбец. Запрос начинает работу сразу с нужной ячейки и не просматривает всю таблицу.

---

### Структура

```go
type NeighborSum struct {
    grid   [][]int
    coords map[int][2]int
}
```

- `grid` хранит исходную квадратную матрицу;
- `coords` связывает значение с координатами `[row, column]`.

Например, для такой матрицы:

```text
0 1 2
3 4 5
6 7 8
```

в карте будут записи:

```text
coords[0] = [0, 0]
coords[4] = [1, 1]
coords[8] = [2, 2]
```

---

### Constructor

```go
func Constructor(grid [][]int) NeighborSum {
    coords := make(map[int][2]int)

    for i := 0; i < len(grid); i++ {
        for v := 0; v < len(grid[i]); v++ {
            coords[grid[i][v]] = [2]int{i, v}
        }
    }

    return NeighborSum{
        grid:   grid,
        coords: coords,
    }
}
```

Два цикла обходят матрицу один раз. Значение используется как ключ, а пара индексов — как координаты ячейки.

После построения карты поиск позиции занимает O(1) в среднем.

---

### AdjacentSum

Сначала метод находит координаты переданного значения:

```go
pos, exists := ns.coords[value]
if !exists {
    return 0
}

row, column := pos[0], pos[1]
```

Для обхода смежных соседей используется массив направлений:

```go
fill := [4][2]int{
    {-1, 0}, // up
    {1, 0},  // down
    {0, -1}, // left
    {0, 1},  // right
}
```

Каждая пара задаёт смещение относительно найденной позиции:

- `{-1, 0}` — строка выше;
- `{1, 0}` — строка ниже;
- `{0, -1}` — столбец слева;
- `{0, 1}` — столбец справа.

Для значения `4` из примера будут проверены элементы `1`, `7`, `3` и `5`:

```text
1 + 7 + 3 + 5 = 16
```

Перед обращением к матрице новые координаты проверяются:

```go
if nextRow >= 0 &&
    nextRow < n &&
    nextColumn >= 0 &&
    nextColumn < n {
    sum += ns.grid[nextRow][nextColumn]
}
```

Поэтому один и тот же цикл работает для центральных, граничных и угловых ячеек.

---

### DiagonalSum

Поиск координат и проверка границ выполняются так же, как в `AdjacentSum`. Отличается только массив направлений:

```go
fill := [4][2]int{
    {-1, -1}, // up-left
    {-1, 1},  // up-right
    {1, -1},  // down-left
    {1, 1},   // down-right
}
```

Для центрального значения `4` диагональными соседями будут `0`, `2`, `6` и `8`:

```text
0 + 2 + 6 + 8 = 16
```

Если значение расположено в углу или возле края, координаты за пределами матрицы пропускаются.

---

### Значение отсутствует

Оба метода сначала проверяют карту:

```go
pos, exists := ns.coords[value]
if !exists {
    return 0
}
```

По условию запросы содержат значения из матрицы. Дополнительная проверка не позволяет использовать нулевые координаты карты, если метод всё же вызван с неизвестным значением.

---

### Почему это работает

Карта `coords` позволяет сразу перейти к нужной ячейке. После этого остаётся проверить только четыре фиксированных направления.

`AdjacentSum` использует вертикальные и горизонтальные смещения, а `DiagonalSum` — диагональные. Проверка границ исключает координаты, которые не принадлежат матрице.

---

### Сложность

| Операция | Время | Дополнительная память |
| --- | --- | --- |
| `Constructor` | O(n²) | O(n²) |
| `AdjacentSum` | O(1) | O(1) |
| `DiagonalSum` | O(1) | O(1) |

Здесь `n` — сторона квадратной матрицы.

Конструктор обходит все `n²` элементов и сохраняет их координаты. Каждый запрос проверяет ровно четыре направления, поэтому работает за постоянное время.

---

## English

### Idea

For a given value, the service calculates two kinds of neighbor sums:

- adjacent — above, below, left, and right;
- diagonal — in the four diagonal directions.

Every value in `grid` is unique. The constructor therefore stores the row and column of each value. A query starts at the required cell without scanning the entire grid.

---

### Structure

```go
type NeighborSum struct {
    grid   [][]int
    coords map[int][2]int
}
```

- `grid` stores the original square matrix;
- `coords` maps a value to its `[row, column]` coordinates.

For this grid:

```text
0 1 2
3 4 5
6 7 8
```

the map contains:

```text
coords[0] = [0, 0]
coords[4] = [1, 1]
coords[8] = [2, 2]
```

---

### Constructor

```go
func Constructor(grid [][]int) NeighborSum {
    coords := make(map[int][2]int)

    for i := 0; i < len(grid); i++ {
        for v := 0; v < len(grid[i]); v++ {
            coords[grid[i][v]] = [2]int{i, v}
        }
    }

    return NeighborSum{
        grid:   grid,
        coords: coords,
    }
}
```

The two loops visit the matrix once. A cell value becomes the key, while its two indexes become the stored coordinates.

After the map is built, locating a value takes O(1) average time.

---

### AdjacentSum

The method first finds the coordinates of the supplied value:

```go
pos, exists := ns.coords[value]
if !exists {
    return 0
}

row, column := pos[0], pos[1]
```

An array of directions describes the adjacent cells:

```go
fill := [4][2]int{
    {-1, 0}, // up
    {1, 0},  // down
    {0, -1}, // left
    {0, 1},  // right
}
```

Each pair is an offset from the located position:

- `{-1, 0}` — one row above;
- `{1, 0}` — one row below;
- `{0, -1}` — one column to the left;
- `{0, 1}` — one column to the right.

For value `4` in the example, the method checks `1`, `7`, `3`, and `5`:

```text
1 + 7 + 3 + 5 = 16
```

The new coordinates are validated before accessing the grid:

```go
if nextRow >= 0 &&
    nextRow < n &&
    nextColumn >= 0 &&
    nextColumn < n {
    sum += ns.grid[nextRow][nextColumn]
}
```

The same loop can therefore handle center, edge, and corner cells.

---

### DiagonalSum

Coordinate lookup and boundary checking work the same way as in `AdjacentSum`. Only the direction array is different:

```go
fill := [4][2]int{
    {-1, -1}, // up-left
    {-1, 1},  // up-right
    {1, -1},  // down-left
    {1, 1},   // down-right
}
```

For the center value `4`, the diagonal neighbors are `0`, `2`, `6`, and `8`:

```text
0 + 2 + 6 + 8 = 16
```

When the value is located in a corner or near an edge, coordinates outside the grid are skipped.

---

### Missing value

Both methods check the map first:

```go
pos, exists := ns.coords[value]
if !exists {
    return 0
}
```

The problem guarantees that query values exist in the matrix. This extra check prevents the methods from treating the map's zero value as valid coordinates when an unknown value is supplied.

---

### Why it works

The `coords` map provides direct access to the required cell. After that, only four fixed directions need to be checked.

`AdjacentSum` uses vertical and horizontal offsets, while `DiagonalSum` uses diagonal offsets. The boundary check excludes coordinates that do not belong to the grid.

---

### Complexity

| Operation | Time | Extra space |
| --- | --- | --- |
| `Constructor` | O(n²) | O(n²) |
| `AdjacentSum` | O(1) | O(1) |
| `DiagonalSum` | O(1) | O(1) |

Here, `n` is the side length of the square matrix.

The constructor visits all `n²` elements and stores their coordinates. Each query checks exactly four directions, so it runs in constant time.

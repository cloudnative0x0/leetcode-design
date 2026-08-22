# 1603 · Design Parking System

**Difficulty:** Easy | **Time:** O(1) per вызов | **Space:** O(1)

---

## Решение хранит три счётчика свободных мест

Никакой структуры под слоты не нужно — важно только количество оставшихся мест каждого типа. Три поля `slotBig`, `slotMedium`, `slotSmall` уменьшаются при постановке машины и не дают уйти в минус.

```go
type ParkingSystem struct {
    slotBig    int
    slotMedium int
    slotSmall  int
}
```

## Почему счётчики, а не массив/срез слотов

- Задаче не нужно знать, какое именно место занято — только «есть свободное или нет».
- Массив слотов дал бы ту же гарантию корректности, но добавил бы O(n) память и лишний перебор при каждой попытке поставить машину.
- Три int'а решают задачу за константную память и константное время на операцию.

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

Просто копирует стартовые значения вместимости в поля структуры.

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

- `carType` напрямую соответствует одному из трёх счётчиков (1 — big, 2 — medium, 3 — small), поэтому `switch` без доп. маппинга.
- Место есть — счётчик уменьшается и функция возвращает `true`.
- Места нет — счётчик не трогаем, возвращаем `false`.

**Edge cases:**
- Вызов с `carType` вне диапазона 1–3 не попадает ни в один `case` и падает на `return false` — паники не будет.
- Вместимость 0 для какого-то типа изначально просто всегда возвращает `false` для него, отдельной обработки не требует.
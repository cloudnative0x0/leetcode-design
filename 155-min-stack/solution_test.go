package _155_min_stack

import "testing"

func TestPushAndTop(t *testing.T) {
	ms := Constructor()
	ms.Push(5)
	if got := ms.Top(); got != 5 {
		t.Fatalf("Top() = %d, хотел 5", got)
	}

	ms.Push(8)
	if got := ms.Top(); got != 8 {
		t.Fatalf("Top() = %d, хотел 8", got)
	}
}

func TestGetMinBasic(t *testing.T) {
	ms := Constructor()
	ms.Push(5)
	ms.Push(2)
	ms.Push(8)

	if got := ms.GetMin(); got != 2 {
		t.Fatalf("GetMin() = %d, хотел 2", got)
	}
}

func TestGetMinAfterPop(t *testing.T) {
	ms := Constructor()
	ms.Push(5)
	ms.Push(2)
	ms.Push(8)

	ms.Pop() // убрали 8
	if got := ms.GetMin(); got != 2 {
		t.Fatalf("после Pop() GetMin() = %d, хотел 2", got)
	}

	ms.Pop() // убрали 2
	if got := ms.GetMin(); got != 5 {
		t.Fatalf("после второго Pop() GetMin() = %d, хотел 5", got)
	}
}

func TestGetMinWithDuplicates(t *testing.T) {
	// Проверяем ключевой случай: несколько одинаковых минимумов подряд.
	// Если бы в Push стояло строгое "<" вместо "<=", тест бы упал.
	ms := Constructor()
	ms.Push(1)
	ms.Push(1)
	ms.Push(1)

	ms.Pop()
	if got := ms.GetMin(); got != 1 {
		t.Fatalf("GetMin() = %d, хотел 1 (дубликаты минимума)", got)
	}

	ms.Pop()
	if got := ms.GetMin(); got != 1 {
		t.Fatalf("GetMin() = %d, хотел 1 (последний дубликат)", got)
	}
}

func TestGetMinDecreasing(t *testing.T) {
	// Каждый следующий элемент меньше предыдущего — все должны попасть в minStack.
	ms := Constructor()
	values := []int{10, 5, 3, 1}
	for _, v := range values {
		ms.Push(v)
	}

	for i := len(values) - 1; i >= 0; i-- {
		if got := ms.GetMin(); got != values[i] {
			t.Fatalf("GetMin() = %d, хотел %d на шаге %d", got, values[i], i)
		}
		ms.Pop()
	}
}

func TestGetMinIncreasing(t *testing.T) {
	// Каждый следующий элемент больше предыдущего — минимум не должен меняться.
	ms := Constructor()
	ms.Push(1)
	ms.Push(2)
	ms.Push(3)
	ms.Push(4)

	if got := ms.GetMin(); got != 1 {
		t.Fatalf("GetMin() = %d, хотел 1", got)
	}

	ms.Pop()
	ms.Pop()
	ms.Pop()

	if got := ms.GetMin(); got != 1 {
		t.Fatalf("после трёх Pop() GetMin() = %d, хотел 1", got)
	}
}

func TestPopUntilEmpty(t *testing.T) {
	ms := Constructor()
	ms.Push(1)
	ms.Push(2)

	ms.Pop()
	ms.Pop()

	// Стек пуст: Pop() не должен паниковать.
	ms.Pop()

	if got := ms.Top(); got != 0 {
		t.Fatalf("Top() на пустом стеке = %d, хотел 0", got)
	}
	if got := ms.GetMin(); got != 0 {
		t.Fatalf("GetMin() на пустом стеке = %d, хотел 0", got)
	}
}

func TestEmptyStackDefaults(t *testing.T) {
	ms := Constructor()

	if got := ms.Top(); got != 0 {
		t.Fatalf("Top() на новом стеке = %d, хотел 0", got)
	}
	if got := ms.GetMin(); got != 0 {
		t.Fatalf("GetMin() на новом стеке = %d, хотел 0", got)
	}
}

func TestNegativeAndZeroValues(t *testing.T) {
	ms := Constructor()
	ms.Push(0)
	ms.Push(-3)
	ms.Push(-1)

	if got := ms.GetMin(); got != -3 {
		t.Fatalf("GetMin() = %d, хотел -3", got)
	}

	ms.Pop() // убрали -1
	if got := ms.GetMin(); got != -3 {
		t.Fatalf("после Pop() GetMin() = %d, хотел -3", got)
	}

	ms.Pop() // убрали -3
	if got := ms.GetMin(); got != 0 {
		t.Fatalf("после второго Pop() GetMin() = %d, хотел 0", got)
	}
}

func TestSingleElement(t *testing.T) {
	ms := Constructor()
	ms.Push(42)

	if got := ms.Top(); got != 42 {
		t.Fatalf("Top() = %d, хотел 42", got)
	}
	if got := ms.GetMin(); got != 42 {
		t.Fatalf("GetMin() = %d, хотел 42", got)
	}

	ms.Pop()
	if got := ms.Top(); got != 0 {
		t.Fatalf("Top() после Pop() = %d, хотел 0", got)
	}
}

func TestFullSequenceFromLeetCode(t *testing.T) {
	// Классический пример из условия задачи LeetCode 155.
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)

	if got := ms.GetMin(); got != -3 {
		t.Fatalf("GetMin() = %d, хотел -3", got)
	}

	ms.Pop()

	if got := ms.Top(); got != 0 {
		t.Fatalf("Top() = %d, хотел 0", got)
	}
	if got := ms.GetMin(); got != -2 {
		t.Fatalf("GetMin() = %d, хотел -2", got)
	}
}

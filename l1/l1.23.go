package l1

// Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.
// Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента (copy(slice[i:], slice[i+1:])) и уменьшить длину слайса на 1.

func SolveL1_23(slice []int, i int) []int {
	if len(slice) == 0 || i >= len(slice) {
		return slice
	}

	copy(slice[i:], slice[i+1:])
	slice[len(slice)-1] = 0
	return slice[:len(slice)-1]
}

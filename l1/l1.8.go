package l1

// Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.
// Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).
// Подсказка: используйте битовые операции (|, &^).

func SolveL1_8(num int64, i int, bit int) int64 {
	if i < 1 || i > 64 {
		panic("invalid bit number")
	}

	if bit != 0 && bit != 1 {
		panic("invalid bit value")
	}

	if (num>>(i-1))&1 == int64(bit) {
		return num
	}

	return num ^ (1 << (i - 1))
}

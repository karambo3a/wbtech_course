package l1

// Разработать программу, которая переворачивает подаваемую на вход строку.
// Например: при вводе строки «главрыба» вывод должен быть «абырвалг».

// Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.), то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).

func SolveL1_19(s string) string {
	runes := []rune(s)
	if len(runes) == 0 || len(runes) == 1 {
		return string(s)
	}

	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}

	return string(runes)
}

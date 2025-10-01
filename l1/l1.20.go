package l1

import "strings"

// Разработать программу, которая переворачивает порядок слов в строке.
// Пример: входная строка:
// «snow dog sun», выход: «sun dog snow».

// Считайте, что слова разделяются одиночным пробелом.
// Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».

func SolveL1_20(s string) string {
	runes := []rune(s)
	var sb strings.Builder
	end := len(runes)

	for i := len(runes) - 1; i >= -1; i-- {
		if i == -1 || runes[i] == ' ' {
			sb.WriteString(string(runes[i+1 : end]))
			if i != -1 {
				sb.WriteByte(' ')
			}
			end = i
		}
	}

	return sb.String()
}

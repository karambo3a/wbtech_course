package l1

import "strings"

// Разработать программу, которая проверяет, что все символы в строке встречаются один раз (т.е. строка состоит из уникальных символов).
// Вывод: true, если все символы уникальны, false, если есть повторения. Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.
// Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.
// Подумайте, какой структурой данных удобно воспользоваться для проверки условия.

func SolveL1_26(s string) bool {
	letters := make(map[rune]struct{})
	lowerCaseS := strings.ToLower(s)
	for _, c := range lowerCaseS {
		_, ok := letters[c]
		if ok {
			return false
		}
		letters[c] = struct{}{}
	}
	return true
}

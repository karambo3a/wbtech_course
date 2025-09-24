package l1

// Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.
// Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.

func SolveL1_12(strings []string) []string {
	m := make(map[string]struct{})
	var res []string

	for _, s := range strings {
	   _, ok := m[s]
	   if !ok {
			m[s] = struct{}{}
			res = append(res, s)
	   }
	}

	return res
}

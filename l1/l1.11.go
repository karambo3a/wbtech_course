package l1

// Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) — т.е. вывести элементы, присутствующие и в первом, и во втором.
// Пример:
// A = {1,2,3}
// B = {2,3,4}
// Пересечение = {2,3}

func SolveL1_11(s1 []int, s2 []int) []int {
	m1 := make(map[int]struct{})
	for _, value := range s1 {
		m1[value] = struct{}{}
	}

	m2 := make(map[int]struct{})
	for _, value := range s2 {
		m2[value] = struct{}{}
	}

	var res []int
	for key := range m1 {
		_, ok := m2[key]
		if ok {
			res = append(res, key)
		}
	}

	return res
}

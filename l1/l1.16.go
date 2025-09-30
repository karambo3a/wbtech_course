package l1

// Реализовать алгоритм быстрой сортировки массива встроенными средствами языка. Можно использовать рекурсию.

// Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел.
// Для выбора опорного элемента можно взять середину или первый элемент.

func SolveL1_16(nums []int) []int {
	return quickSort(nums)
}

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	if len(nums) == 2 {
		if nums[0] <= nums[1] {
			return nums
		}
		return []int{nums[1], nums[0]}
	}

	pivot := nums[len(nums)/2]
	var left, middle, right []int

	for _, n := range nums {
		if n < pivot {
			left = append(left, n)
		} else if n == pivot {
			middle = append(middle, pivot)
		} else {
			right = append(right, n)
		}
	}

	newLeft := quickSort(left)
	newRight := quickSort(right)

	var res []int
	res = append(res, newLeft...)
	res = append(res, middle...)
	res = append(res, newRight...)
	return res
}

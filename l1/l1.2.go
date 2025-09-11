package l1

import "fmt"

func ConcurrentSquare() {
	nums := []int{2, 4, 6, 8, 10}
	squares := make([]int, len(nums))
	ch := make(chan struct{}, len(nums))

	for i := range nums {
		go func() {
			squares[i] = nums[i] * nums[i]
			ch <- struct{}{}
		}()
	}

	for range nums {
		<-ch
	}

	for i, square := range squares {
		fmt.Printf("%d^2 = %d\n", nums[i], square)
	}
}

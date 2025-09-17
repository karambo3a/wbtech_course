package l1

import (
	"fmt"
	"sync"
)

// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа x из массива, во второй – результат операции x*2.
// После этого данные из второго канала должны выводиться в stdout.
// То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка.
// Убедитесь, что чтение из второго канала корректно завершается.

func SolveL1_9(nums []int) {
	chan1 := make(chan int)
	chan2 := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go generate(nums, chan1)
	go square(chan1, chan2)
	go print(chan2, &wg)

	wg.Wait()
}

func generate(nums []int, chan1 chan<- int) {
	for _, n := range nums {
		chan1 <- n
	}

	close(chan1)
}

func square(chan1 <-chan int, chan2 chan<- int) {
	for n := range chan1 {
		chan2 <- n * 2
	}

	close(chan2)
}

func print(chan2 <-chan int, wg *sync.WaitGroup) {
	for n := range chan2 {
		fmt.Println(n)
	}

	wg.Done()
}

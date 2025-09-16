package l1

import (
	"fmt"
	"sync"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала – читать эти значения.
// По истечении N секунд программа должна завершаться.
// Подсказка: используйте time.After или таймер для ограничения времени работы.

func producer(ch chan<- int, done chan struct{}, wg *sync.WaitGroup) {
	for {
		select {
		case <-done:
			fmt.Println("producer done")
			wg.Done()
			return
		case ch <- 1:
			continue
		}
	}
}

func consumer(ch <-chan int, done chan struct{}, wg *sync.WaitGroup) {
	for {
		select {
		case <-done:
			fmt.Println("consumer done")
			wg.Done()
			return
		case <-ch:
			continue
		}
	}
}

func SolveL1_5(N int) {
	ch := make(chan int)
	done := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(ch, done, &wg)
	go consumer(ch, done, &wg)

	<-time.After(time.Duration(N) * time.Second)
	close(done)
	wg.Wait()
}

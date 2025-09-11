// Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).
// Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.
// Подсказка: можно использовать контекст (context.Context) или канал для оповещения о завершении.

package l1

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func Solve2() {
	if len(os.Args) != 2 {
		fmt.Println("wrong number of args")
		return
	}

	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("error converting string to integer")
		return
	}
	if N < 1 {
		fmt.Println("must be at least one worker")
		return
	}

	ch := make(chan int)
	sigCh := make(chan os.Signal, 1)
	doneCh := make(chan struct{})
	signal.Notify(sigCh, syscall.SIGINT)

	for i := range N {
		go Worker2(ch, doneCh, i)
	}

	Producer2(ch, sigCh, doneCh)
	fmt.Println("done")
}

func Worker2(ch <-chan int, doneCh chan struct{}, id int) {
	for {
		select {
		case data, ok := <-ch:
			if !ok {
				return
			}
			fmt.Printf("worker %d: %d\n", id, data)
		case <-doneCh:
			return
		}
	}
}

func Producer2(ch chan int, sigCh chan os.Signal, doneCh chan struct{}) {
	for {
		select {
		case <-sigCh:
			close(doneCh)
			return
		default:
			ch <- 1
		}
	}
}

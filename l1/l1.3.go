package l1

import (
	"fmt"
	"os"
	"strconv"
)

// call in main goroutine
func Solve() {
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
	for i := range N {
		go Worker(ch, i)
	}

	Producer(ch)
}

func Worker(ch <-chan int, id int) {
	for data := range ch {
		fmt.Printf("worker %d: %d\n", id, data)
	}
}

func Producer(ch chan int) {
	for {
		ch <- 1
	}
}

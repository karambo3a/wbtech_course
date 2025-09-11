package l1

import (
	"fmt"
	"os"
	"strconv"
)

func Worker(ch <-chan int, id int) {
	for data := range ch {
		fmt.Printf("worker %d: %d\n", id, data)
	}
}

// call in main goroutine
func Producer() {
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

	for {
		ch <- 1
	}
}

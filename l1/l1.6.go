package l1

import (
	"context"
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.
// Классические подходы: выход по условию, через канал уведомления, через контекст, прекращение работы runtime.Goexit() и др.
// Продемонстрируйте каждый способ в отдельном фрагменте кода.

func SolveL1_6() {
	cancelWithCond()
	cancelWithChannelSend()
	cancelWithChannelClose()
	cancelWithContext()
	cancelWithGoexit()
	cancelWithTimeAfter()
	cancelWithSharedVariable()
}

func cancelWithCond() {
	fmt.Println("start cancelWithCond")

	go func() {
		for i := range 5 {
			fmt.Println(i)
		}
		fmt.Println("done goroutine cancelWithCond")

	}()

	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("done cancelWithCond")
	fmt.Println()
}

func cancelWithChannelSend() {
	fmt.Println("start cancelWithChannelSend")
	done := make(chan struct{})

	go func() {
		i := 0
		for {
			select {
			case <-done:
				fmt.Println("done goroutine cancelWithChannelSend")
				return
			default:
				if i < 5 {
					fmt.Println(i)
					i++
				}
			}
		}
	}()

	time.Sleep(time.Duration(1) * time.Second)
	done <- struct{}{}
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("done cancelWithChannelSend")
	fmt.Println()
}

func cancelWithChannelClose() {
	fmt.Println("start cancelWithChannelClose")
	done := make(chan struct{})

	go func() {
		i := 0
		for {
			select {
			case <-done:
				fmt.Println("done goroutine cancelWithChannelClose")
				return
			default:
				if i < 5 {
					fmt.Println(i)
					i++
				}
			}
		}
	}()

	time.Sleep(time.Duration(1) * time.Second)
	close(done)
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("done cancelWithChannelClose")
	fmt.Println()
}

func cancelWithContext() {
	fmt.Println("start cancelWithContext")

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done goroutine cancelWithContext")
				return
			default:
				if i < 5 {
					fmt.Println(i)
					i++
				}
			}
		}
	}()

	time.Sleep(time.Duration(1) * time.Second)
	cancel()
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("done cancelWithContext")
	fmt.Println()
}

func cancelWithGoexit() {
	fmt.Println("start cancelWithGoexit")

	go func() {
		defer fmt.Println("done goroutine cancelWithGoexit")

		for i := range 5 {
			fmt.Println(i)
		}

		runtime.Goexit()
	}()

	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("done cancelWithGoexit")
	fmt.Println()
}

func cancelWithTimeAfter() {
	fmt.Println("start cancelTimeAfter")

	go func() {
		i := 0
		after := time.After(time.Duration(1) * time.Second)
		for {
			select {
			case <-after:
				fmt.Println("done goroutine cancelTimeAfter")
				return
			default:
				if i < 5 {
					fmt.Println(i)
					i++
				}
				time.Sleep(time.Millisecond)
			}
		}
	}()

	time.Sleep(time.Duration(3) * time.Second)
	fmt.Println("done cancelTimeAfter")
	fmt.Println()
}

func cancelWithSharedVariable() {
	fmt.Println("start cancelWithSharedVariable")

	var stop atomic.Bool
	go func() {
		i := 0
		for {
			if stop.Load() {
				fmt.Println("done goroutine cancelWithSharedVariable")
				return
			}
			if i < 5 {
				fmt.Println(i)
				i++
			}
		}
	}()

	time.Sleep(time.Duration(1) * time.Second)
	stop.Store(true)
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("done cancelWithSharedVariable")
	fmt.Println()
}

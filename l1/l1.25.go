package l1

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep(duration) аналогично встроенной функции time.Sleep, которая приостанавливает выполнение текущей горутины.
// Важно: в отличие от настоящей time.Sleep, ваша функция должна именно блокировать выполнение (например, через таймер или цикл), а не просто вызывать time.Sleep :) — это упражнение.
// Можно использовать канал + горутину, или цикл на проверку времени (не лучший способ, но для обучения).

func sleep(sec int) {
	start := time.Now()
	for time.Since(start) < time.Duration(sec)*time.Second {
	}
}

func SolveL1_25(sec int) {
	fmt.Println("before sleep")
	sleep(sec)
	fmt.Printf("after %d sec sleep\n", sec)
}

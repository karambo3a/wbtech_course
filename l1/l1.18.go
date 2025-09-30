package l1

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин).
// По завершению программы структура должна выводить итоговое значение счётчика.

// Подсказка: вам понадобится механизм синхронизации, например, sync.Mutex или sync/Atomic для безопасного инкремента.

type concurrentCounter struct {
	counter int
	rwMutex sync.RWMutex
}

func NewConcurrentCounter() *concurrentCounter {
	return &concurrentCounter{counter: 0}
}

func (cc *concurrentCounter) Inc() {
	cc.rwMutex.Lock()
	defer cc.rwMutex.Unlock()
	cc.counter++
}

func (cc *concurrentCounter) Get() int {
	cc.rwMutex.RLock()
	defer cc.rwMutex.RUnlock()
	return cc.counter
}

func SolveL1_18() {
	counter := NewConcurrentCounter()
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		for range 10 {
			counter.Inc()
		}
		wg.Done()
	}()

	go func() {
		for range 10 {
			counter.Inc()
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(counter.Get())
}

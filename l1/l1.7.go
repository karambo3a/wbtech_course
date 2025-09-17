package l1

import (
	"fmt"
	"sync"
)

// Реализовать безопасную для конкуренции запись данных в структуру map.
// Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).
// Проверьте работу кода на гонки (util go run -race).

type concurrentMap[T comparable, V any] struct {
	m       map[T]V
	rwMutex sync.RWMutex
}

func NewConcurrentMap[T comparable, V any]() *concurrentMap[T, V] {
	return &concurrentMap[T, V]{
		m: make(map[T]V),
	}
}

func (cm *concurrentMap[T, V]) Put(key T, value V) {
	cm.rwMutex.Lock()
	defer cm.rwMutex.Unlock()

	cm.m[key] = value
}

func (cm *concurrentMap[T, V]) Get(key T) (value V, ok bool) {
	cm.rwMutex.RLock()
	defer cm.rwMutex.RUnlock()

	value, ok = cm.m[key]
	return
}

func SolveL1_7() {
	cm := NewConcurrentMap[int, int]()

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		for i := range 100 {
			value, ok := cm.Get(i)
			fmt.Printf("GET: value=%d  ok=%t\n", value, ok)
		}
		wg.Done()
	}()

	go func() {
		for i := range 100 {
			cm.Put(i, i)
			fmt.Printf("PUT: value=%d\n", i)
		}
		wg.Done()
	}()

	wg.Wait()
}

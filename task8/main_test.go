package main

import (
	"sync"
	"testing"
	"time"
)

func NewCustomWaitGroup() *CustomWaitGroup {
	wg := &CustomWaitGroup{}
	wg.cond = sync.NewCond(&wg.lock)
	return wg
}

func TestAdd(t *testing.T) {
	wg := NewCustomWaitGroup()

	someNum := 5
	wg.Add(someNum)

	if wg.counter != someNum {
		t.Errorf("Не правильная работа функции Add")
	}
}

func TestDone(t *testing.T) {
	wg := NewCustomWaitGroup()
	someNum := 5
	wg.Add(someNum)

	wg.Done()

	if wg.counter != someNum-1 {
		t.Errorf("Не правильная работа функции Done")
	}
}

func TestWait(t *testing.T) {
	wg := NewCustomWaitGroup()
	var counter int

	numWorkers := 5

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(100 * time.Millisecond)
			counter++
			wg.Done()
		}()
	}

	wg.Wait()

	if counter != numWorkers {
		t.Errorf("Ожидалось завершённых горутин %d, получили %d", numWorkers, counter)
	}
}

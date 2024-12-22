package main

import "sync"

type CustomWaitGroup struct {
	counter int
	lock    sync.Mutex
	cond    *sync.Cond
}

func (wg *CustomWaitGroup) Add(delta int) {
	wg.lock.Lock()
	defer wg.lock.Unlock()
	wg.counter += delta
	if wg.counter == 0 {
		wg.cond.Broadcast()
	}
}

func (wg *CustomWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *CustomWaitGroup) Wait() {
	wg.lock.Lock()
	defer wg.lock.Unlock()
	for wg.counter > 0 {
		wg.cond.Wait()
	}
}

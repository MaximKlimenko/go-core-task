package main

import (
	"sync"
)

func merge(chs ...<-chan int) <-chan int {
	resultChannel := make(chan int, 1)

	wg := &sync.WaitGroup{}
	for _, ch := range chs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for c := range ch {
				resultChannel <- c
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	return resultChannel
}

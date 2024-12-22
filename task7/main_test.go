package main

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)

	go func() {
		ch1 <- 1
		ch1 <- 3
		close(ch1)
	}()

	go func() {
		ch2 <- 2
		ch2 <- 4
		close(ch2)
	}()

	go func() {
		ch3 <- 0
		ch3 <- 5
		close(ch3)
	}()
	sum := 0
	for ch := range merge(ch1, ch2, ch3) {
		sum += ch
		fmt.Println(ch)
	}
	if sum != 15 {
		t.Errorf("Не верная сумма из каналов") //Хз как еще проверить прввильность
	}
}

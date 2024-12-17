package main

import (
	"fmt"
	"math/rand"
)

func randomNumberGenerator(ch chan int) {
	num := rand.Intn(100)
	ch <- num
}
func main() {
	ch := make(chan int)

	go randomNumberGenerator(ch)

	num := <-ch
	fmt.Printf("Generated number: %d\n", num)
}

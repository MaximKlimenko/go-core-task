package main

import "fmt"

func Conveyor(chIn <-chan uint8, chOut chan<- float64) {
	for num := range chIn {
		chOut <- float64(num) * float64(num) * float64(num)
	}

	close(chOut)
}

func main() {
	inputChannel := make(chan uint8)
	outputChannel := make(chan float64)

	go Conveyor(inputChannel, outputChannel)

	go func() {
		for i := uint8(1); i <= 10; i++ {
			inputChannel <- i
		}
		close(inputChannel)
	}()

	fmt.Println("Results:")
	for result := range outputChannel {
		fmt.Println(result)
	}
}

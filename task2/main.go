package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//1
	originalSlice := []int{}
	for i := 0; i < 10; i++ {
		originalSlice = append(originalSlice, rand.Intn(100))
	}
	fmt.Println("Оригинальный слайс:", originalSlice)
	fmt.Println("Чётный слайс", sliceExample(originalSlice))
	a := addElements(originalSlice, 2)
	fmt.Println("Слайс с доп. элементом", a)
	copy := copySlice(originalSlice)
	copy[0] = 101
	fmt.Println("Копированный слайс с изм 0 элементом", copy)
	fmt.Println("Слайс с удалённым элементом:", removeElement(originalSlice, 0))
	fmt.Println("Оригинальный слайс остался целым:", originalSlice)
}

// 2
func sliceExample(slice []int) []int {
	ansSlice := []int{}
	for _, num := range slice {
		if num%2 == 0 {
			ansSlice = append(ansSlice, num)
		}
	}

	return ansSlice
}

// 3
func addElements(slice []int, num int) []int {
	ansSlice := make([]int, len(slice))
	copy(ansSlice, slice) //Копируем слайс, потому-что иначе он будет передаваться по ссылке, а не по значению
	ansSlice = append(ansSlice, num)
	return ansSlice //На выходе имеем новый слайс, не связанный со старым
}

// 4
func copySlice(slice []int) []int { //Можно использовать функцию copy
	ansSlice := make([]int, len(slice))
	for i := range slice {
		ansSlice[i] = slice[i]
	}

	return ansSlice
}

// 5
func removeElement(slice []int, index int) []int {
	ansSlice := []int{}
	for i := range slice {
		if i == index {
			continue
		}

		ansSlice = append(ansSlice, slice[i])
	}

	return ansSlice
}

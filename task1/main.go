package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	//1
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64
	//2
	fmt.Printf("%T\n%T\n%T\n%T\n%T\n%T\n%T\n", numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	//3
	bigString := fmt.Sprintln(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println(bigString)
	//4
	runes := []rune(bigString)
	//5
	salt := "go-2024"
	midIndex := len(runes) / 2
	runesWithSalt := append(runes[:midIndex], append([]rune(salt), runes[midIndex:]...)...)

	hasher := sha256.New()
	hasher.Write([]byte(string(runesWithSalt)))
	hashed := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashed)

	fmt.Println(hashString)
}

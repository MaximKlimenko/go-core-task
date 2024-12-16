package main

import (
	"fmt"
	"testing"
)

func TestConcatenatedString(t *testing.T) {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	expected := "42 42 42 3.14 Golang true (1+2i)"
	concatenated := fmt.Sprintf("%d %d %d %.2f %s %t %v", numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)

	if concatenated != expected {
		t.Errorf("Expected concatenated string to be '%s', got '%s'", expected, concatenated)
	}
}

package main

import (
	"testing"
)

func TestSliceExample(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}

	result := sliceExample(originalSlice)

	if !Equal(result, expected) {
		t.Errorf("Неправильная работа функции")
	}
}

func TestAddElements(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	num := 11
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	result := addElements(originalSlice, num)

	if !Equal(result, expected) {
		t.Errorf("Неправильная работа функции")
	}
}

func TestCopySlice(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := copySlice(originalSlice)

	if !Equal(result, expected) {
		t.Errorf("Неправильная работа функции")
	}
	result[0] = 123
	if Equal(result, expected) {
		t.Errorf("Изменения в оригинальном слайсе влияют на его копию.")
	}
}

func TestRemoveElement(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index := 5
	expected := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}

	result := removeElement(originalSlice, index)

	if !Equal(result, expected) {
		t.Errorf("Неправильная работа функции")
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

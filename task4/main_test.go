package main

import (
	"reflect"
	"testing"
)

func TestNeMain(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	expected := []string{"apple", "cherry", "43", "lead", "gno1"}
	actual := neMain(slice1, slice2)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Не правильная работа функции")
	}
}

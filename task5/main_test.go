package main

import (
	"reflect"
	"testing"
)

func TestEqNums(t *testing.T) { //Порядок важен
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	expected := []int{3, 64}
	exsists, actual := eqNums(a, b)
	if !reflect.DeepEqual(actual, expected) || !exsists {
		t.Errorf("не работает")
	}
}

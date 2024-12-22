package main

import (
	"reflect"
	"testing"
)

var newMap StringIntMap

func TestAdd(t *testing.T) {
	newMap.storage = make(map[string]int) // Инициализация мапы
	expected := make(map[string]int)
	expected["test"] = 123
	expected["test2"] = 321

	newMap.Add("test", 123)
	newMap.Add("test2", 321)

	if newMap.storage["test"] != expected["test"] {
		t.Errorf("Значения не совпадают")
	}
}

func TestRemove(t *testing.T) {
	newMap.Remove("test")

	_, exists := newMap.storage["test"]
	if exists {
		t.Errorf("Удаление не произошло")
	}
}

func TestCopy(t *testing.T) {
	copyMap := newMap.Copy()

	if !reflect.DeepEqual(newMap.storage, copyMap) {
		t.Errorf("Мапа не скопировалась")
	}

	//Можно добавить проверку на то, что мапа скопировалась по значению а не по ссылке
}

func TestExists(t *testing.T) {
	if !newMap.Exists("test2") { //Ранее было добавлено данное значение в мапу
		t.Errorf("Пары ключ значение test2:321 не существует")
	}
}

func TestGet(t *testing.T) {
	expected := 321
	actual, exists := newMap.Get("test2")

	if expected != actual || !exists {
		t.Errorf("Ошибка получения значения")
	}
}

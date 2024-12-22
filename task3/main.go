package main

import "fmt"

type StringIntMap struct {
	storage map[string]int
}

func (m *StringIntMap) Add(key string, value int) {
	m.storage[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.storage, key)
}

func (m StringIntMap) Copy() map[string]int {
	newStorage := make(map[string]int)
	for key, value := range m.storage {
		newStorage[key] = value
	}

	return newStorage
}

func (m StringIntMap) Exists(key string) bool {
	_, exists := m.storage[key]

	return exists
}

func (m StringIntMap) Get(key string) (int, bool) {
	value, exists := m.storage[key]
	if exists {
		return value, true
	}

	return 0, false
}

func main() {
	var maap StringIntMap
	maap.storage = make(map[string]int) // Инициализация мапы
	maap.Add("Hello", 1)
	fmt.Println(maap.storage)
}

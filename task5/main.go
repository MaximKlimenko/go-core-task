package main

func eqNums(a, b []int) (bool, []int) { //По идее можно придумать алгоритм получше, и получить не квадратичную сложность, а напрмиер линейную или даже log
	var ansSlice []int
	consists := false
	for _, numA := range a {
		for _, numB := range b {
			if numA == numB {
				consists = true
				ansSlice = append(ansSlice, numA)
			}
		}
	}

	return consists, ansSlice
}

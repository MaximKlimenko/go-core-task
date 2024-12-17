package main

func neMain(slice1, slice2 []string) []string {
	var ansSlice []string
	var index1, index2 int
	lenght1 := len(slice1)
	lenght2 := len(slice2)

	for index1 < lenght1 {
		if slice1[index1] == slice2[index2] {
			index1++
			continue
		}

		if index2+1 == lenght2 {
			ansSlice = append(ansSlice, slice1[index1])
			index2 = 0
			index1++
		} else {
			index2++
		}

	}

	return ansSlice
}

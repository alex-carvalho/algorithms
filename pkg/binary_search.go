package main

func search(array []int, searchValue int) bool {
	min := 0
	max := len(array) - 1

	for max >= min {
		middle := (min + max) / 2

		if array[middle] == searchValue {
			return true
		} else if array[middle] > searchValue {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}

	return false
}

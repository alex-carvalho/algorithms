package main

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	lower := []int{}
	bigger := []int{}
	for _, v := range arr[1:] {
		if v <= pivot {
			lower = append(lower, v)
		} else {
			bigger = append(bigger, v)
		}
	}
	return append(append(quickSort(lower), pivot), quickSort(bigger)...)
}

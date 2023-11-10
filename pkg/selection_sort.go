package main

func findSmallest(arr []int) int {
	smaller := arr[0]
	smallerIndex := 0

	for i, v := range arr {
		if v < smaller {
			smaller = v
			smallerIndex = i
		}
	}

	return smallerIndex
}

func selectionSort(arr []int) []int {
	buffer := []int{}
	for range arr {
		smallerIndex := findSmallest(arr)
		buffer = append(buffer, arr[smallerIndex])

		// remove index from arr
		arr = append(arr[:smallerIndex], arr[smallerIndex+1:]...)
	}

	return buffer
}

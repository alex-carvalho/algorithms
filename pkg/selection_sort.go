package main

import "fmt"

func findSmallest(arr []int) int {
	smaller := arr[0]
	smallerIndex := 0

	for i, v := range(arr){
		if v < smaller {
			smaller = v
			smallerIndex = i
		}
	}

	return smallerIndex
}

func sort(arr []int) []int {
	buffer := []int{}
	for i := range(arr) {
		print(i)
		smallerIndex := findSmallest(arr)
		buffer = append(buffer, arr[smallerIndex])

		// remove index from arr
		arr = append(arr[:smallerIndex], arr[smallerIndex+1:]...)
	}
	
	return buffer
}

func main() {
	print(fmt.Printf("%v", sort([]int{8,2,5,3,4,1})))
}
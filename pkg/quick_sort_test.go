package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestQuickSortList(t *testing.T) {
	unsorted := [...]int{7, 1, 5, 3, 2, 4, 6}
	expected := [...]int{1, 2, 3, 4, 5, 6, 7}

	result := quickSort(unsorted[:])

	if !slices.Equal(expected[:], result) {
		t.Error("Value not sorted")
	}
}

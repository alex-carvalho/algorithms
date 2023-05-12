package main

import "testing"

func TestFindMiddleValue(t *testing.T) {
	searchValue := 3
	array := [...]int{0, 1, 2, 3, 4, 5, 6}

	result := search(array[:], searchValue)

	if !result {
		t.Error("Value not found")
	}

}

func TestFindLastValue(t *testing.T) {
	searchValue := 6
	array := [...]int{0, 1, 2, 3, 4, 5, 6}

	result := search(array[:], searchValue)

	if !result {
		t.Error("Value not found")
	}

}

func TestFindFirstValue(t *testing.T) {
	searchValue := 6
	array := [...]int{0, 1, 2, 3, 4, 5, 6}

	result := search(array[:], searchValue)

	if !result {
		t.Error("Value not found")
	}

}

func TestNotFoundValue(t *testing.T) {
	searchValue := 10
	array := [...]int{0, 1, 2, 3, 4, 5, 6}

	result := search(array[:], searchValue)

	if result {
		t.Error("Value found")
	}

}
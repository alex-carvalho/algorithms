package main

import (
	"testing"
)

func TestDijkstra(t *testing.T) {
	// 		  		 A
	//      	↗6	 ↑	↘1
	// START		 3		END
	//      	↘2	 ↑	↗5
	// 		  		 B
	//
	// fast = START → B → A → END
	//               2 +  3 +  1 = 6

	g := make(map[string]map[string]int)
	g["START"] = make(map[string]int)
	g["START"]["A"] = 6
	g["START"]["B"] = 2

	g["A"] = make(map[string]int)
	g["A"]["END"] = 1

	g["B"] = make(map[string]int)
	g["B"]["A"] = 3
	g["B"]["END"] = 5

	g["END"] = make(map[string]int)

	costs := make(map[string]int)
	costs["A"] = 6
	costs["B"] = 2
	costs["END"] = MaxInt

	parents := make(map[string]string)
	parents["A"] = "START"
	parents["B"] = "START"

	result := dijkstra(g, costs, parents)
	expected := "START B A END = 6"

	if expected != result {
		t.Errorf("found = '%s' expected '%s'", result, expected)
	}
}

func TestDijkstra2(t *testing.T) {
	// 		  		 A		→4	   C
	//      	↗5	 					↘3
	// START		 8↑		2↘	  6↓		END
	//      	↘2	 					↗1
	// 		  		 B		→7	   D
	//
	// fast = START → A → D → END
	//                5 + 2 + 1 = 8

	g := make(map[string]map[string]int)
	g["START"] = make(map[string]int)
	g["START"]["A"] = 5
	g["START"]["B"] = 2

	g["A"] = make(map[string]int)
	g["A"]["C"] = 4
	g["A"]["D"] = 2

	g["B"] = make(map[string]int)
	g["B"]["A"] = 8
	g["B"]["D"] = 7

	g["C"] = make(map[string]int)
	g["C"]["D"] = 6
	g["C"]["END"] = 3

	g["D"] = make(map[string]int)
	g["D"]["END"] = 1

	g["END"] = make(map[string]int)

	costs := make(map[string]int)
	costs["A"] = 5
	costs["B"] = 2
	costs["C"] = MaxInt
	costs["D"] = MaxInt
	costs["END"] = MaxInt

	parents := make(map[string]string)
	parents["A"] = "START"
	parents["B"] = "START"

	result := dijkstra(g, costs, parents)
	expected := "START A D END = 8"

	if expected != result {
		t.Errorf("found = '%s' expected '%s'", result, expected)
	}
}

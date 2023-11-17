package main

import (
	"strconv"

	"golang.org/x/exp/slices"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func dijkstra(g map[string]map[string]int, costs map[string]int, parents map[string]string) string {
	visited := make([]string, 1)

	node := findLowCost(costs, visited)

	for node != "" {
		cost := costs[node]
		neighbors := g[node]

		for n, value := range neighbors {
			newCost := cost + value
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}

		visited = append(visited, node)
		node = findLowCost(costs, visited)
	}

	result := ""
	n := "END"
	for n != "" {
		result = n + " " + result
		n = parents[n]
	}

	return result + "= " + strconv.Itoa(costs["END"])
}

func findLowCost(costs map[string]int, visited []string) string {
	lowCost := MaxInt
	nodeLowCost := ""

	for node, cost := range costs {
		if cost < lowCost && !slices.Contains(visited, node) {
			lowCost = cost
			nodeLowCost = node
		}
	}
	return nodeLowCost
}

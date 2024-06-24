package main

import (
	"fmt"
	"math"
)

func findLowestCostNode(costsFromStart map[string]uint16, processed []string) string {
	var lowestCost uint16 = math.MaxUint16
	lowestCostNode := ""

	for node := range costsFromStart {
		cost := costsFromStart[node]
		nodeInProcessed := false
		if cost < lowestCost {

			for i := 0; i < len(processed); i++ {
				if node == processed[i] {
					nodeInProcessed = true
				}
			}

			if nodeInProcessed != true {
				lowestCost = cost
				lowestCostNode = node
			}
		}
	}

	return lowestCostNode
}

func main() {

	var graf = map[string]map[string]uint16{
		"start": {"a": 5, "b": 2, "c": math.MaxUint16, "d": math.MaxUint16, "fin": math.MaxUint16},
		"a":     {"c": 4, "d": 2, "fin": math.MaxUint16},
		"b":     {"a": 8, "d": 7, "fin": math.MaxUint16},
		"c":     {"fin": 3},
		"d":     {"fin": 1},
	}

	var costsFromStart = map[string]uint16{

		"a":   5,
		"b":   2,
		"c":   math.MaxUint16,
		"d":   math.MaxUint16,
		"fin": math.MaxUint16,
	}

	var parents = map[string]string{

		"a":   "start",
		"b":   "start",
		"c":   "",
		"d":   "",
		"fin": "",
	}

	var processed []string

	node := findLowestCostNode(costsFromStart, processed)
	cost := costsFromStart[node]
	neighbors := graf[node]

	for node != "" {
		cost = costsFromStart[node]
		neighbors = graf[node]
		for k := range neighbors {
			new_cost := cost + neighbors[k]
			if costsFromStart[k] > new_cost {
				costsFromStart[k] = new_cost
				parents[k] = node
			}
		}
		processed = append(processed, node)
		node = findLowestCostNode(costsFromStart, processed)
	}

	fmt.Println(graf)
	fmt.Println(graf["start"])
	fmt.Println(graf["start"]["a"])
	fmt.Println(costsFromStart)
}

package main

import (
	"fmt"
)

func bfsearch(name string) bool {
	var nameInt []int32 = []int32(name)
	var lastLetter int32 = nameInt[len(name)-1]
	if lastLetter == 109 {
		return true
	} else {
		return false
	}
}

func main() {
	var graf = map[string][]string{
		"you":   {"Alice", "Bob", "Clare"},
		"Alice": {"you", "Peggi"},
		"Bob":   {"Anug", "Peggi", "you"},
		"Clare": {"Jonni", "you"},
		"Peggi": {"Alice", "Bob"},
		"Anug":  {"Bob"},
		"Jonni": {"Clare"},
		"Tom":   {"Clare"},
	}

	var nameQueue []string = graf["you"]
	var searchedNames []string

	for i := 0; i < len(nameQueue); i++ {
		searchedNames = append(searchedNames, nameQueue[0])

		for i := 0; i < len(searchedNames); i++ {
			if nameQueue[0] != searchedNames[i] {
				if bfsearch(nameQueue[0]) {
					fmt.Printf("%s is the One!", nameQueue[0])
					break
				} else {
					i--
					nameQueue = append(nameQueue, graf[nameQueue[0]]...)
					nameQueue = nameQueue[1:]
					fmt.Println(nameQueue)
				}
			}
		}

	}

}

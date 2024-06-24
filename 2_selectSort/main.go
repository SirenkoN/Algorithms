package main

import (
	"fmt"
)

const MAX_INT32_VALUE = 2147483647

func findSmallestIndex(slice []int32) (smallestIndex int32) {
	var smallest int32 = MAX_INT32_VALUE
	for i := 0; i < len(slice); i++ {
		if slice[i] < smallest {
			smallestIndex = int32(i)
			smallest = slice[smallestIndex]
		}
	}
	return smallestIndex
}

func selectSort(unsortedSlice []int32) (sortedSlice []int32) {
	var minIndex int32 = 0
	for i := 0; i < len(unsortedSlice); i++ {
		minIndex = findSmallestIndex(unsortedSlice)
		sortedSlice = append(sortedSlice, unsortedSlice[minIndex])
		unsortedSlice[minIndex] = MAX_INT32_VALUE
	}
	return sortedSlice
}

func main() {
	var slice []int32 = []int32{-100, 500, 6, 8, -25, 100, -200, 500, 800}
	fmt.Println("Unsorted slice = ", slice)
	fmt.Println("Sorted slice = ", selectSort(slice))
	fmt.Println("Unsorted slice = ", slice)
	//sort.Ints(slice)
}

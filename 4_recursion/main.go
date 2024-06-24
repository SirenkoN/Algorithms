package main

import "fmt"

func sliceRecursive(slice []int32) int32 {
	if len(slice) == 0 {
		return 0
	} else {
		return slice[0] + sliceRecursive(slice[1:])
	}
}

func elementsInSliceRecursibe(slice []int32) int32 {
	if len(slice) == 0 {
		return 0
	} else {
		return 1 + elementsInSliceRecursibe(slice[1:])
	}
}

func biggestInSliceRecurcive(slice []int32) int32 {
	if len(slice) == 1 {
		return slice[0]
	} else {
		if slice[0] > slice[1] {
			slice[1] = slice[0]
			return biggestInSliceRecurcive(slice[1:])
		} else {
			return biggestInSliceRecurcive(slice[1:])
		}
	}
}

func qsortArray(slice []int32) []int32 {
	var less, greater, ret []int32

	if len(slice) < 2 {
		return slice
	} else {
		pivot := slice[0]
		for i := 1; i < len(slice); i++ {
			if slice[i] <= pivot {
				less = append(less, slice[i])
			} else if slice[i] > pivot {
				greater = append(greater, slice[i])
			}
		}

		ret = append(ret, qsortArray(less)...)
		ret = append(ret, pivot)
		ret = append(ret, qsortArray(greater)...)
		return ret
	}

}

func main() {
	var slice []int32 = []int32{5, 7, 6, 15, 10, 35, 11, 7, 12, -15, 21}
	fmt.Println("Slice = ", slice)
	fmt.Println("Sorted slice = ", qsortArray(slice))
	fmt.Println("Sum of slice = ", sliceRecursive(slice))
	fmt.Println("Elements in slice = ", elementsInSliceRecursibe(slice))
	fmt.Println("Biggest element = ", biggestInSliceRecurcive(slice))

}

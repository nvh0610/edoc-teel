package main

import (
	"math"
)

func maxDistance(arrays [][]int) int {
	minNum := int(math.Pow(10, 5))
	maxNum := -int(math.Pow(10, 5))
	numBreak := 0
	for index, array := range arrays {
		if minNum > array[0] {
			minNum = array[0]
			numBreak = index
		}

	}

	for index, array := range arrays {
		if index == numBreak {
			continue
		}

		if maxNum < array[len(array)-1] {
			maxNum = array[len(array)-1]
		}
	}
	result1 := maxNum - minNum

	maxNum = -int(math.Pow(10, 5))
	numBreak = 0
	minNum = int(math.Pow(10, 5))
	for index, array := range arrays {
		if maxNum < array[len(array)-1] {
			maxNum = array[len(array)-1]
			numBreak = index
		}
	}

	for index, array := range arrays {
		if index == numBreak {
			continue
		}

		if minNum > array[0] {
			minNum = array[0]
		}

	}

	result2 := maxNum - minNum
	if result1 >= result2 {
		return result1
	}

	return result2
}

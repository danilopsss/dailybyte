package main

import "math"


type TwoSum struct {
	slice []int
	k int
}

func IsSumMatching(a int, b int, k int) bool {
	return a + b == k
}

func (ts TwoSum) hasSum() bool {
	for i := 0; i < len(ts.slice); i++ {
		for j := i + 1; j < len(ts.slice); j++ {
			if IsSumMatching(ts.slice[i], ts.slice[j], ts.k) {
				return true
			}
		}
	}
	return false
}

func (ts TwoSum) hasSumWithNoNestedLoop() bool {
	startingIndex := 0
	movingPoint := len(ts.slice) - 1
	
	loopSize := int(math.Pow(float64(len(ts.slice) - 1), 2))

	for i := 0; i < loopSize; i++ {
		if IsSumMatching(ts.slice[startingIndex], ts.slice[movingPoint], ts.k) {
			return true
		}
		if startingIndex == movingPoint {
			startingIndex = 0
			movingPoint = startingIndex + 1
		} else {
			startingIndex++
		}
	}
	return false
}
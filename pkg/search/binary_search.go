package search

import (
	"math"
)

// BinarySearch implements binary search - O(log n).
// Data must be ordered in order for this search to work.
// https://frontendmasters.com/courses/algorithms/binary-search-algorithm/
func BinarySearch(a []int, x int) bool {
	low := 0
	high := len(a)

	for low < high {
		m := int(math.Floor(float64(low) + float64(high-low)/2))

		mv := a[m]
		if x == mv {
			return true
		} else if mv > x {
			high = m
		} else {
			low = m + 1
		}
	}

	return false
}

package search

import (
	"math"
)

// TwoCrystalBalls When given two crystal balls that will break if dropped from a high enough distance,
// determine the exact spot in which it will break in the most optimized way.
// This segment demonstrates breaking down a search problem without using a linear search.
func TwoCrystalBalls(breaks []bool) int {
	jmpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := jmpAmount
	for ; i < len(breaks); i += jmpAmount {
		if breaks[i] {
			break
		}
	}

	i -= jmpAmount
	for j := 0; j <= jmpAmount && i < len(breaks); j, i = j+1, i+1 {
		if breaks[i] {
			return i
		}
	}

	return -1
}

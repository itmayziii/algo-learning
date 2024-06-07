package search

import (
	"log"
	"math"
)

// TwoCrystalBalls When given two crystal balls that will break if dropped from a high enough distance,
// determine the exact spot in which it will break in the most optimized way.
// This segment demonstrates breaking down a search problem without using a linear search.
func TwoCrystalBalls(breaks []bool) int {
	log.Printf("arr: %v\n", breaks)
	jmpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	log.Printf("jmpAmount: %d\n", jmpAmount)

	i := jmpAmount
	didBreak := -1
	for ; i < len(breaks); i += jmpAmount {
		if breaks[i] {
			log.Printf("1st ball breaking on: %d\n", i)
			didBreak = i
			break
		}
	}

	i -= jmpAmount
	for j := 0; j < jmpAmount && i < len(breaks); j, i = j+1, i+1 {
		log.Printf("manually searching through with i: %d, j: %d\n", i, j)
		if breaks[i] {
			log.Printf("2nd ball breaking on: %d\n", i)
			return i
		}
	}

	if didBreak != -1 {
		return didBreak
	}

	log.Printf("no ball was broken for arr: %v\n", breaks)
	return -1
}

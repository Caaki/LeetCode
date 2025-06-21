package main

import (
	"math"
	"slices"
)

func main() {}

func minimumDeletions(word string, k int) int {

	values := make(map[rune]int, 0)

	for _, v := range word {
		values[v]++
	}

	frequenc := make([]int, 0)

	for _, v := range values {
		frequenc = append(frequenc, v)
	}

	slices.Sort(frequenc)
	minDif := math.MaxInt

	for i := 0; i < len(frequenc); i++ {
		cur := frequenc[i]
		temp := 0
		for j := 0; j < len(frequenc); j++ {
			if cur > frequenc[j] {
				temp += frequenc[j]
				continue
			}

			if cur+k < frequenc[j] {
				temp += frequenc[j] - (cur + k)
				continue
			}
		}
		if temp < minDif {
			minDif = temp
		}
	}

	return minDif
}

package main

import (
	"slices"
)

func main() {}

func pivotArray(nums []int, pivot int) []int {
	less := []int{}
	greater := []int{}
	equal := []int{}

	for _, v := range nums {
		if v == pivot {
			equal = append(equal, v)
			continue
		}
		if v > pivot {
			greater = append(greater, v)
			continue
		}
		less = append(less, v)
	}

	return slices.Concat(less,
		equal,
		greater,
	)
}

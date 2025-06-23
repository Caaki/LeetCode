package main

import "slices"

func main() {}

func findKDistantIndices(nums []int, key int, k int) []int {
	n := len(nums)
	rVals := make([]int, 0)
	for i, _ := range nums {
		j := i - k
		if j < 0 {
			j = 0
		}
		for ; j-i <= k && j < n; j++ {
			if nums[j] != key {
				continue
			} else {
				rVals = append(rVals, i)
				break
			}
		}
	}
	slices.Sort(rVals)

	return rVals
}

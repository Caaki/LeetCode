package main

import "slices"

func findLHS(nums []int) int {

	slices.Sort(nums)

	min := nums[0]
	max := nums[0]

	l := 0
	r := 0

	maxLen := 0

	for r < len(nums) {
		if l > r {
			r = l
		}

		max = nums[r]
		if max-min == 1 {
			if maxLen < r-l+1 {
				maxLen = r - l + 1
			}
			r++
			continue
		}

		if max-min == 0 {
			r++
			continue
		}

		if max-min > 1 {
			l++
			min = nums[l]
		}
	}

	return maxLen

}

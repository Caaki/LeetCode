package main

import "math"

func canStealKHouses(nums []int, k, capability int) bool {
	count := 0
	for i := 0; i < len(nums); {
		if nums[i] <= capability {
			count++
			i += 2
		} else {
			i++
		}
	}
	return count >= k
}

func minCapability(nums []int, k int) int {
	l := 1
	r := math.MaxInt32

	for l < r {
		mid := l + (r-l)/2
		if canStealKHouses(nums, k, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

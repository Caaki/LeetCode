package main

import "os"

func main() {}

/*
1 <= nums.lenght <= 10^5
1 <= nums[i], k<= 10^9

*/

func countGood(nums []int, k int) int64 {
	values := make(map[int]int, 0)

	if len(nums) < 2 {
		return 0
	}
	pairCount := 0
	pairCount += addPairs(nums[0], values)
	result := int64(0)

	for l, r := 0, 0; l < len(nums)-1; {
		if r < len(nums)-1 && pairCount < k {
			r++
			pairCount += addPairs(nums[r], values)
			if pairCount >= k {
				result += int64(len(nums) - r)
			}
			continue
		}
		pairCount -= reducePairs(nums[l], values)

		if pairCount >= k {
			if r == len(nums) {
				result++
			} else {
				result += int64(len(nums) - r)
			}
		}
		l++
	}
	return result
}

func reducePairs(n int, vals map[int]int) int {
	v := vals[n]
	vals[n] -= 1
	v1 := vals[n]
	return (v * (v - 1) / 2) - (v1 * (v1 - 1) / 2)
}

func addPairs(n int, vals map[int]int) int {
	v1 := vals[n]
	vals[n] += 1
	v := vals[n]
	if v < 1 {
		return 0
	}
	return (v * (v - 1) / 2) - (v1 * (v1 - 1) / 2)
}

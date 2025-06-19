package main

import "sort"

func main() {}

func partitionArray(nums []int, k int) int {

	sort.Ints(nums)

	count := 1
	current := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]-current > k {
			count++
			current = nums[i]
		}
	}

	return count
}

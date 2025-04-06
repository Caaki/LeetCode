package main

import (
	"slices"
)

func main() {}

func largestDivisibleSubset(nums []int) []int {
	slices.Sort(nums)
	cache := make(map[int][]int, 0)
	result := []int{}
	for i := 0; i < len(nums); i++ {
		temp := dfs(i, nums, cache)
		if len(result) < len(temp) {
			result = temp
		}
	}

	return result
}

func dfs(i int, nums []int, cache map[int][]int) []int {
	if i == len(nums) {
		return []int{}
	}

	if _, ok := cache[i]; ok {
		return cache[i]
	}
	res := []int{nums[i]}

	for j := i + 1; j < len(nums); j++ {
		if nums[j]%nums[i] == 0 {
			temp := append([]int{nums[i]}, dfs(j, nums, cache)...)
			if len(temp) > len(res) {
				res = temp
			}
		}
	}
	cache[i] = res

	return res
}

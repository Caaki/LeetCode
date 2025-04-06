package main

import "fmt"

func main() {
	fmt.Print(subsetXORSum([]int{3, 4, 5, 6, 7, 8}))
}

func subsetXORSum(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	return dfs(0, 0, nums)
}

func dfs(i int, val int, nums []int) int {
	if i >= len(nums) {
		return 0
	}
	v := dfs(i+1, val, nums)
	v2 := val ^ nums[i] + dfs(i+1, val^nums[i], nums)
	return v + v2
}

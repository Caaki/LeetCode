package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print([]int{-1, -1, -1, 1})
}

func threeSum(nums []int) [][]int {

	returnVals := make(map[[3]int]bool, 0)
	numbers := make(map[int]int, 0)
	for _, v := range nums {
		numbers[v]++
	}
	nums = []int{}
	for k, v := range numbers {
		if v >= 3 {
			v = 3
		}
		for v > 0 {
			nums = append(nums, k)
			v--
		}
	}
	fmt.Println(nums)
	sort.Ints(nums)

	l := 0
	r := len(nums) - 1

	for l+1 < r {
		val := nums[r] + nums[l]

		if val < 0 {
			for k := r - 1; k > l; k-- {
				if val+nums[k] < 0 {
					l++
					break
				}
				if val+nums[k] == 0 {
					returnVals[[3]int{nums[l], nums[k], nums[r]}] = true
					l++
				}
			}
		} else {
			for k := l + 1; k < r; k++ {
				if val+nums[k] > 0 {
					r--
					break
				}
				if val+nums[k] == 0 {
					returnVals[[3]int{nums[l], nums[k], nums[r]}] = true
					l++
				}
			}
		}
	}
	rV := make([][]int, 0)
	for k, _ := range returnVals {
		rV = append(rV, []int{k[0], k[1], k[2]})
	}
	return rV
}

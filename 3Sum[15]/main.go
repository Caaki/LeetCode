package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(threeSum([]int{-1, -1, -1, 1}))
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
	sort.Ints(nums)

	r := len(nums) - 1
	for i := 0; i < r; i++ {
		for j := r; j >= i+1; j-- {
			needed := (nums[j] + nums[i]) * -1
			if v, ok := numbers[needed]; ok {
				c := 0
				if nums[j] == needed {
					c++
				}
				if nums[i] == needed {
					c++
				}
				if c >= v {
					break
				}
				minVal := nums[i]
				maxVal := nums[j]
				if needed > maxVal {
					needed, maxVal = maxVal, needed
				} else if needed < minVal {
					needed, minVal = minVal, needed
				}

				returnVals[[3]int{minVal, needed, maxVal}] = true
			} else {
				if v > 0 {
					r--
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

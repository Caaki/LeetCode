package main

import (
	"math"
)

func main() {}

func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}

	maxVal := nums[0]
	values := make(map[int]int, 0)

	l := 0
	r := 0
	for ; r < k; r++ {
		values[nums[r]]++
		if nums[r] > maxVal {
			maxVal = nums[r]
		}
	}

	returnArr := make([]int, 0)
	returnArr = append(returnArr, maxVal)

	for ; r < len(nums); r++ {
		values[nums[r]]++
		if nums[r] > maxVal {
			values = map[int]int{nums[r]: values[nums[r]]}
			maxVal = nums[r]
			returnArr = append(returnArr, maxVal)
			values[nums[l]]--
			if values[nums[l]] == 0 {
				delete(values, nums[l])
			}
			l++
		} else {
			values[nums[l]]--
			if values[nums[l]] == 0 {
				delete(values, nums[l])
				if maxVal == nums[l] {
					maxVal = findmaxVal(values)
				}
			}
			returnArr = append(returnArr, maxVal)
			l++
		}

	}

	return returnArr
}

func findmaxVal(m map[int]int) int {
	maxVal := math.MinInt
	for k, _ := range m {
		if k > maxVal {
			maxVal = k
		}
	}
	return maxVal
}

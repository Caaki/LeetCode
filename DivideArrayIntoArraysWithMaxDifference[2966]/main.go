package main

import (
	"slices"
)

func main() {
}

func divideArray(nums []int, k int) [][]int {

	slices.Sort(nums)

	returnVals := make([][]int, 0)

	temp := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if (i+1)%3 == 0 {
			temp = append(temp, nums[i])
			returnVals = append(returnVals, temp)
			if !CheckDifference(temp, k) {
				return [][]int{}
			}
			temp = make([]int, 0)
		} else {
			temp = append(temp, nums[i])
		}
	}

	return returnVals
}

func CheckDifference(n []int, k int) bool {
	a := n[0] - n[1]
	b := n[0] - n[2]
	c := n[1] - n[2]

	if a < 0 {
		a *= -1
	}

	if b < 0 {
		b *= -1
	}

	if c < 0 {
		c *= -1
	}

	if a > k || b > k || c > k {
		return false
	}
	return true
}

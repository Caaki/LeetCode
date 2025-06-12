package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print(maxAdjacentDistance([]int{1, 2, -1, 2, -10, -4, 24, 25, -10}))
}

func maxAdjacentDistance(nums []int) int {
	maximun := math.MinInt

	prev := nums[len(nums)-1]

	for i := 0; i < len(nums); i++ {
		temp := prev - nums[i]
		if temp < 0 {
			temp *= -1
		}
		if temp > maximun {
			maximun = temp
		}
		prev = nums[i]
	}

	return maximun
}

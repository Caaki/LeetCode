package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxAbsoluteSum([]int{-3, -5, -3, -2, -6, 3, 10, -10, -8, -3, 0, 10, 3, -5, 8, 7, -9, -9, 5, -8}))
}

func maxAbsoluteSum(nums []int) int {
	maxSubArray := float64(nums[0])
	maxVal := maxSubArray
  minVal:=maxVal
  for i := 1; i < len(nums); i++ {
		maxSubArray = CheckMax(maxSubArray, float64(nums[i]))
		if maxSubArray >= maxVal{
			maxVal = maxSubArray
		}
	}
	maxSubArray = float64(nums[0])
  for i := 1; i < len(nums); i++ {
		maxSubArray = CheckMin(maxSubArray, float64(nums[i]))
		if maxSubArray< minVal {
			minVal = maxSubArray
		}
	}

	return int(math.Max(maxVal,math.Abs(minVal)))
}

func CheckMax(maxVal, current float64) float64 {
	if current >= maxVal+current {
		return current
	}
	return current + maxVal
}

func CheckMin(maxVal, current float64) float64 {
	if current >= current+maxVal{
		return current+maxVal
	}
	return current
}

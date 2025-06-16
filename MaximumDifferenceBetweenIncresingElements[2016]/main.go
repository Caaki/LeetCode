package main

func main() {}

func maximumDifference(nums []int) int {
	current := nums[len(nums)-1]
	difference := 0
	for i := len(nums) - 2; i >= 0; i-- {
		if current < nums[i] {
			current = nums[i]
			continue
		}

		temp := current - nums[i]
		if temp > difference {
			difference = temp
		}
	}
	if difference == 0 {
		return -1
	}
	return difference
}

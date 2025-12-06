package main

import "fmt"

/*
Given a binary array nums, return the maximum number of consecutive 1's in the array.
*/

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1}))
}

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	temp := 0
	for _, v := range nums {
		if v == 1 {
			temp++
		} else {
			if count < temp {
				count = temp
			}
			temp = 0
		}
	}
	if count < temp {
		return temp
	}
	return count
}

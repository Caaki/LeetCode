package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 3, 4, 4, 5, 5, 5, 6}))
}

func removeDuplicates(nums []int) int {
	position := 1
	current := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if current != nums[i] {
			nums[position] = nums[i]
			position++
			current = nums[i]
			count++
		}
	}
	return count
}

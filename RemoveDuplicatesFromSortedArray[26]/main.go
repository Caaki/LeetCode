package main

func main() {}

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

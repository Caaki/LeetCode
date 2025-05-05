package main

func main() {}

func buildArray(nums []int) []int {
	answers := make([]int, len(nums))

	for i, _ := range nums {
		answers[i] = nums[nums[i]]
	}
	return answers
}

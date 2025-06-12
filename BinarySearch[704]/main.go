package main

func main() {}

func search(nums []int, target int) int {

	l := 0
	h := len(nums) - 1

	for l <= h {
		mid := l + (h-l)/2
		num := nums[mid]
		if target == num {
			return mid
		}
		if num < target {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}

	return -1
}

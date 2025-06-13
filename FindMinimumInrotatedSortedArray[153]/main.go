package main

func main() {}

func findMin(nums []int) int {

	l := 0
	r := len(nums) - 1

	minVal := nums[0]

	for l <= r {
		mid := l + (r-l)/2
		val := nums[mid]

		if nums[0] > val {
			r = mid - 1
			if minVal > val {
				minVal = val
			}
		} else {
			l = mid + 1
		}
	}

	return minVal

}

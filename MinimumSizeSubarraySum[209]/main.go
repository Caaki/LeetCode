package main

func main() {}

func minSubArrayLen(target int, nums []int) int {

	l := 0
	size := len(nums)
	sum := 0
	t := false
	for r := 0; r < len(nums); r++ {

		sum += nums[r]
		for sum >= target {
			if size > r-l {
				size = r - l + 1
				t = true
			}
			sum -= nums[l]
			l++
		}

	}

	if !t {
		return 0
	}

	return size
}

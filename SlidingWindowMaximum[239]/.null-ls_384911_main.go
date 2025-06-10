package main

func main() {}

func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}

	deque := make([]int, 0)

	returnVal := make([]int, 0)
	l := 0
	r := 0
	for r < len(nums) {
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[r] {
			deque = deque[:len(deque)-1]
		}
		deque[len(deque)-1] = nums[r]

		if l > deque[0] {
			deque = deque[1:]
		}

		if (r + 1) >= k {
			returnVal = append(returnVal, deque[0])
			l++
		}
		r++
	}

	return returnVal
}

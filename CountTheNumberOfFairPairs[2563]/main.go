package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(countFairPairs([]int{1, 7, 9, 2, 5}, 11, 11))
}

func countFairPairs(nums []int, lower int, upper int) int64 {

	sort.Ints(nums)
	result := int64(0)
	lastR := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		l := findLowest(nums, lower-nums[i], upper-nums[i], i+1, lastR)
		if l == -1 {
			continue
		}

		r := findHighest(nums, upper-nums[i], lower-nums[i], i+1, lastR)
		if r == -1 {
			break
		}
		if nums[lastR] > nums[r] {
			lastR = r
		}
		result += int64(r - l + 1)
	}

	return result
}
func findLowest(nums []int, lowest, h, i, j int) int {
	l, r := i, j
	res := -1

	for l <= r {
		mid := l + (r-l)/2
		if lowest > nums[mid] {
			l = mid + 1
		} else if h < nums[mid] {
			r = mid - 1
		} else {
			res = mid
			r = mid - 1
		}
	}
	return res

}

func findHighest(nums []int, upper, low, j, k int) int {
	l, r := j, k
	res := -1

	for l <= r {
		mid := l + (r-l)/2
		if upper < nums[mid] {
			r = mid - 1
		} else if low > nums[mid] {
			l = mid + 1
		} else {
			res = mid
			l = mid + 1
		}
	}
	return res
}

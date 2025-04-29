package main

func main() {}

func containsNearbyDuplicate(nums []int, k int) bool {

	vals := make(map[int]int16, 0)
	l := 0
	for r := 0; r < len(nums); r++ {
		for r-l > k {
			vals[nums[l]]--
			l++
		}
		vals[nums[r]]++
		if r-l <= k && vals[nums[r]] >= 2 {
			return true
		}
	}

	return false

}

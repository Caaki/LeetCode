package main

func main() {}

func countSubarrays(nums []int, k int) int64 {

	maxVal := 0

	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	l := 0
	count := 0
	result := 0

	for r := 0; r < len(nums); r++ {
		if nums[r] == maxVal {
			count++
		}

		for count >= k {
			result += len(nums) - r
			if nums[l] == maxVal {
				count--
			}
			l++
		}
	}
	return int64(result)
}

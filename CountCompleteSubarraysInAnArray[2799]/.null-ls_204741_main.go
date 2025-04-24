package main

func countCompleteSubarrays(nums []int) int {
	dis := make(map[int]bool)
	for _, v := range nums {
		dis[v] = true
	}
	distinct := len(dis)
	count := 0

	n := len(nums)

	for l := 0; l < n; l++ {
		tempMap := make(map[int]int)
		tempCount := 0
		for r := l; r < n; r++ {
			tempMap[nums[r]]++
			if tempMap[nums[r]] == 1 {
				tempCount++
			}
			if tempCount == distinct {
				count++
			}
		}
	}

	return count
}

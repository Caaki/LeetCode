package main

import (
	"slices"
)

func main() {
}

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	values := make(map[int]int, 0)
	order := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		if _, ok := values[nums1[i][0]]; !ok {
			order = append(order, nums1[i][0])
		}
		values[nums1[i][0]] += nums1[i][1]
	}
	for i := 0; i < len(nums2); i++ {
		if _, ok := values[nums2[i][0]]; !ok {
			order = append(order, nums2[i][0])
		}
		values[nums2[i][0]] += nums2[i][1]
	}

	slices.Sort(order)
	nums3 := make([][]int, len(order))
	for i, v := range order {
		nums3[i] = []int{order[i], values[v]}
	}
	return nums3
}

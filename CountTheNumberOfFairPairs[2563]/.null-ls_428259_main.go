package main

import (
	"fmt"
	"sort"
)

func main() {
	countFairPairs([]int{5, 2, 3, 1, 67, 1, 3}, 1, 2)
}

func countFairPairs(nums []int, lower int, upper int) int64 {

	sort.Ints(nums)
	fmt.Println(nums)
	return 1
}

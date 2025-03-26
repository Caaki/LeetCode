package main

import (
	"fmt"
	"slices"
)

func main() {
	values := [][]int{{596, 904, 960, 232, 120, 932, 176}, {372, 792, 288, 848, 960, 960, 764}, {652, 92, 904, 120, 680, 904, 120}, {372, 960, 92, 680, 876, 624, 904}, {176, 652, 64, 344, 316, 764, 316}, {820, 624, 848, 596, 960, 960, 372}, {708, 120, 456, 92, 484, 932, 540}}
	x := 28
	fmt.Println(minOperations(values, x))
}
func minOperations(grid [][]int, x int) int {
	values := make([]int, len(grid)*len(grid[0]))

	c := 0

	for _, v := range grid {
		for _, v2 := range v {
			values[c] = v2
			c++
		}
	}
	slices.Sort(values)
	median := values[len(values)/2]
	count := 0
	for i := 0; i < len(values); i++ {

		if values[i]%x != median%x {
			return -1
		}

		if i == len(values)/2 {
			count *= -1
		}
		count += (values[i] - median) / x
	}
	return count
}

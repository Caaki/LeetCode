package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countDays(57, [][]int{{3, 49}, {23, 44}, {21, 56}, {26, 55}, {23, 52}, {2, 9}, {1, 48}, {3, 31}}))
}

type byFirst [][]int

func (b byFirst) Len() int {
	return len(b)
}

func (b byFirst) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byFirst) Less(i, j int) bool {
	return b[i][0] < b[j][0]
}

func countDays(days int, meetings [][]int) int {
	sort.Sort(byFirst(meetings))
	newArray := [][]int{}
	for i := 0; i < len(meetings); i++ {
		minV := meetings[i][0]
		maxV := meetings[i][1]
		for j := i + 1; j < len(meetings); j++ {
			if meetings[j][0] <= maxV {
				if maxV < meetings[j][1] {
					maxV = meetings[j][1]
				}
				i = j
				continue
			}
			break
		}
		newArray = append(newArray, []int{minV, maxV})
	}
	sum := 0
	for _, v := range newArray {
		sum += (v[1] - v[0]) + 1
	}
	return days - sum
}

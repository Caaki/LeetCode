package main

import (
	"fmt"
)

type Values struct {
	Lower  []int
	Higher []int
}

type Values2 struct {
	Lowest  int
	Highest int
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 2, 6, 3, 1, 5, 4}))
}

func longestConsecutive(nums []int) int {
	values := make(map[int]*Values2)
	longest := 0
	for _, num := range nums {
		if _, ok := values[num]; ok {
			continue
		}
		values[num] = &Values2{num, num}
		getFirst(values, num, num)
		getLast(values, num, num)
		if longest < int(values[num].Highest-values[num].Lowest)+1 {
			longest = int(values[num].Highest-values[num].Lowest) + 1
		}
	}
	return longest
}

func getFirst(values map[int]*Values2, num int, appendTo int) {
	if _, ok := values[num-1]; ok {
		values[appendTo].Lowest = values[num-1].Lowest
		getFirst(values, values[appendTo].Lowest, appendTo)
	}
}

func getLast(values map[int]*Values2, num int, appendTo int) {

	if _, ok := values[num+1]; ok {
		values[appendTo].Highest = values[num+1].Highest
		getLast(values, values[appendTo].Highest, appendTo)
	}
}

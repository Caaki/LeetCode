package main

import "fmt"

func main() {
	fmt.Print(minimumIndex([]int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}))
}

func minimumIndex(nums []int) int {
	valueMap := make(map[int]int)

	for _, v := range nums {
		valueMap[v]++
	}

	value := -1
	occuered := 0
	for k, v := range valueMap {
		if v > occuered {
			value = k
			occuered = v
		}
	}
	ocurances := make([]int, len(nums))
	count := 0
	for i, v := range nums {
		if v == value {
			count++
		}
		ocurances[i] = count
	}

	for i := 0; i < len(ocurances); i++ {
		lLen := len(nums) - (len(nums) - i) + 1
		rLen := len(nums) - (i)
		if (ocurances[i]*2) > lLen && (occuered-ocurances[i])*2 >= rLen {
			return i
		}
	}

	return -1

}

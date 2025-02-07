package main

import "fmt"

func main() {
	fmt.Println(tupleSameProduct([]int{2, 3, 4, 6}))
}

func tupleSameProduct(nums []int) int {
	values := make(map[int][][2]int)

	for i, v := range nums {
		if i+1 == len(nums) {
			break
		}
		for j := i + 1; j < len(nums); j++ {
			product := v * nums[j]
			values[product] = append(values[product], [2]int{i, j})
		}
	}

	sum := 0
	for _, v := range values {
		if len(v) > 1 {
			sum += len(v) * 2 * ((len(v) - 1) * 2)
		}

	}

	return sum
}

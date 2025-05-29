package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{5, 3, 1, 5, 6, 87, 1, 2, 4, 56, 7, 1, 1, 643, 3, 21, 34, 1, 25, 623}))
}

func maxProfit(prices []int) int {
	maximum := 0

	l := 0
	r := 1
	n := len(prices)
	for r < n {
		if prices[r] <= prices[l] {
			l = r
			r = r + 1
			continue
		}
		temp := prices[r] - prices[l]
		if temp > maximum {
			maximum = temp
		}
		r++
	}
	return maximum
}

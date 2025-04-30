package main

import "strconv"

func main() {}

func findNumbers(nums []int) int {
	count := 0

	for _, v := range nums {
		temp := strconv.Itoa(v)
		if len(temp)%2 == 0 {
			count++
		}
	}

	return count
}

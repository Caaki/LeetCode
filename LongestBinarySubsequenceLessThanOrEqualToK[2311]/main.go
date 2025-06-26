package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(longestSubsequence("1001010", 5))
}

func longestSubsequence(s string, k int) int {

	count := 0
	sum := float64(0)
	k2 := float64(k)

	n := len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 49 {
			temp := math.Pow(2, float64(n-i))
			if sum+temp > k2 {
				continue
			} else {
				sum += temp
				count++
			}
		} else {
			count++
		}
	}

	return count
}

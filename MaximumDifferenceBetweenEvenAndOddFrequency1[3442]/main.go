package main

import "math"

func maxDifference(s string) int {

	values := make([]int, 26)

	for i, _ := range s {
		values[s[i]-'a']++
	}

	maxOdd := math.MinInt
	minEven := math.MaxInt

	for _, v := range values {
		if v == 0 {
			continue
		}

		if v%2 == 0 {
			if minEven > v {
				minEven = v
			}
		} else {
			if maxOdd < v {
				maxOdd = v
			}
		}
	}

	return maxOdd - minEven
}

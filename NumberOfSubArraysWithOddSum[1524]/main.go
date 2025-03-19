package main

import (
	"fmt"
)

func main() {
	fmt.Println(numOfSubarrays([]int{1, 3, 5}))
}

func numOfSubarrays(arr []int) int {
	sum := 0
	odd := false
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			odd = false
		} else {
			odd = true
			sum++
		}
		for j := i + 1; j < len(arr); j++ {
			odd2 := arr[j]%2 == 1
			if odd == odd2 {
				odd = false
			} else {
				odd = true
				sum++
			}
		}

	}
	return sum % 1000000007
}

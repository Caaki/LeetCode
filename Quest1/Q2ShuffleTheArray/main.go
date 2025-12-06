package main

import "fmt"

/*
Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].
Return the array in the form [x1,y1,x2,y2,...,xn,yn]
*/
func main() {
	nusm := []int{1, 2, 3, 4}
	n := 2
	fmt.Println(shuffle(nusm, n))
}

func shuffle(nums []int, n int) []int {
	answer := make([]int, n*2)
	for q, i, j := 0, 0, n; i < n; {
		answer[q] = nums[i]
		q++
		answer[q] = nums[j]
		i++
		j++
		q++
	}
	return answer
}

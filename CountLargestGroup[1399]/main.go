package main

import "fmt"

func main() {
	fmt.Println(countLargestGroup(1000))
}

func countLargestGroup(n int) int {

	values := make(map[int]int, 0)
	for i := 1; i <= n; i++ {
		values[digSum(i)]++
	}

	maxCount := 0
	for _, v := range values {
		if v > maxCount {
			maxCount = v
		}
	}

	result := 0
	for _, v := range values {
		if v == maxCount {
			result++
		}
	}

	return result

}

func digSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

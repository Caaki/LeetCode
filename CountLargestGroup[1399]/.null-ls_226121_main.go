package main

import "fmt"

func main() {
	fmt.Println(digSum(5))
	fmt.Println(digSum(25))
	fmt.Println(digSum(115))
}

func countLargestGroup(n int) int {

	values := make(map[int]int, 0)
	for i := 1; i <= n; i++ {
		values[digSum(i)] = i
	}

	return 0

}

func digSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

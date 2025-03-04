package main

import "fmt"

func main() {

	fmt.Println("Checking for number 45:", checkPowersOfThree(45))
	fmt.Println("Checking for number 12:", checkPowersOfThree(12))
}

func checkPowersOfThree(n int) bool {
	for n > 0 {
		if n%3 == 2 {
			return false
		}
		n /= 3
	}
	return true
}

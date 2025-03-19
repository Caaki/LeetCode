package main

import "fmt"

func main() {

	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}

func containsDuplicate(nums []int) bool {

	m := make(map[int]bool, len(nums))

	for _, n := range nums {
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = true
	}

	return true

}

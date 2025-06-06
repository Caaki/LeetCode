package main

import "fmt"

func main() {

	fmt.Print(getLongestSubsequence([]string{"f", "o", "l", "i", "v", "r", "a", "q", "z", "g", "p", "h", "n", "c", "y", "u", "e", "b", "k", "t", "w", "m", "x", "j", "s", "d"}, []int{0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0}))
}

func getLongestSubsequence(words []string, groups []int) []string {

	ans1 := make([]string, 0)
	ans2 := make([]string, 0)

	cur0 := int8(0)
	cur1 := int8(1)

	for i := 0; i < len(words); i++ {
		if cur0 != int8(groups[i]) {
			ans1 = append(ans1, words[i])
			cur0 = int8(groups[i])
		}

		if cur1 != int8(groups[i]) {
			ans2 = append(ans2, words[i])
			cur1 = int8(groups[i])
		}
	}

	if len(ans1) > len(ans2) {
		return ans1
	}
	return ans2
}

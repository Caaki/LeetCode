package main

import "fmt"

var letters []string = []string{"a", "b", "c"}

func main() {
	fmt.Print(getHappyString(1, 3))
}

func getHappyString(n int, k int) string {
	results := make([]string, 0)
	getLetterCombination(&results, "", n, k)
	if len(results) < k {
		return ""
	} else {
		return results[k-1]
	}
}

func getLetterCombination(results *[]string, current string, n int, k int) {
	if len(current) == n {
		*results = append(*results, current)
		return
	}
	for _, c := range letters {
		if len(current) <= 0 {
			getLetterCombination(results, c, n, k)
			continue
		}
		if len(current) < n && current[len(current)-1] != c[0] {
			getLetterCombination(results, current+c, n, k)
			continue
		}
	}
}

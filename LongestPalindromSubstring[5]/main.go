package main

import "fmt"

func main() {

	fmt.Println(longestPalindrome("bb"))

}

func longestPalindrome(s string) string {
	maxPalindrom := ""
	for i := 0; i < len(s); i++ {
		countPair := 1
		countNotEqual := 1
		var current string
		maxPair := "" + string(s[i])
		maxNotEqual := "" + string(s[i])
		for min, max := i-1, i+1; min >= 0 && max < len(s); {
			if s[min] == s[max] {
				countNotEqual += 2
				maxNotEqual = string(s[min]) + maxNotEqual + string(s[max])
				max++
				min--
				continue
			}
			break
		}

		for min, max := i, i+1; min >= 0 && max < len(s); {
			if countPair == 1 {
				if s[i] == s[max] {
					countPair++
					maxPair = maxPair + string(s[max])
					max++
					min--
					continue
				}
			}

			if s[min] == s[max] {
				countPair += 2
				maxPair = string(s[min]) + maxPair + string(s[max])
				max++
				min--
				continue
			}
			break
		}

		if countPair > countNotEqual {
			current = maxPair
		} else {
			current = maxNotEqual
		}
		if len(current) >= len(maxPalindrom) {
			maxPalindrom = current
		}
	}
	return maxPalindrom
}

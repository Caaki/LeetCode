package main

import "fmt"

// a-z = 97 - 122
// A-Z = 65 - 90
// 0-9 = 48 - 57
func main() {
	isPalindrome("aAzZ")
	fmt.Print(int('a'), int('A')+32, int(97-65))
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		v1 := int(s[i])
		for !(v1 >= 65 && v1 <= 90) && !(v1 <= 122 && v1 >= 97) && !(v1 >= 48 && v1 <= 57) {
			i++
			if i >= j {
				break
			}
			v1 = int(s[i])
		}
		if v1 < 97 {
			v1 += 32
		}

		v2 := int(s[j])
		for !(v2 >= 65 && v2 <= 90) && !(v2 <= 122 && v2 >= 97) && !(v2 >= 48 && v2 <= 57) {
			j--
			if i >= j {
				break
			}
			v2 = int(s[j])
		}
		if v2 < 97 {
			v2 += 32
		}
		if i >= j {
			return true
		}
		if v1 != v2 {
			return false
		}
		i++
		j--
	}
	return true
}

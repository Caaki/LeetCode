package main

func main() {}

func characterReplacement(s string, k int) int {
	caracters := make(map[byte]int, 0)
	max := 1

	l := 0
	r := 0
	n := len(s)
	tempCount := 0

	for r < n && l < n {
		caracters[s[r]]++
		for _, v := range caracters {
			if v > tempCount {
				tempCount = v
			}
		}

		for (r-l+1)-tempCount > k {
			caracters[s[l]]--
			l++
		}

		if r-l+1 > max {
			max = r - l + 1
		}
		r++
	}

	return max

}

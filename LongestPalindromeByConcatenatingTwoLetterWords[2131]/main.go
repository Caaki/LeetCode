package main

func main() {}

func longestPalindrome(words []string) int {
	combination := make(map[string]int, 0)
	same := make(map[string]int, 0)
	extra := false
	result := 0

	for _, v := range words {
		if v[0] == v[1] {
			same[v]++
		} else {
			combination[v]++
		}
	}

	for _, v := range same {
		if v%2 == 0 {
			result += 2 * v
			continue
		}
		if v == 1 {
			extra = true
			continue
		} else {
			result += (v - 1) * 2
			extra = true
			continue
		}
	}

	for k, v := range combination {
		k2 := string(k[1]) + string(k[0])
		if v1, ok := combination[k2]; ok {
			if v1 < v {
				result += 4 * v1
			} else {
				result += 4 * v
			}
			combination[k2] = 0
		}
	}

	if extra {
		return result + 2
	}
	return result
}

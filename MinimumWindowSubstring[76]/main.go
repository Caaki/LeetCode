package main

func minWindow(s string, t string) string {

	n := len(s)
	if len(t) > n {
		return ""
	}
	if t == s {
		return t
	}

	caracters := make(map[byte]int, 0)
	sum := 0
	for i, _ := range t {
		caracters[t[i]]++
		sum++
	}

	needed := make(map[byte]int, 0)
	for k, _ := range caracters {
		needed[k] = 0
	}

	l := 0
	r := 0
	count := 0
	returnVal := s + t
	for r < n || count == sum {
		if sum == count {
			if r-l < len(returnVal) {
				returnVal = s[l:r]
			}
			if v, ok := needed[s[l]]; ok {
				if v <= caracters[s[l]] {
					count--
				}
				needed[s[l]]--
			}
			l++
			continue
		}
		if v, ok := needed[s[r]]; ok {
			needed[s[r]]++
			if v < caracters[s[r]] {
				count++
			}
		}
		r++
	}

	if len(returnVal) == len(t)+len(s) {
		return ""
	}
	return returnVal
}

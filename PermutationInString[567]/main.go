package main

func checkInclusion(s1, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	contants := make(map[byte]int, 0)
	sum := 0
	for i, _ := range s1 {
		contants[s1[i]]++
		sum++
	}

	needed := make(map[byte]int, 0)
	for k, _ := range contants {
		needed[k] = 0
	}

	count := 0

	l := 0
	r := len(s1) - 1

	for i := 0; i < len(s1); i++ {
		if _, ok := needed[s2[i]]; ok {
			needed[s2[i]]++
			if needed[s2[i]] <= contants[s2[i]] {
				count++
			}
		}
	}

	if sum == count {
		return true
	}
	for r < len(s2)-1 {

		r++
		if _, ok := needed[s2[r]]; ok {
			needed[s2[r]]++
			if needed[s2[r]] <= contants[s2[r]] {
				count++
			}
		}
		if _, ok := needed[s2[l]]; ok {
			needed[s2[l]]--
			if needed[s2[l]] < contants[s2[l]] {
				count--
			}
		}
		l++

		if sum == count {
			return true
		}
	}
	return false
}

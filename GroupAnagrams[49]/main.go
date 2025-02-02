package main

func main() {

}

func groupAnagrams(strs []string) [][]string {
	m := [][]string{{strs[0]}}
	test := false
	for i := 1; i < len(strs); i++ {
		test = false
		for j := 0; j < len(m); j++ {
			if isAnagram(m[j][0], strs[i]) {
				m[j] = append(m[j], strs[i])
				test = true
				break
			}
		}
		if !test {
			m = append(m, []string{strs[i]})
		}
	}

	return m
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := map[rune]int16{}

	for _, c := range s {
		m[c]++
	}
	for _, c := range t {
		m[c]--
		if m[c] < 0 {
			return false
		}
	}
	return true
}

func groupAnagrams2(strs []string) [][]string {
	m := make(map[[26]int8][]string)

	for _, s := range strs {
		var cnt [26]int8
		for _, c := range s {
			cnt[c-'a']++
		}
		m[cnt] = append(m[cnt], s)
	}

	r := make([][]string, 0, len(m))

	for _, v := range m {
		r = append(r, v)
	}
	return r
}

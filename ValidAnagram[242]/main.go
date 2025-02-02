package main

func main() {

}

func isAnagram(s string, t string) bool {
	var m1 map[rune]int16
	var m2 map[rune]int16
	createMapFromString(&s, m1)
	createMapFromString(&t, m2)
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if v2, ok := m2[k]; !ok || v2 != v {
			return false
		}
	}
	return true

}
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := map[byte]int16{}

	for _, c := range []byte(s) {
		m1[c]++
	}

	for _, c := range []byte(t) {
		m1[c]--
		if m1[c] < 0 {
			return false
		}
	}

	return true

}

func createMapFromString(s *string, m map[rune]int16) {
	for _, c := range *s {
		if i, ok := m[c]; ok {
			m[c] = i + 1
		} else {
			m[c] = 1
		}
	}
}

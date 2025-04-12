package main

func main() {}

func removeDuplicateLetters(s string) string {
	letters := make([][]int, 26)
	count := 0

	for i, v := range s {
		if len(letters[v-'a']) == 0 {
			count++
		}
		letters[v-'a'] = append(letters[v-'a'], i)
	}

	val := ""
	start := 0
	for count > 0 {
		for i := 0; i < len(letters); i++ {
			if len(letters[i]) == 0 {
				continue
			}
			first := -1
			for _, idx := range letters[i] {
				if idx >= start {
					first = idx
					break
				}
			}
			if first == -1 {
				continue
			}

			ok := true
			for j := 0; j < len(letters); j++ {
				if len(letters[j]) == 0 || j == i {
					continue
				}
				found := false
				for _, idx := range letters[j] {
					if idx > first {
						found = true
						break
					}
				}
				if !found {
					ok = false
					break
				}
			}

			if ok {
				val += string('a' + i)
				letters[i] = []int{}
				start = first + 1
				count--
				break
			}
		}
	}

	return val
}

package main

func main() {}

func findWordsContaining(words []string, x byte) []int {

	result := make([]int, 0)

	for i, v := range words {
		for j := 0; j < len(v); j++ {
			if v[j] == x {
				result = append(result, i)
				break
			}
		}
	}

	return result
}

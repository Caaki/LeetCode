package main

func main() {}

func numEquivDominoPairs(dominoes [][]int) int {

	values := make(map[[2]int8]int, 0)

	for _, v := range dominoes {
		if v[0] > v[1] {
			values[[2]int8{int8(v[0]), int8(v[1])}]++
		} else {
			values[[2]int8{int8(v[1]), int8(v[0])}]++
		}
	}
	result := 0
	for _, v := range values {
		if v > 1 {
			result += (v * (v - 1)) / 2
		}
	}

	return result
}

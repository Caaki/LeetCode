package main

func main() {}

func minDominoRotations(tops []int, bottoms []int) int {

	values := make(map[int8]int, 0)

	for i := 0; i < len(tops); i++ {
		if tops[i] == bottoms[i] {
			values[int8(tops[i])]++
		} else {
			values[int8(tops[i])]++
			values[int8(bottoms[i])]++
		}
	}

	none := true
	possible := make([]int, 7)
	needed := len(tops)
	for k, v := range values {
		if v >= needed {
			possible[k] = v
			none = false
		}
	}
	if none {
		return -1
	}

	maxTops := make(map[int8]int, 0)
	maxBottoms := make(map[int8]int, 0)
	for i := 0; i < len(tops); i++ {
		if possible[tops[i]] > 0 {
			maxTops[int8(tops[i])]++
		}
		if possible[bottoms[i]] > 0 {
			maxBottoms[int8(bottoms[i])]++
		}
	}

	maxVal := 0
	for _, v := range maxTops {
		if v > maxVal {
			maxVal = v
		}
	}
	for _, v := range maxBottoms {
		if v > maxVal {
			maxVal = v
		}
	}
	return len(tops) - maxVal
}

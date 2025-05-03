package main

func main() {}

func minDominoRotations(tops []int, bottoms []int) int {

	values := make([]int, 7)

	maxTops := make([]int, 7)
	maxBottoms := make([]int, 7)
	for i := 0; i < len(tops); i++ {
		if tops[i] == bottoms[i] {
			values[tops[i]]++
		} else {
			values[tops[i]]++
			values[bottoms[i]]++
		}
		maxTops[tops[i]]++
		maxBottoms[bottoms[i]]++
	}
	none := true
	for _, v := range values {
		if v >= len(tops) {
			none = false
			break
		}
	}
	if none {
		return -1
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

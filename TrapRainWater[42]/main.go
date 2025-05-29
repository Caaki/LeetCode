package main

func main() {}

func trap(height []int) int {

	if len(height) < 2 {
		return 0
	}
	n := len(height)
	maxL := 0
	maxR := 0

	l := 0
	r := n - 1

	result := 0
	for l <= r {
		if maxL < maxR {
			current := height[l]

			temp := maxL - current
			if temp > 0 {
				result += temp
			}
			if current > maxL {
				maxL = current
			}
			l++
		} else {
			current := height[r]
			temp := maxR - current
			if temp > 0 {
				result += temp
			}
			if current > maxR {
				maxR = current
			}
			r--
		}
	}
	return result
}

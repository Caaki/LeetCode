package main

func main() {}

func maxArea(height []int) int {

	maxVolume := 0
	l := 0
	r := len(height)

	for l <= r {
		if height[l] <= 0 {
			l++
		}
		if height[r] <= 0 {
			r--
		}
		if height[l] > 0 && height[r] > 0 {
			break
		}
	}

	if height[l] == 0 || height[r] == 0 {
		return 0
	}

	for l <= r {
		temp := (r - l) * getMin(height[l], height[r])
		if temp > maxVolume {
			maxVolume = temp
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return maxVolume
}

func getMin(i, j int) int {
	if i < j {
		return i
	}
	return j
}

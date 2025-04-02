package main

func main() {}

func maximumTripletValue(nums []int) int64 {
	maxVal := 0
	maxDifference := 0
	maxTriplet := 0

	for i, v := range nums {
		if maxTriplet < maxDifference*v && len(nums)-i >= 1 {
			maxTriplet = maxDifference * v
		}
		if maxDifference < maxVal-v && len(nums)-i >= 2 {
			maxDifference = (maxVal - v)
			continue
		}
		if maxVal < v && len(nums)-i >= 3 {
			maxVal = v
			continue
		}
	}

	if maxTriplet <= 0 {
		return 0
	}

	return int64(maxTriplet)
}

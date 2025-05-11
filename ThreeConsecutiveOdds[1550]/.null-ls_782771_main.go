package main

func main() {}

func threeConsecutiveOdds(arr []int) bool {
	count := int8(0)

	for _, v := range arr {
		if v%2 == 0 {
			count = 0
		} else {
			count++
			if count == 3 {
				return true
			}
		}
	}
	return false
}

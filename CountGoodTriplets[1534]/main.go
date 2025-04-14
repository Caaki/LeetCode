package main

func main() {}

func countGoodTriplets(arr []int, a, b, c int) int {
	result := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {

				if getVal(arr[i], arr[j]) <= a {
					if getVal(arr[j], arr[k]) <= b {
						if getVal(arr[i], arr[k]) <= c {
							result++
						}
					}
				}
			}
		}
	}

	return result

}

func getVal(a, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}

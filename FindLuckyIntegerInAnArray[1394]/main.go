package main

func findLucky(arr []int) int {
	rVal := -1
	values := make(map[int]int, 0)

	for _, v := range arr {
		values[v]++
	}

	for k, v := range values {
		if k == v && rVal < v {
			rVal = v
		}
	}

	return rVal
}

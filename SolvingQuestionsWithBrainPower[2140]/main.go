package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mostPoints([][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}))
}

func mostPoints(questions [][]int) int64 {
	cache := make([]int, len(questions))
	return backTrack(0, &questions, &cache)
}

func backTrack(i int, questons *[][]int, cache *[]int) int64 {
	if i >= len(*questons) {
		return 0
	}
	if (*cache)[i] != 0 {
		return int64((*cache)[i])
	}

	p, bp := (*questons)[i][0], (*questons)[i][1]
	(*cache)[i] = int(math.Max(float64(backTrack(i+1, questons, cache)), float64(p+int(backTrack(i+1+bp, questons, cache)))))

	return int64((*cache)[i])
}

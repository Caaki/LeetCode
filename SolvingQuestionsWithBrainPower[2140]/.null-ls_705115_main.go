package main

import "fmt"

func main() {
	questons := make([][]int, 0)
	questons = append(questons, []int{2})
	backTrack(1, questons)
	fmt.Println(questons)
}

func backTrack(i int, questons [][]int) {
	questons[0][0] = i
}

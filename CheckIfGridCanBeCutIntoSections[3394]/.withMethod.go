package main

import "sort"

func main() {}

type byX [][]int

func (b byX) Len() int {
	return len(b)
}
func (b byX) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byX) Less(i, j int) bool {
	return b[i][0] < b[j][0]
}

type byY [][]int

func (b byY) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byY) Len() int {
	return len(b)
}

func (b byY) Less(i, j int) bool {
	return b[i][1] < b[j][1]
}

func doCheck(rectangles [][]int, x int, y int) int {
	lines := 0
	minX := 0
	finAll := false
	for i := 0; i < len(rectangles) && !finAll; i++ {
		minX = rectangles[i][y]
		for j := i + 1; j < len(rectangles); j++ {
			if minX <= rectangles[j][x] {
				lines++
				i = j - 1
				break
			}
			if minX < rectangles[j][y] {
				minX = rectangles[j][y]
			}
			if j == len(rectangles)-1 {
				finAll = true
			}
		}
	}
	return lines
}

func checkValidCuts(n int, rectangles [][]int) bool {
	sort.Sort(byX(rectangles))
	lines := doCheck(rectangles, 0, 2)

	if lines >= 2 {
		return true
	}

	sort.Sort(byY(rectangles))
	lines = doCheck(rectangles, 1, 3)

	return lines >= 2
}

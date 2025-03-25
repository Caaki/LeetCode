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

func checkValidCuts(n int, rectangles [][]int) bool {
	sort.Sort(byX(rectangles))
	lines := 0
	minX := 0
	finAll := false
	for i := 0; i < len(rectangles) && !finAll; i++ {
		minX = rectangles[i][2]
		for j := i + 1; j < len(rectangles); j++ {
			if minX <= rectangles[j][0] {
				lines++
				i = j - 1
				break
			}
			if minX < rectangles[j][2] {
				minX = rectangles[j][2]
			}
			if j == len(rectangles)-1 {
				finAll = true
			}
		}
	}

	if lines >= 2 {
		return true
	}

	lines = 0
	sort.Sort(byY(rectangles))
	finAll = false
	minY := 0
	for i := 0; i < len(rectangles) && !finAll; i++ {
		minY = rectangles[i][3]
		for j := i + 1; j < len(rectangles); j++ {
			if minY <= rectangles[j][1] {
				lines++
				i = j - 1
				break
			}
			if minY < rectangles[j][3] {
				minY = rectangles[j][3]
			}
			if j == len(rectangles)-1 {
				finAll = true
			}
		}
	}

	if lines >= 2 {
		return true
	}
	return false
}

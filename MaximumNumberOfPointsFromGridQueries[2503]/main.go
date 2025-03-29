package main

import (
	"container/heap"
	"slices"
)

func main() {}

type cell struct {
	i, j int
	val  int
}
type minHeap []cell

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(cell))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func maxPoints(grid [][]int, queries []int) []int {
	qSorted := append([]int{}, queries...)
	slices.Sort(qSorted)

	order := make(map[int]int, 0)
	for i, v := range queries {
		order[v] = i
	}

	answers := make([]int, len(queries))
	var toVisit = &minHeap{}
	heap.Init(toVisit)
	heap.Push(toVisit, cell{i: 0, j: 0, val: grid[0][0]})
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	count := 0
	for _, q := range qSorted {
		for toVisit.Len() > 0 && grid[(*toVisit)[0].i][(*toVisit)[0].j] < q {

			cordinates := heap.Pop(toVisit).(cell)

			count++
			for _, dir := range directions {
				x, y := cordinates.i+dir[0], cordinates.j+dir[1]
				if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && !visited[x][y] {
					visited[x][y] = true
					heap.Push(toVisit, cell{x, y, grid[x][y]})
				}
			}
		}
		if count > 1 {
			answers[order[q]] = count - 1

		} else {
			answers[order[q]] = count

		}
	}
	return answers
}

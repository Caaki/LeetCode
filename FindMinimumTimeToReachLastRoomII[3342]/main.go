package main

import (
	"container/heap"
	"math"
)

func main() {}

func minTimeToReach(moveTime [][]int) int {
	n := len(moveTime)
	m := len(moveTime[0])
	dist := make([][]int, n)

	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)

		for j := 0; j < m; j++ {
			dist[i][j] = math.MaxInt
		}
	}

	minHeap := &IntHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, [3]int{0, 0, 0})

	for minHeap.Len() > 0 {
		elements := heap.Pop(minHeap).([3]int)
		w, i, j := elements[0], elements[1], elements[2]

		for _, v := range [][3]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			ni, nj := i+v[0], j+v[1]

			if ni < 0 || nj < 0 || ni >= n || nj >= m {
				continue
			}
			cur := Max(w, moveTime[ni][nj]) + ((i+j)%2 + 1)

			if cur < dist[ni][nj] {
				dist[ni][nj] = cur
				heap.Push(minHeap, [3]int{cur, ni, nj})
			}
		}
	}

	return dist[n-1][m-1]

}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

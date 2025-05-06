package main

import (
	"container/heap"
	"math"
)

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

	dist[0][0] = 0

	minHeap := &IntHeap{}

	heap.Init(minHeap)
	heap.Push(minHeap, [3]int{0, 0, 0})

	for minHeap.Len() > 0 {
		element := minHeap.Pop().([3]int)
		w, i, j := element[0], element[1], element[2]
		for _, v := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			ni := i + v[0]
			nj := j + v[1]

			if ni < 0 || nj < 0 || ni >= n || nj >= m {
				continue
			}

			current := MaxInt(w, moveTime[ni][nj]) + 1
			if current < dist[ni][nj] {
				dist[ni][nj] = current
				heap.Push(minHeap, [3]int{current, ni, nj})
			}
		}
	}

	return dist[n-1][m-1]

}

func MaxInt(n, j int) int {
	if n > j {
		return n
	}
	return j
}

package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
}

func minOperations(nums []int, k int) int {
	h := &IntHeap{}
	arr2 := make([]int, len(nums))
	for i, v := range nums {
		arr2[i] = int(v)
	}
	*h = append(*h, arr2...)
	heap.Init(h)
	count := 0
	minVal := 0
	for minVal < k {
		minVal = add(heap.Pop(h).(int), heap.Pop(h).(int))
		count++
		heap.Push(h, int(minVal))
		fmt.Print(h)
	}

	return count
}

func add(x int, y int) int {

	return int(math.Min(float64(x), float64(y))*2 + math.Max(float64(x), float64(y)))

}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

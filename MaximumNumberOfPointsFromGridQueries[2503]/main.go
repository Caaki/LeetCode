package main

import "slices"

func main() {}

type Stack [][2]int

func (s Stack) Add(a [2]int) {
	s = append(s, a)
}

func (s Stack) Pop() [2]int {
	if len(s) < 1 {
		return [2]int{}
	}
	val := s[len(s)-1]
	s = s[:len(s)-1]
	return [2]int{val[0], val[1]}
}

func (s Stack) Peek() [2]int {
	if len(s) < 1 {
		return [2]int{}
	}
	val := s[len(s)-1]
	return [2]int{val[0], val[1]}
}

func maxPoints(grid [][]int, queries []int) []int {
	qSorted := append([]int{}, queries...)
	slices.Sort(qSorted)

	order := make(map[int]int, 0)
	for i, v := range queries {
		order[v] = i
	}

	answers := make([]int, len(queries))
	var toVisit = Stack{}
	toVisit.Add([2]int{0, 0})
	visited := make(map[[2]int]bool, 0)

	count := 0
	for _, q := range qSorted {
		for len(toVisit) > 0 && grid[toVisit[0][0]][toVisit[0][1]] < q {
			if _, ok := visited[toVisit.Peek()]; ok {
				toVisit.Pop()
				continue
			}
			count++
			cordinates := toVisit.Pop()
			visited[cordinates] = true
			if cordinates[0] < len(grid) && !visited[[2]int{cordinates[0] + 1, cordinates[1]}] {
				toVisit.Add([2]int{cordinates[0] + 1, cordinates[1]})
			}
			if cordinates[0] < len(grid[0]) && !visited[[2]int{cordinates[0], cordinates[1] + 1}] {
				toVisit.Add([2]int{cordinates[0], cordinates[1] + 1})
			}
		}
		answers[order[q]] = count
	}
	return answers
}

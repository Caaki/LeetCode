package main

import (
	"fmt"
	"math"
)

func main() {}

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	if len(edges) == 1 {
		return amount[0]
	}
	alieceMax := math.MinInt32
	fmt.Println(alieceMax)
	a := transformToTree(edges)
	allPaths := removeExtraPaths(findEdgeNumbers(a, bob))
	for _, v := range allPaths {
		am := append([]int{}, amount...)
		current := calculateMax(v, bob, am)
		fmt.Println("CUrrent is", current, v)
		if current > alieceMax {
			alieceMax = current
		}
	}

	return alieceMax
}

type Paths struct {
	path   []int
	isBob  bool
	bobPos int
}

func findEdgeNumbers(all [][]int, bob int) map[int]Paths {
	paths := map[int]Paths{}

	currentRoot := all[0][0]
	paths[all[0][0]] = Paths{path: []int{currentRoot}, isBob: false}
	for i := 0; i < len(all); i++ {
		parent, child := all[i][0], all[i][1]
		if child == bob {
			paths[child] = Paths{path: append(paths[parent].path, child), isBob: true, bobPos: len(paths[parent].path)}
		} else {
			paths[child] = Paths{path: append(paths[parent].path, child), isBob: paths[parent].isBob, bobPos: paths[parent].bobPos}
		}
		if currentRoot != parent {
			delete(paths, currentRoot)
			currentRoot = parent
		}
	}
	return paths
}

func calculateMax(paths Paths, bob int, amount []int) int {
	current := 0
	if !paths.isBob {
		for _, v := range paths.path {
			current += amount[v]
		}
		return current
	} else {
		bobPosition := paths.bobPos
		for i := 0; i < len(paths.path); i++ {
			if bobPosition > 0 {
				if bobPosition == i {
					amount[paths.path[i]] = amount[paths.path[i]] / 2
				}
				amount[paths.path[bobPosition]] = 0
				current += amount[paths.path[i]]

				bobPosition--
			} else {
				current += amount[paths.path[i]]
			}
		}
		return current
	}
}

func transformToTree(edges [][]int) [][]int {
	adjList := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}

	var tree [][]int
	visited := make(map[int]bool)
	queue := []int{0}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		visited[node] = true

		for _, neighbor := range adjList[node] {
			if !visited[neighbor] {
				tree = append(tree, []int{node, neighbor})
				queue = append(queue, neighbor)
			}
		}
	}

	return tree
}

func removeExtraPaths(paths map[int]Paths) map[int]Paths {
	filteredPaths := make(map[int]Paths)

	for key, p := range paths {
		isExtra := false
		for _, other := range paths {
			if key != other.path[len(other.path)-1] && isSubPath(p.path, other.path) {
				isExtra = true
				break
			}
		}
		if !isExtra {
			filteredPaths[key] = p
		}
	}

	return filteredPaths
}

func isSubPath(small, large []int) bool {
	if len(small) >= len(large) {
		return false
	}
	for i := 0; i <= len(large)-len(small); i++ {
		if equalSlices(small, large[i:i+len(small)]) {
			return true
		}
	}
	return false
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

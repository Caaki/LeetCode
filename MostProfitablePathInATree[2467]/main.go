package main

func main() {}

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	if len(edges) == 1 {
		return amount[0]
	}
	alieceMax := amount[0]
	allPaths := findEdgeNumbers(edges, bob)
	for _, v := range allPaths {
		current := calculateMax(v, bob, amount)
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
			paths[child] = Paths{path: append(paths[parent].path, child), isBob: paths[parent].isBob}
		}
		if currentRoot != parent {
			delete(paths, currentRoot)
			currentRoot = child
		}
	}
	return paths
}

func calculateMax(paths Paths, bob int, amount []int) int {
	maxAliece := amount[0]
	current := amount[0]
	if !paths.isBob {
		for _, v := range paths.path {
			current += amount[v]
			if current > maxAliece {
				maxAliece = current
			}
		}
		return maxAliece
	} else {
		bobPosition := paths.bobPos
		for i := 0; i < len(paths.path); i++ {
			if bobPosition >= 0 {
				if bobPosition == i {
					amount[i] = amount[i] / 2
				}
				amount[bob] = 0
				current += amount[paths.path[i]]
				if current > maxAliece {
					maxAliece = current
				}
			} else {
				current += amount[paths.path[i]]
				if current > maxAliece {
					maxAliece = current
				}
			}
		}
		return maxAliece
	}
}

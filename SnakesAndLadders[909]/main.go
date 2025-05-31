package main

import (
	"errors"
)

func main() {
}

var n int

func snakesAndLadders(board [][]int) int {

	n = len(board)

	visited := make(map[int]bool, 0)
	queue := [][]int{{1, 0}}
	visited[1] = true

	for len(queue) > 0 {
		val, err := PopFirst(&queue)
		if err != nil {
			break
		}
		position, moves := val[0], val[1]
		for i := 1; i < 7; i++ {
			nextPosition := position + i
			r, c := intToPos(nextPosition)
			if board[r][c] != -1 {
				nextPosition = board[r][c]
			}
			if nextPosition == n*n {
				return moves + 1
			}
			if _, ok := visited[nextPosition]; !ok {
				visited[nextPosition] = true
				queue = append(queue, []int{nextPosition, moves + 1})
			}
		}
	}

	return -1
}
func intToPos(pos int) (int, int) {

	row := (pos - 1) / n
	col := (pos - 1) % n
	if row%2 == 1 {
		col = n - 1 - col
	}
	return n - 1 - row, col

}

func PopFirst(a *[][]int) ([]int, error) {
	if len(*a) <= 0 {
		return []int{}, errors.New("")
	} else {
		val := (*a)[0]
		*a = (*a)[1:]
		return val, nil
	}
}

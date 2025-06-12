package main

func main() {}

func searchMatrix(matrix [][]int, target int) bool {

	m := len(matrix)
	n := len(matrix[0])

	r := m*n - 1
	l := 0

	for l <= r {
		mid := l + (r-l)/2
		row, col := convertToMatrix(n, mid)
		val := matrix[row][col]
		if val == target {
			return true
		}

		if val < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}

func convertToMatrix(n, val int) (int, int) {
	row := val / n
	col := val % n
	return row, col
}

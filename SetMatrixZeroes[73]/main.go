package main

func main() {}

// x= 11111
// y= 22222
// xy=33333
func setZeroes(matrix [][]int) {

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				if matrix[0][j] == 11111 || matrix[0][j] == 33333 || i == 0 {
					matrix[0][j] = 33333
				} else {
					matrix[0][j] = 22222
				}
				if matrix[i][0] == 22222 || matrix[i][0] == 33333 {
					matrix[i][0] = 33333
				} else {
					matrix[i][0] = 11111
				}
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[0][j] == 22222 || matrix[0][j] == 33333 {
				if matrix[i][j] == 11111 || matrix[i][j] == 22222 || matrix[i][j] == 33333 {
					continue
				} else {
					matrix[i][j] = 0
					continue
				}
			}

			if matrix[i][0] == 11111 || matrix[i][0] == 33333 {
				if matrix[i][j] == 11111 || matrix[i][j] == 22222 || matrix[i][j] == 33333 {
					continue
				} else {
					matrix[i][j] = 0
					continue
				}
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 11111 || matrix[i][j] == 22222 || matrix[i][j] == 33333 {
				matrix[i][j] = 0
			}
		}
	}
}

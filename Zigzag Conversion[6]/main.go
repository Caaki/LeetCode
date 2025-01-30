package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))

	//fmt.Println(calculateNeddedCol(15, 5))
}

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	if numRows == 2 {
		return forTwoRows(&s)
	}

	neddedColums := calculateNeddedCol(float64(len(s)), float64(numRows))
	matrix := make([][]rune, numRows)
	for i := range matrix {
		matrix[i] = make([]rune, neddedColums)
	}

	down := true
	i := 0
	j := 0

	for _, char := range s {

		if down {
			matrix[i][j] = char
			i++
			if i == numRows {
				down = false
				j++
				i -= 2
			}
		} else {
			matrix[i][j] = char
			i--
			j++
			if i <= 0 {
				i = 0
				down = true
			}
		}
	}
	toReturn := make([]rune, len(s), len(s))
	current := 0
	for i := 0; i < len(matrix); i++ {
		if i == 0 || i == len(matrix)-1 {
			for j := 0; j < len(matrix[i]); j += numRows - 1 {
				if matrix[i][j] != 0 {
					toReturn[current] = matrix[i][j]
					current++
				}
			}
		} else {
			for j := 0; j < len(matrix[i]); {
				if matrix[i][j] != 0 {
					toReturn[current] = matrix[i][j]
					current++
				}
				j += numRows - 1 - i
				if j < len(matrix[i]) {
					if matrix[i][j] != 0 {
						toReturn[current] = matrix[i][j]
						current++
					}
					j += i
				}
			}
		}
	}
	fmt.Println(len(toReturn))
	return string(toReturn)

}

func calculateNeddedCol(stringLen float64, rows float64) int {
	if stringLen <= rows {
		return 1
	}
	numbsInSquare := (rows * 2) - 2
	toMultiply := stringLen / numbsInSquare
	return int(math.Ceil((rows - 1) * toMultiply))
}

func forTwoRows(s *string) string {
	s1 := ""
	s2 := ""
	for index, ch := range *s {
		if index == 0 || index%2 == 0 {
			s1 += string(ch)
		} else {
			s2 += string(ch)
		}
	}

	return s1 + s2
}

package main

import "fmt"

func main() {
	var b byte = 56
	fmt.Println(b - '0')
}

func isValidSudoku(board [][]byte) bool {
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := range cols {
		cols[i] = make(map[byte]bool, 9)
	}
	for i := range boxes {
		boxes[i] = make(map[byte]bool, 9)
	}

	box := 0
	temp := make(map[byte]bool, 9)
	for i, r := range board {
		for j, v := range r {
			if j != 0 && j%3 == 0 {
				box++
			}
			if !checkAndAdd(boxes[box], v) || !checkAndAdd(temp, v) || !checkAndAdd(cols[j], v) {
				return false
			}
		}
		if i != 0 && (i+1)%3 == 0 {
			box++
		} else {
			box -= 2
		}
		clear(temp)
	}
	return true
}

func checkAndAdd(l map[byte]bool, v byte) bool {
	if v == 46 {
		return true
	}
	if _, ok := l[v]; ok {
		return false
	}
	l[v] = true
	return true
}

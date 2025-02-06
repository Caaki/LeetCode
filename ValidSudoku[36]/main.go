package main

import "fmt"

func main() {
  var b byte = 56
  fmt.Println(b-'0')
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
				continue
			}
			if !checkAndAdd(boxes[box], v) {

				fmt.Println("PUCA KOD BOXA", box)
				return false
			}
			if !checkAndAdd(temp, v) {

				fmt.Println("PUCA KOD REDA")
				return false
			}
			if !checkAndAdd(cols[j], v) {

				fmt.Println("PUCA KOD KOLONA")
				return false
			}
		}
		if i != 0 && (i+1)%3 == 0 {
			fmt.Println("Povecava se box")
			box++
		} else {
			fmt.Printf("Smanjuje se box sa %d na %d, i = %d\n", box, box-2, i)
			box -= 2
		}
		clear(temp)
	}
	for _, m := range cols {
		for k, _ := range m {
			fmt.Printf("%d, ", k)
		}
		fmt.Println()
	}

	return true
}

func checkAndAdd(l map[byte]bool, v byte) bool {
	if v == 46 {
		return true
	}
	if _, ok := l[v]; ok {
		fmt.Println("Vraca false kod vrednosti: ", v)
		return false
	}
	l[v] = true
	return true
}

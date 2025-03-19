package main

import (
	"fmt"
)

func main() {
	//fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	arr := []string{}
	getPosible(n, 1, 0, "(", &arr)
	return arr
}

func getPosible(n int, opened int, closed int, str string, arr *[]string) {
	if opened == n && closed == n {
		*arr = append(*arr, str)
		return
	}
	if opened < n {
		getPosible(n, opened+1, closed, str+"(", arr)
	}
	if opened > closed {
		getPosible(n, opened, closed+1, str+")", arr)
	}

}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(7))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	result := "1"
	for n-1 > 0 {
		temp := makeString(makeArray(result))
		result = temp
		n--
	}
	return result
}

func makeArray(nStr string) [][]int {
	arr := make([][]int, 0)
	for i := 0; i < len(nStr); {
		count := 1
		for i+count < len(nStr) && nStr[i+count] == nStr[i] {
			count++
		}
		v, _ := strconv.Atoi(string(nStr[i]))
		arr = append(arr, []int{v, count})
		i += count
	}
	return arr
}

func makeString(arr [][]int) string {
	str := ""
	for _, v := range arr {
		str += strconv.Itoa(v[1]) + strconv.Itoa(v[0])
	}
	return str
}

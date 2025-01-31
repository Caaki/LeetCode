package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	fmt.Println(reverse(-123))
}

func reverse(x int) int {
	strVal := strconv.Itoa(x)
	if strVal[0] == '-' {
		newStr := ReverseString(strVal[1:])
		intVal, err := strconv.ParseInt(newStr, 10, 32)
		if err != nil {
			return 0
		}
		return int(-(intVal))
	} else {
		newStr := ReverseString(strVal)
		intVal, err := strconv.ParseInt(newStr, 10, 32)
		if err != nil {
			return 0
		}
		return int(intVal)
	}
}

func ReverseString(s string) string {

	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)

}

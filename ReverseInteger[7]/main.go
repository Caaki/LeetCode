package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

func main() {
	fmt.Println(math.MaxInt32)
	fmt.Println(reverse2(1534236469))
}

func reverse2(x int) int {
	var returnVal int32 = 0
	negative := x < 0
	counter := 0
	for x != 0 {
		if counter >= 10 {
			if !negative {
				if (int64(returnVal)*10 + int64(x%10)) > math.MaxInt32 {
					return 0
				}
			} else {
				if (int64(returnVal)*10 + int64(x%10)) < math.MinInt32 {
					return 0
				}
			}
			returnVal = returnVal*10 + int32(x%10)
			x /= 10
			counter++
		} else {
			returnVal = returnVal*10 + int32(x%10)
			x /= 10
			counter++
		}
	}
	return int(returnVal)
}

func reverse(x int) int {
	strVal := strconv.Itoa(x)
	if x < 0 {
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

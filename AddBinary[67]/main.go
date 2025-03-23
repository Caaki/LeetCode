package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("11101", "1101001"))
}
func addBinary(a, b string) string {

	if len(b) > len(a) {
		temp := b
		b = a
		a = temp
	}

	for i := len(b); i < len(a); i++ {
		b = "0" + b
	}
	result := make([]byte, len(a))
	remainder := false
	for i := len(a) - 1; i >= 0; i-- {
		if b[i] != a[i] {
			if remainder {
				result[i] = '0'
				continue
			}
			result[i] = '1'
			continue
		}
		if b[i] == '0' {
			if remainder {
				result[i] = '1'
				remainder = false
				continue
			}
			result[i] = '0'
			continue
		}
		if remainder {
			remainder = true
			result[i] = '1'
			continue
		}
		result[i] = '0'
		remainder = true
	}
	if remainder {
		if a[0] == '1' {
			return "1" + string(result)
		}
		return "1" + string(result[1:])
	}
	return string(result)
}

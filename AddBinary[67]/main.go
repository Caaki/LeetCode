package main

import "fmt"

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
	result := ""
	remainder := false
	for i := len(a) - 1; i >= 0; i-- {
		if b[i] != a[i] {
			if remainder {
				result = "0" + result
				continue
			}
			result = "1" + result
			continue
		}
		if b[i] == '0' {
			if remainder {
				result = "1" + result
				remainder = false
				continue
			}
			result = "0" + result
			continue
		}
		if remainder {
			remainder = true
			result = "1" + result
			continue
		}
		result = "0" + result
		remainder = true
	}

	if remainder {
		if a[0] == '1' {
			return "1" + result
		}
		return "1" + result[1:]
	}
	return result
}

package main

import (
	"fmt"
	"math"
)

// For memory efficiancy i manualy converted runes to their int representation
// + == 43
// " " == 32
// - == 44
// 1 == 49
// 9 == 57
func main() {
	test1 := "-2147483647"
	fmt.Println(myAtoi(test1))
}

func myAtoi(s string) int {
	var n int32 = 0
	st := false
	neg := false
	i := 0
	for _, c := range s {
		if !st {
			if c != 32 && c != 43 && c != 45 {
				if c > 57 || c < 48 {
					return 0
				}
				n = n*10 + c - 48
				st = true
				if n != 0 {
					i++
				}
			}
			if c == 32 {
				continue
			}
			if c == 45 {
				neg = true
				st = true
			}
			if c == 43 {
				st = true
			}
		} else {
			if c > 57 || c < 48 {
				break
			}
			if i > 10 {
				if neg {
					return math.MinInt32
				} else {
					return math.MaxInt32
				}
			}
			if i >= 9 {
				if (int(n)*10)+int(c)-48 > math.MaxInt32 {
					if neg {
						return math.MinInt32
					}
					return math.MaxInt32
				}
				n = n*10 + c - 48
				i++
				continue
			}
			n = n*10 + c - 48
			if n != 0 {
				i++
			}
		}
	}
	if neg {
		return -(int(n))
	}
	return int(n)
}

package main

import "fmt"

var realValues2 = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func main() {

	fmt.Println(romanToInt2("CDIX"))

}
func romanToInt2(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			value, ok := realValues2[fmt.Sprintf("%c%c", s[i], s[i+1])]
			if ok {
				sum += value
				i++
				continue
			}
		}
		sum += realValues2[fmt.Sprintf("%c", s[i])]
	}
	return sum
}

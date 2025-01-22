package main

import (
	"fmt"
)

var greatersMap = map[rune][]rune{'I': {'V', 'X'}, 'V': {'X'}, 'X': {'L', 'C'}, 'L': {'C'}, 'C': {'D', 'M'}, 'D': {'M'}, 'M': {}}
var realValues = map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
var neighbours = map[rune][]rune{'I': {'V', 'X'}, 'V': {'I', 'X'}, 'X': {'I', 'V', 'L', 'C'}, 'L': {'X', 'C'}, 'C': {'X', 'L', 'D', 'M'}, 'D': {'C', 'M'}, 'M': {'C', 'D'}}
var oneLess = map[rune]rune{'I': 'A', 'V': 'A', 'X': 'V', 'L': 'X', 'C': 'L', 'D': 'L', 'M': 'D'}

// func main() {
//
//	fmt.Println(romanToInt("CCCI"))
//
// }
func romanToInt(s string) int {

	sum := 0
	runes := []rune(s)
	var first rune
	var secound rune
	fCount := 0
	sCount := 0
	for i := 0; i < len(runes); i++ {
		first = runes[i]
		fCount = 1
		sCount = 0
		j := i + 1

		for ; j < len(runes) && first == runes[j]; j++ {
			fCount++
		}
		if j >= len(runes) {
			sum += realValues[runes[i]] * fCount
			fmt.Println("1", sum)
			break
		}
		if !isNeighbour(runes[i], runes[j]) {
			sum += realValues[runes[i]] * fCount
			fmt.Println("2", sum, string(runes[i]), string(runes[j]))
			i = j - 1
			continue
		}
		secound = runes[j]
		sCount = 0
		h := j
		for ; h < len(runes) && secound == runes[h]; h++ {
			sCount++
		}

		if h == len(runes) {
			sum += calculated(first, secound, fCount, sCount)
			break
		}
		if runes[h] == first || (runes[h] == oneLess[first] && realValues[first] > realValues[secound]) {
			sum += realValues[runes[i]] * fCount
			fmt.Println(string(runes[h]), string(first), "One less is: ", string(oneLess[first]), sum)
			i = j - 1
			continue
		}

		sum += calculated(first, secound, fCount, sCount)
		fmt.Println(sum)
		i = h - 1
	}
	return sum
}

func calculated(first rune, secound rune, fCount int, sCount int) int {

	runesFromMap := greatersMap[first]

	isGreater := false
	for _, greater := range runesFromMap {
		if greater == secound {
			isGreater = true
		}
	}

	if isGreater {
		result := (realValues[secound] * sCount) - (realValues[first] * fCount)
		return result
	} else {

		result := (realValues[first] * fCount) + (realValues[secound] * sCount)
		return result
	}

}
func isNeighbour(first rune, secound rune) bool {
	for _, value := range neighbours[first] {
		if value == secound {
			return true
		}
	}
	return false
}

package main

import "strconv"

func main() {}

func minMaxDifference(num int) int {

	temp := strconv.Itoa(num)
	minVal := ""
	maxVal := ""

	c1 := byte('-')
	c2 := byte('-')
	for i, v := range temp {
		if c1 != '-' && temp[i] == c1 {
			maxVal += "9"
			continue
		}
		if temp[i] == '9' {
			maxVal += "9"
			continue
		}
		if c1 == '-' {
			c1 = temp[i]
			maxVal += "9"
			continue
		}
		maxVal += string(v)
	}
	for i, v := range temp {
		if c2 != '-' && temp[i] == c2 {
			minVal += "0"
			continue
		}
		if temp[i] == '0' {
			minVal += "0"
			continue
		}
		if c2 == '-' {
			c2 = temp[i]
			minVal += "0"
			continue
		}
		minVal += string(v)
	}

	l, _ := strconv.Atoi(maxVal)

	m, _ := strconv.Atoi(minVal)
	return l - m
}

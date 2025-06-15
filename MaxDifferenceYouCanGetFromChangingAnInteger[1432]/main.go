package main

import "strconv"

func main() {

}

func maxDiff(num int) int {

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

	stack := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if temp[0] != '1' {
		stack = stack[1 : len(stack)-1]
	}

	for i, v := range temp {
		strToInt, _ := strconv.Atoi(string(v))
		intToStr := strconv.Itoa(stack[0])
		if c2 != '-' && temp[i] == c2 {
			minVal += intToStr
			continue
		}
		if temp[i] == '0' {
			minVal += "0"
			continue
		}
		if temp[i] == '1' {
			minVal += "1"
			continue
		}

		if c2 == '-' {
			if stack[0] < strToInt {
				c2 = temp[i]
				minVal += intToStr
				continue
			}
		}
		minVal += string(v)
	}
	l, _ := strconv.Atoi(maxVal)

	m, _ := strconv.Atoi(minVal)
	return l - m
}

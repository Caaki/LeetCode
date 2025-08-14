package main

import (
	"fmt"
	"strconv"
)

func main() {}

func largestGoodInteger(num string) string {
	largest := -1
	selected := num[0]
	count := 1
	toReturn := ""
	for i := 1; i < len(num); i++ {
		if num[i] == selected {
			count++
			if count == 3 {
				temp,_ := strconv.Atoi(string(selected))
				if temp > largest{
					toReturn = fmt.Sprintf("%c%c%c",selected,selected,selected)
					largest = temp
				}
			}
		}else{

			selected = num[i]
			count = 1
		}
	}
	return toReturn
}

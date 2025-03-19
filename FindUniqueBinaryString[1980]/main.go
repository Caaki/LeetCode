package main

import (
	"fmt"
	"strconv"
)

func main() {

}
func findDifferentBinaryString(nums []string) string {
	n := int64(len(nums))
	mapedValues := make(map[int64]bool, len(nums))
	for _, v := range nums {
		value, _ := strconv.ParseInt(v, 2, 0)
		mapedValues[value] = true
	}
	format := fmt.Sprintf("%%0%db", n)
	for i := int64(0); i < n+1; i++ {
		if _, ok := mapedValues[i]; !ok {
			return fmt.Sprintf(format, i)
		}
	}
	return ""
}

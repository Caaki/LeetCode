package main

import "fmt"


func main(){
	fmt.Println(countMaxOrSubsets([]int{1,4}))
}

func countMaxOrSubsets(nums []int) int{
	maxVal:=0

	for _,num := range nums{
		maxVal |=num
	}
	return maxVal
}

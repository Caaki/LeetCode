package main

import (
	"fmt"
	"maps"
	"slices"
	"sort"
)

func main (){
	fmt.Println(findErrorNums([]int{1,2,2,4}))
}

func findErrorNums(nums []int) []int {
	
	duped := make(map[int]bool,0)
	start :=1
	sort.Ints(nums)
	toReturn := make(map[int]bool,0)

	for i:=0; i < len(nums)|| start > len(nums); i++  {
		fmt.Println(nums[i], start)
		if _,ok := duped[nums[i]]; ok {
			fmt.Println("first")
			toReturn[nums[i]]=true
		}
		duped[nums[i]]=true
		if start < nums[i] {
			fmt.Println("second")
			toReturn[start]=true
			start++
			i--
			continue
		}
		if start!=nums[i]{
			fmt.Println("third")
			toReturn[start]=true
		}
		start++
	}
	values :=slices.Collect(maps.Keys(toReturn))
	return values
}

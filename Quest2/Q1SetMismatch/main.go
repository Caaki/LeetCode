package main

import (
	"fmt"
	"sort"
)

func main (){
	fmt.Println(findErrorNums([]int{1,2,2,4}))
}

func findErrorNums(nums []int) []int{

	if len(nums) ==2{
		if nums[0]==1{
			return []int{1,2}
		}else{
			return []int{2,1}
		}
	}

	sort.Ints(nums)
	m:=0
	e:=0
	c,i:=1,0
	for ; i<len(nums);i++{
		if e != 0 && m != 0{
			return []int{e,m}
		}
		if c == nums[i]{
			c++
			continue
		}
		if c < nums[i]{
			m = c
			c++
			i--
			continue
		}
		if c > nums[i]{
			e = nums[i]
		}
	}
	if m == 0{
		m = c
	}

	return []int{e,m}
}


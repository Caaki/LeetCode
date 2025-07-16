package main

import "fmt"


func main() {
	fmt.Println(findDuplicate([]int{2,3,4,1,5,6,7,1}))
}

func findDuplicate(nums []int) int {

	values := make(map[int]bool,0)


	for _, v:= range nums{
		_,ok:=values[v]
		if ok{
			return v
		}else{
			values[v]=true
		}
	}
	return 0
}

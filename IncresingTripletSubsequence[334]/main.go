package main

import "math"


func increasingTriplet(nums []int) bool {
	if len(nums)<3{
		return false
	}

	min1 :=math.MaxInt
	min2 := math.MaxInt
	
	for _,v := range nums {
		
		if v <=min1{
			min1 = v
		}else if v <= min2{
			min2=v
		}else{
			return true 
		}
	}
	return false
}

package main

/*
Given the array nums, for each nums[i] find out how many numbers in the array are smaller than it.
That is, for each nums[i]
you have to count the number of valid j's such that j != i and nums[j] < nums[i].

Return the answer in an array.
*/

func main(){}

type N struct {
	num, count int
}

func smallerNumbersThanCurrent(nums []int) []int {
	nCount := make(map[int] int)

	for _,v := range nums{
		nCount[v]++
	}
	
	rValues := make([]int,len(nums))
	
	for i,cur := range nums{
		count := 0
		for k,v :=range nCount{
			if cur > k{
				cur+= v
			}
		}
		rValues[i] = count
	}

	return rValues
}

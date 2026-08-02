package main

import "sort"

func findDisappearedNumbers(nums []int) []int {
	sort.Ints(nums)
	current :=1
	missing := make([]int,0)
	for i:=0; i < len(nums);{
		if nums[i] > current{
			missing = append(missing, current)
			current++
			continue
		}
		if nums[i] == current{
			current++
		}
		i++
	}

	for current <= len(nums){
		missing = append(missing, current)
		current++
	}

	return missing

}

package main

func main(){}


func countHillValley(nums []int) int{
	lh := nums[0]
	current:=nums[1]
	count:=0
	for i:=1; i < len(nums);i++{
		for lh == current {
			current=nums[i]
			i++
			if i == len(nums){
				return count
			}
		}

		if current == nums[i]{
			continue
		}
		if (lh < current && current < nums[i]) || (lh > current && current > nums[i]){
			lh = current
			current = nums[i]
			continue
		}
		if lh > current && current < nums[i]{
			count++
			lh = current
			current = nums[i]
			continue
		}
		if lh < current && current > nums[i]{
			count++
			lh = current
			current = nums[i]
			continue
		}
	}
	return count
}


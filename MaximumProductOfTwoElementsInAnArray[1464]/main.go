package main

func maxProduct(nums []int ) int{
	l := 0
	l2 := 0

	if nums[0] > nums[1]{
		l = nums[0]
		l2 = nums[1] 
	}else {
		l = nums[1]
		l2 = nums[0]
	}

	for i := 2; i <len(nums);i++{
		if nums[i] >= l {
			l2 = l
			l = nums[i]
		}else if nums[i]>l2{
			l2 = nums[i]
		}
	}

	return (l-1)*(l2-1)
}

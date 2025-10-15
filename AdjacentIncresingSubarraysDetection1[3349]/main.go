package main



func hasIncreasingSubarrays(nums []int, k int) bool {
	if k == 1 && len(nums)>=2{
		return true
	}
	count := 1
	current := nums[0]
	last := 0
	for i :=1 ; i < len(nums)-1 ; i++{
		if current < nums[i]{
			count++
			current=nums[i]
			if count == k {
				count = 1
                current = nums[i+1]
				for i=i+2; i < len(nums); i++{
					if current < nums[i]{
						count++
						current=nums[i]
						if count >=k {
							return true 
						}
					}else{
						last++
						i = last
						count = 1
						current = nums[i]
						break
					}
				}
			}
		}else {
			last++
			i = last
			count = 1
			current = nums[i]	}
		}
		return false
	}

package main

func productExceptSelf(nums []int) []int {
	
	prefixSum := make([]int, len(nums))
	sufixSum := make([]int, len(nums))
	
	sum :=1
	for i,v:= range nums{
        sum = sum*v
		prefixSum[i]=sum
	}
	
	sum =1
	for i:=len(sufixSum)-1; i >=0;i--{
        sum = sum*nums[i]
		sufixSum[i]=sum
	}

	toReturn := make([]int,len(nums))
	
	for i:=0; i<len(nums); i++{
		if i == 0{
			toReturn[0] = sufixSum[1]
            continue
		}
		if i == len(nums)-1{
			toReturn[len(nums)-1] = prefixSum[len(nums)-2]
            continue
		}
		toReturn[i]=prefixSum[i-1]*sufixSum[i+1]
	}
	return toReturn
}

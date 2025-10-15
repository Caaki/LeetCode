package main


func moveZeroes(nums []int) {
	zeroesCount :=0
	for _,v := range nums {
		if v == 0{
			zeroesCount++
		}
	}
	if zeroesCount == 0 || zeroesCount == len(nums){
		return
	}
	
	for i :=0 ; i < len(nums) ; i++{
		if nums[i]==0{
			for j:=i+1; j<len(nums); j++{
				if nums[j]!=0{
					nums[i]=nums[j]
                    nums[j]=0
					break
				}
			}
		}
	}
}

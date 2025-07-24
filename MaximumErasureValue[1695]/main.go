package main

func main(){}
//delete(m,k)


func maximumUniqueSubarray(nums []int) int{
	valies :=make(map[int]bool,0)

	
	l:=0
	r:=1
	valies[nums[0]]=true
	tempScore:=nums[0]
	maxScore:=nums[0]
	for r < len(nums) {
       for valies[nums[r]] {
            delete(valies, nums[l])
            tempScore -= nums[l]
            l++
        }
		valies[nums[r]]=true
		tempScore+=nums[r]
		if tempScore>maxScore{
			maxScore=tempScore
		}
        r++

	}
	return maxScore
}


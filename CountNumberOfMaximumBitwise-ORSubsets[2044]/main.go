package main



func main(){
}

func countMaxOrSubsets(nums []int) int{
	maxVal:=0
	count :=0
	for _,num := range nums{
		maxVal |=num
	}

	var backTrack func(index,currentOR int) 
	backTrack=func(index, currentOR int){
	if index == len(nums){
		if maxVal==currentOR{
			count++
		}
		return
	}
		backTrack(index+1,currentOR)
		backTrack(index+1,currentOR|nums[index])
	} 

	backTrack(0,0)

	return count
}


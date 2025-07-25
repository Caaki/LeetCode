package main



func main(){}


func maxSum(nums []int ) int {
	values := make(map[int]bool,0)
	sum:=0
	maxVal:=nums[0]
	for _,v := range nums {
		if v >0{
			values[v] = true
		}
		if maxVal < v{
			maxVal = v
		}
	}
	for k,_:=range values {
		sum +=k
	}
	if maxVal<0{
		return maxVal
	}

	return sum
}

func maxSum2(nums []int) int{
	values := make([]int8,101)
	maxVal:=nums[0]
	sum := 0
	for _,v :=range nums{
		if v > 0{
            if values[v]==0{
	            sum += v
            }
			values[v]=1
		}
		if maxVal < v{
			maxVal = v
		}
	}
	if maxVal <= 0 {
		return maxVal
	}
	return sum
}




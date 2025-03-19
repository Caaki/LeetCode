package main

func main(){}

func applyOperations(nums []int) []int {
  temp:=0
  for i,j:=0,0 ; i<len(nums);i++{
    if len(nums)>i+1 && nums[i] == nums[i+1] && nums[i]!=0{
      temp=nums[i]*2
      nums[i]=0
      nums[i+1]=0
      nums[j]=temp
      j++
      i++
      continue
    }
    if nums[i]>0{
      temp=nums[i]
      nums[i]=0
      nums[j]=temp
      j++
    }
  }

  return nums
}

package main


func main (){
}

func productExceptSelf(nums []int) []int {
  if len(nums)==1{
    return []int{0}
  }
  prefix:= make([]int32,len(nums))
  sufix:= make([]int32,len(nums))


  var sum int32= 1
  for index, i:= range nums {
    sum *= int32(i)
    prefix[index]= sum
  }
  sum = 1
  for i := len(nums)-1; i>=0 ; i--{
    sum *= int32(nums[i])
    sufix[i] =sum
  }
  var output []int
  for i := 0 ; i<len(nums) ; i++{
    if i == 0 {
      output = append(output, 1* int(sufix[i+1]))
      continue
    }
    if i == len(nums)-1{
      output = append(output, 1*int(prefix[len(nums)-2]))
      continue
    }
    output = append(output, int(prefix[i-1] * sufix[i+1]))
  }
  return output
}




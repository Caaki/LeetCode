package main

import(
  "fmt"
)

func main() {
  fmt.Println(minOperations([]int{0,1,1,1,0,0}))
}

func minOperations(nums []int) int{
  count:=0

  l,r:=0,2
  for r<len(nums) {
    if nums[l] == 0{
      changeVal(nums,l,r)
      count++
    }
    l++
    r++
  }
  if nums[l-1] != nums[l] || nums[r-1] != nums[l-1]{
    return -1
  }
  return count
}

func changeVal(nums []int, l,r int){
  for l<=r{
    nums[l] = nums[l]^(1 << 0)
    l++
  }
}

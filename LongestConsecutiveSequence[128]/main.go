package main

import (

  "fmt"
  "time"
)

func main(){

  fmt.Println(longestConsecutive([]int{0,3,7,2,5,8,4,6,0,1}))
}

func longestConsecutive(nums []int) int{
  values:= make(map[int]bool)
  longest:=0
  for _,num :=range nums{
    if _,ok:= values[num];ok{
      continue
    }
    tempCount:=1
    values[num] = true
    tempCount +=getFirst(values, num)
    tempCount+=getLast(values,num)
    if longest < tempCount{
      longest =tempCount
    }
  }
  return longest
}

func getFirst(values map[int]bool,num int) int{
  if _,ok:= values[num-1];ok{
    return 1+getFirst(values, num-1)
  }
  return 0
}

func getLast(values map[int]bool, num int) int{
  if _,ok:= values[num+1];ok{
    return 1+ getLast(values,num+1)
  }
  return 0
}


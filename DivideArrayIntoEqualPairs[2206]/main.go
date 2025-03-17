package main

import (
  "fmt"
)
func main (){
  fmt.Println(divideArray([]int{3,2,3,2,2,2}))
}


func divideArray(nums []int)bool{
  maxVal:=1
  for _,v:=range nums{
    if v>maxVal{
      maxVal = v
    }
  }

  allVals:=make([]int16,maxVal+1)
  for _,v:=range nums{
    allVals[v]++
  }

  for _,v:=range allVals{
    if v%2!=0 && v!=0{
      return false
    }
  }
  return true
}

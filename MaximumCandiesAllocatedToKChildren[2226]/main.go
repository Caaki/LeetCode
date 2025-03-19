package main

import (
  "math"
  "fmt"
)

func main (){
  fmt.Println(maximumCandies([]int{9,10,1,2,10,9,9,10,2,2}, 3))
}

func maximumCandies(candies []int, k int64) int {
  l:=0
  count:=0
  for _,v:= range candies{
    count +=v
  }
  result :=1
  r := count/int(k)

  if count < int(k){
    return 0
  }
  if count == int(k){
    return 1
  }
  
  
  for l <= r{
    mid:= (l+r)/2
    temp:=0
    for _,v := range candies{
      temp+=int(math.Floor(float64(v)/float64(mid)))
    }
    if temp >= int(k){
      result=mid
      l=mid+1
    }else{
      r = mid-1
    }
  }
  return result
}

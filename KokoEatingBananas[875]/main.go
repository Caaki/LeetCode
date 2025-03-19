package main 

import(
  "math"
  "fmt"
)
func main(){
  fmt.Println(minEatingSpeed([]int{3,6,7,11},8))
}

func minEatingSpeed(piles []int, h int) int {
  l:=1
  r :=piles[0]
  for _,v :=range piles{
    if v > r {
      r = v
    }
  }
  result :=r
  for l <=r {
    mid := (r+l)/2
    temp:=0
    for _,v := range piles{
      temp+= int(math.Ceil(float64(v)/float64(mid)))
    }
    if temp <= h{
      result = mid
      r = mid-1
    }else {
      l = mid+1
    }
  }
  return result
}
